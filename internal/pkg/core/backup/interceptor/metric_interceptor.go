package interceptor

import (
	"context"
	"tasks/internal/pkg/core/metric"

	"google.golang.org/grpc"
)

func ClientMetricInterceptor() grpc.UnaryClientInterceptor {
	return func(parentCtx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		err := invoker(parentCtx, method, req, reply, cc, opts...)
		if err != nil {
			metric.Instance().Inc("backup_req_out_err")
		} else {
			metric.Instance().Inc("backup_req_out_ok")
		}
		return err
	}
}

func ServerMetricInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		res, err := handler(ctx, req)
		if err != nil {
			metric.Instance().Inc("backup_req_in_err")
		} else {
			metric.Instance().Inc("backup_req_in_ok")
		}
		return res, err
	}
}
