package grpc

import (
	"context"
	"errors"
	apperrors "property-service/pkg/errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ErrToStatus converts an error into the corresponding gRPC status,
// adding the AppError's Type as an extra detail.
func ErrToStatus(err error) error {
	var appErr apperrors.AppError
	if errors.As(err, &appErr) {
		grpcCode := appErr.Code() // mapErrorCode converts your code (e.g. codes.Internal) to a gRPC code.
		baseMsg := appErr.Unwrap().Error()
		st := status.New(codes.Code(grpcCode), baseMsg)
		return st.Err()
	}
	return status.Errorf(codes.Internal, err.Error())
}

// UnaryErrorInterceptor is a unary interceptor that converts internal errors to gRPC statuses.
func UnaryErrorInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	resp, err := handler(ctx, req)
	if err != nil {
		// Convert the error using your helper.
		return nil, ErrToStatus(err)
	}
	return resp, nil
}
