package grpc

import (
	"context"
	"google.golang.org/grpc/metadata"
)

func withAuditMD(ctx context.Context, id, name, email, role string) context.Context {
	md := metadata.New(map[string]string{
		"ts-user-id":    id,
		"ts-user-name":  name,
		"ts-user-email": email,
		"ts-user-role":  role,
	})
	return metadata.NewIncomingContext(ctx, md)
}
