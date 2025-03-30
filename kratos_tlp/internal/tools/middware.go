package tools

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/google/uuid"
)

// TraceIDKey Context 中 TraceID 的建
const (
	TraceIDKey     = "trace_id"
	TraceHeaderKey = "X-Trace-ID"
)

// TraceMiddleware 处理 TraceID 的中间件
func TraceMiddleware() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			traceID := ""

			// 从 Header 中获取 TraceID
			tr, ok := transport.FromServerContext(ctx)
			if ok {
				traceID = tr.RequestHeader().Get(TraceHeaderKey)
			}

			if traceID == "" {
				// 没有就新生成一个
				traceID = uuid.New().String()
			}

			// 将traceID设置到context中
			ctx = context.WithValue(ctx, TraceIDKey, traceID)
			tr.ReplyHeader().Set(TraceHeaderKey, traceID)
			return handler(ctx, req)
		}
	}
}

// TempTraceID 获取 traceID
func TempTraceID() log.Valuer {
	return func(ctx context.Context) interface{} {
		if ctx == nil {
			return ""
		}

		if traceID, ok := ctx.Value(TraceIDKey).(string); ok {
			return traceID
		}
		return ""
	}
}
