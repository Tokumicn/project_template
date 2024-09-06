package main

import (
	"context"
	"flag"
	assetfs "github.com/elazarl/go-bindata-assetfs"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpc_tlp/global"
	"grpc_tlp/internal/middleware"
	"grpc_tlp/pkg/swagger"
	"grpc_tlp/pkg/tracer"
	pb "grpc_tlp/proto"
	"grpc_tlp/server"
	"log"
	"net/http"
	"path"
	"strings"
)

var port string

func init() {
	flag.StringVar(&port, "port", "8004", "启动端口号")
	flag.Parse()

	// 启动链路追踪
	err := setupTracer()
	if err != nil {
		log.Fatalf("init.setpuTracer err: %v", err)
	}
}

func setupTracer() error {
	// 创建 jaeger client
	jaegerTracer, _, err := tracer.NewJaegerTracer("grpc-tlp-service", "127.0.0.1:6831")
	if err != nil {
		return err
	}
	global.Tracer = jaegerTracer
	return nil
}

func main() {
	err := RunServer(port)
	if err != nil {
		log.Fatalf("Run Serve err: %v", err)
	}
}

func RunServer(port string) error {
	httpMux := runHttpServer()
	grpcS := runGRPCServer()
	gatewayMux := runGrpcGatewayServer()

	httpMux.Handle("/", gatewayMux)
	return http.ListenAndServe(":"+port, grpcHandlerFunc(grpcS, httpMux))
}

// 启动 GRPC Gateway 服务
func runGrpcGatewayServer() *gwruntime.ServeMux {
	endpoint := "0.0.0.0:" + port
	gwmux := gwruntime.NewServeMux()
	options := []grpc.DialOption{grpc.WithInsecure()}
	_ = pb.RegisterTagServiceHandlerFromEndpoint(context.Background(), gwmux, endpoint, options)
	return gwmux
}

// 启动 httpServer
func runHttpServer() *http.ServeMux {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("pong"))
	})
	prefix := "/swagger-ui/"
	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    swagger.Asset,
		AssetDir: swagger.AssetDir,
		Prefix:   "third_party/swagger-ui",
	})

	serveMux.Handle(prefix, http.StripPrefix(prefix, fileServer))
	serveMux.HandleFunc("/swagger/", func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasSuffix(r.URL.Path, "swagger.json") {
			http.NotFound(w, r)
			return
		}

		p := strings.TrimPrefix(r.URL.Path, "/swagger/")
		p = path.Join("proto", p)
		http.ServeFile(w, r, p)
	})
	return serveMux
}

// 启动GRPC服务
func runGRPCServer() *grpc.Server {
	opts := []grpc.ServerOption{
		// 虽然 grpc-go 官方只允许设置一个拦截器
		// github.com/grpc-ecosystem/go-grpc-middleware 提供的链式插件方案
		// 实现多拦截器的需要
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				middleware.AccessLog,
				middleware.ErrorLog,
				middleware.Recovery,
				middleware.ServerTracing,
			)),
	}

	s := grpc.NewServer(opts...)
	pb.RegisterTagServiceServer(s, server.NewTagServer())
	reflection.Register(s)

	return s
}

func grpcHandlerFunc(grpcSvc *grpc.Server, otherHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcSvc.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}
