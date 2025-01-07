package v1

import (
	"context"
	v1 "github.com/tpl-x/goplug/api/plugins/myip/v1"
	"google.golang.org/grpc"
)

var _ v1.MyIpServiceServer = (*MyIpFinderGrpcServer)(nil)

type MyIpFinderGrpcServer struct {
	Impl MyIPFinder
}

func (m *MyIpFinderGrpcServer) FindMyIP(ctx context.Context, request *v1.FindMyIPRequest) (*v1.FindMyIPResponse, error) {
	//TODO implement me
	panic("implement me")
}

var _ v1.MyIpServiceClient = (*MyIpFinderGrpcClient)(nil)

type MyIpFinderGrpcClient struct {
	client v1.MyIpServiceClient
}

func (m *MyIpFinderGrpcClient) FindMyIP(ctx context.Context, in *v1.FindMyIPRequest, opts ...grpc.CallOption) (*v1.FindMyIPResponse, error) {
	//TODO implement me
	panic("implement me")
}
