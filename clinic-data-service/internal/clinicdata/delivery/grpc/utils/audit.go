package utils

import (
	"context"
	"google.golang.org/grpc/metadata"
)

// ExtractAudit returns user metadata (can be empty string if not present)
func ExtractAudit(ctx context.Context) (id, name, email, role string) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", "", "", ""
	}
	return get(md, "ts-user-id"), get(md, "ts-user-name"), get(md, "ts-user-email"), get(md, "ts-user-role")
}

func get(md metadata.MD, key string) string {
	if val := md[key]; len(val) > 0 {
		return val[0]
	}
	return ""
}
