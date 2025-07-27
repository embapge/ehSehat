package utils

import (
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GRPCErrorToHTTPError(err error) error {
	if err == nil {
		return nil
	}
	// MongoDB error handling
	if err == mongo.ErrNoDocuments {
		return status.Error(codes.NotFound, "data not found")
	}
	// MongoDB write error
	if we, ok := err.(mongo.WriteException); ok {
		for _, e := range we.WriteErrors {
			switch e.Code {
			case 11000: // duplicate key
				return status.Error(codes.AlreadyExists, e.Message)
			}
		}
		return status.Error(codes.Internal, we.Error())
	}
	// Custom error type handling
	if be, ok := err.(*badRequestError); ok {
		return status.Error(codes.InvalidArgument, be.Error())
	}
	// Default: error internal
	return status.Error(codes.Internal, err.Error())
}
