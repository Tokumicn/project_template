package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"

	v1 "github.com/Tokumicn/kratos_tlp/api/helloworld/v1"
	"github.com/Tokumicn/kratos_tlp/internal/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc *biz.GreeterUsecase
	l  *log.Helper
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase, logger log.Logger) *GreeterService {
	return &GreeterService{uc: uc, l: log.NewHelper(logger)}
}

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	s.l.WithContext(ctx).Debugf("SayHello") // TODO 增加日志
	//go func() {
	//	defer func() {
	//		if err := recover(); err != nil {
	//			s.l.WithContext(ctx).Errorf("panic: %v", err) // TODO 增加日志
	//		}
	//	}()
	//}()

	time.Sleep(time.Second * 1) // TODO 模拟耗时
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}
