package main

import (
	"context"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"grpc_tlp/global"
	"grpc_tlp/internal/middleware"
	pb "grpc_tlp/proto"
	"lego_lib/tracer"
	"log"
)

func init() {
	err := setupTracer()
	if err != nil {
		log.Fatalf("init.setupTracer err: %v", err)
	}
}

func main() {
	ctx := context.Background()
	// GRPC中 Header中传递的信息通过该方式添加， 使用Append而非New 防止链路信息被覆盖
	newCtx := metadata.AppendToOutgoingContext(ctx, "tom-test", "Tom 学 GRPC")
	clientConn, err := GetClientConn(newCtx, "localhost:8004", []grpc.DialOption{grpc.WithUnaryInterceptor(
		grpc_middleware.ChainUnaryClient(
			middleware.UnaryContextTimeout(), // 超时处理中间件
			middleware.ClientTracing(),       // Tracing中间件
		),
	)})
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	defer clientConn.Close()
	tagServiceClient := pb.NewTagServiceClient(clientConn)
	resp, err := tagServiceClient.GetTagList(newCtx, &pb.GetTagListRequest{Name: "Go"})
	if err != nil {
		log.Fatalf("tagServiceClient.GetTagList err: %v", err)
	}
	log.Printf("resp: %v", resp)
}

// GetClientConn 创建 GRPC Client
func GetClientConn(ctx context.Context, target string, opts []grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(opts, grpc.WithInsecure())
	return grpc.DialContext(ctx, target, opts...)
}

func setupTracer() error {
	var err error
	jaegerTracer, _, err := tracer.NewJaegerTracer("grpc-tlp-cli", "127.0.0.1:6831")
	if err != nil {
		return err
	}
	global.Tracer = jaegerTracer
	return nil
}
