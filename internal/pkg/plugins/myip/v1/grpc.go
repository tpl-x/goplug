package v1

import (
	"context"
	v1 "github.com/tpl-x/goplug/api/plugins/myip/v1"
	"google.golang.org/grpc"
)

var _ v1.MyIpServiceServer = (*myIpFinderGrpcServer)(nil)

type myIpFinderGrpcServer struct {
	Impl MyIPFinder
}

func (m *myIpFinderGrpcServer) FindMyIP(ctx context.Context, request *v1.FindMyIPRequest) (*v1.FindMyIPResponse, error) {
	ip, location, region, err := m.Impl.GetMyIp()
	if err != nil {
		return nil, err
	}
	resp := &v1.FindMyIPResponse{
		IpAddress: ip,
		Location:  location,
		Region:    region,
	}
	return resp, nil
}

var _ v1.MyIpServiceClient = (*myIpFinderGrpcClient)(nil)

type myIpFinderGrpcClient struct {
	client v1.MyIpServiceClient
}

func (m *myIpFinderGrpcClient) FindMyIP(ctx context.Context, in *v1.FindMyIPRequest, opts ...grpc.CallOption) (*v1.FindMyIPResponse, error) {
	return m.client.FindMyIP(ctx, in, opts...)
}
