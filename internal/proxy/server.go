package proxy

import (
	"context"
	"fmt"

	proxyv1 "github.com/crypto-knight/go-rpc-proxy/gen/proxy/v1"
)

type Server struct {
	proxyv1.UnimplementedProxyServiceServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Forward(ctx context.Context, req *proxyv1.ForwardRequest) (*proxyv1.ForwardResponse, error) {
	// TODO: real forwarding logic
	return &proxyv1.ForwardResponse {
		Payload: []byte(fmt.Sprintf("proxied: %s → %s", req.TargetService, req.Method)),
		Status: 200,
	}, nil
}