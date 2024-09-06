package middleware

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"grpc_tlp/global"
	"grpc_tlp/pkg/errcode"
	"grpc_tlp/pkg/metatext"
	"log"
	"runtime/debug"
	"time"
)

// ServerTracing 链路追踪中间件
func ServerTracing(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	// Grpc 调用中传递 Header 操作通过 FromIncomingContext
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		md = metadata.New(nil)
	}

	// TODO  添加注释
	parentSpanContext, _ := global.Tracer.Extract(opentracing.TextMap, metatext.MetadataTextMap{MD: md})
	spanOpts := []opentracing.StartSpanOption{
		opentracing.Tag{Key: string(ext.Component), Value: "gRPC"},
		ext.SpanKindRPCServer,
		ext.RPCServerOption(parentSpanContext),
	}
	span := global.Tracer.StartSpan(info.FullMethod, spanOpts...)
	defer span.Finish()

	// 将带有 tracing 标记的 span 写入上下文 并继续传递
	ctx = opentracing.ContextWithSpan(ctx, span)
	return handler(ctx, req)
}

// AccessLog 访问日志中间件：输出接口、耗时、出入参数
func AccessLog(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	requestLog := "access request log: method: %s, begin_time: %d, request: %v"
	beginTime := time.Now().Local().Unix()
	log.Printf(requestLog, info.FullMethod, beginTime, req)

	resp, err := handler(ctx, req)

	responseLog := "access response log: method: %s, begin_time: %d, end_time: %d, response: %v"
	endTime := time.Now().Local().Unix()
	log.Printf(responseLog, info.FullMethod, beginTime, endTime, resp)
	return resp, err
}

// ErrorLog 错误日志中间件：将内部错误信息重构为业务错误信息
func ErrorLog(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	resp, err := handler(ctx, req)
	if err != nil {
		errLog := "error log: method: %s, code: %v, message: %v, details: %v"
		s := errcode.FromError(err)
		log.Printf(errLog, info.FullMethod, s.Code(), s.Err().Error(), s.Details())
	}
	return resp, err
}

// Recovery panic统一恢复中间件
func Recovery(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	defer func() {
		if e := recover(); e != nil {
			recoveryLog := "recovery log: method: %s, message: %v, stack: %s"
			log.Printf(recoveryLog, info.FullMethod, e, string(debug.Stack()[:]))
		}
	}()

	return handler(ctx, req)
}
