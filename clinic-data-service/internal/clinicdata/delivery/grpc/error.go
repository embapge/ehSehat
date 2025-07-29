package grpc

import (
	"clinic-data-service/internal/clinicdata/app"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// mapErrorToStatus maps app-layer errors to gRPC status codes
func mapErrorToStatus(err error) error {
	switch err {
	case app.ErrNotFound:
		return status.Error(codes.NotFound, err.Error())
	case app.ErrInvalidInput, app.ErrMissingFields, app.ErrMissingID:
		return status.Error(codes.InvalidArgument, err.Error())
	case app.ErrInternal:
		return status.Error(codes.Internal, err.Error())
	default:
		return status.Error(codes.Unknown, err.Error())
	}
}
