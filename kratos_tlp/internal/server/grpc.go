package server

import (
	v1 "github.com/Tokumicn/kratos_tlp/api/helloworld/v1"
	"github.com/Tokumicn/kratos_tlp/internal/conf"
	"github.com/Tokumicn/kratos_tlp/internal/service"
	"github.com/Tokumicn/kratos_tlp/internal/tools"
	"github.com/go-kratos/kratos/v2/middleware/logging"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, greeter *service.GreeterService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			tools.TraceMiddleware(), // 自定义的TraceID中间件
			logging.Server(logger),  // 添加全局日志
			// tracing.Server(), // 官方支持的链路追踪组件 "github.com/go-kratos/kratos/v2/middleware/tracing"
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterGreeterServer(srv, greeter)
	return srv
}
