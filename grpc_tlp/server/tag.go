package server

import (
	"context"
	"google.golang.org/grpc"
	pb "grpc_tlp/proto"
)

type TagServer struct {
	pb.UnimplementedTagServiceServer
}

func NewTagServer() *TagServer {
	return &TagServer{}
}

func (t *TagServer) GetTagList(ctx context.Context, r *pb.GetTagListRequest) (*pb.GetTagListReply, error) {
	//api := bapi.NewAPI("http://127.0.0.1:8000")
	//body, err := api.GetTagList(ctx, r.GetName())
	//if err != nil {
	//	return nil, errcode.TogRPCError(errcode.ErrorGetTagListFail)
	//}

	tagList := pb.GetTagListReply{}
	//err = json.Unmarshal(body, &tagList)
	//if err != nil {
	//	return nil, errcode.TogRPCError(errcode.Fail)
	//}
	return &tagList, nil
}

func GetClientConn(ctx context.Context, target string, opts []grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(opts, grpc.WithInsecure())
	return grpc.DialContext(ctx, target, opts...)
}
