package server

import (
	v1 "github.com/Tokumicn/kratos_tlp/api/helloworld/v1"
	"github.com/Tokumicn/kratos_tlp/internal/conf"
	"github.com/Tokumicn/kratos_tlp/internal/service"
	"github.com/Tokumicn/kratos_tlp/internal/tools"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.GreeterService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tools.TraceMiddleware(), // 自定义的TraceID中间件
			logging.Server(logger),  // 添加全局日志
			// tracing.Server(), // 官方支持的链路追踪组件 "github.com/go-kratos/kratos/v2/middleware/tracing"
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterGreeterHTTPServer(srv, greeter)
	return srv
}
