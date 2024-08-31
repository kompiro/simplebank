package gapi

import (
	"context"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

const (
	grpcGatewayUserAgentHeader = "grpcgateway-user-agent"
	userAgentHeader            = "user-agent"
	xForwardedForHeader        = "x-forwarded-for"
)

type Metadata struct {
	UserAgent string
	ClientIP  string
}

func (server *Server) extractMetadata(ctx context.Context) *Metadata {
	md := &Metadata{}

	if metadata, ok := metadata.FromIncomingContext(ctx); ok {
		if userAgents := metadata.Get(grpcGatewayUserAgentHeader); len(userAgents) > 0 {
			md.UserAgent = userAgents[0]
		}

		if userAgents := metadata.Get(userAgentHeader); len(userAgents) > 0 {
			md.UserAgent = userAgents[0]
		}

		if clientIPs := metadata.Get(xForwardedForHeader); len(clientIPs) > 0 {
			md.ClientIP = clientIPs[0]
		}
	}

	// fetch client IP from peer for gRPC request
	if p, ok := peer.FromContext(ctx); ok {
		md.ClientIP = p.Addr.String()
	}

	return md
}
