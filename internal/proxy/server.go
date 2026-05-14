package proxy

import (
	"context"
	"fmt"

	proxyv1 "github.com/crypto-knight/go-rpc-proxy/gen/proxy/v1"
	"github.com/crypto-knight/go-rpc-proxy/internal/config"
)

type Server struct {
	proxyv1.UnimplementedProxyServiceServer
	cfg *config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{cfg: cfg}
}

func (s *Server) Forward(ctx context.Context, req *proxyv1.ForwardRequest) (*proxyv1.ForwardResponse, error) {
	addr, ok := s.cfg.Services[req.TargetService]

	if !ok {
		return &proxyv1.ForwardResponse{
			Status: 404,
			ErrorMsg: fmt.Sprintf("unknown service: %s", req.TargetService),
		}, nil

	}
	// TODO: real forwarding logic
	return &proxyv1.ForwardResponse {
		Payload: []byte(fmt.Sprintf("proxied to %s → %s::%s", req.TargetService, addr, req.Method)),
		Status: 200,
	}, nil
}