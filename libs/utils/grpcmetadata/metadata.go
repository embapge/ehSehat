package grpcmetadata

import (
	"context"

	"google.golang.org/grpc/metadata"
)

func GetMetadataFromContext(ctx context.Context) (metadata.MD, bool) {
	md, ok := metadata.FromIncomingContext(ctx)
	return md, ok
}
