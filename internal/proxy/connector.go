package proxy

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/crypto-knight/go-rpc-proxy/internal/config"
)

type Connector struct {
	conns map[string]*grpc.ClientConn
}

func NewConnector(cfg config.Config) (*Connector, error) {
	conns := make(map[string]*grpc.ClientConn)

	for name, addr := range cfg.Services {
		conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return nil, fmt.Errorf("connect to %s {%s}: %s", name, addr, err)
		}
		conns[name] = conn
	}

	return &Connector{conns}, nil
}

func (c *Connector) Get(service string) (*grpc.ClientConn, bool) {
	conn, ok := c.conns[service]
	return conn, ok
}

func (c *Connector) Close() {
	for _, conn := range c.conns {
		conn.Close()
	}
}
