package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"

	proxyv1 "github.com/crypto-knight/go-rpc-proxy/gen/proxy/v1"
	"github.com/crypto-knight/go-rpc-proxy/internal/config"
	"github.com/crypto-knight/go-rpc-proxy/internal/proxy"
)

func main() {
	cfg := config.Load()

	lis, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proxyv1.RegisterProxyServiceServer(grpcServer, proxy.NewServer(cfg))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("proxy listening on :%s", cfg.Port)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	<-quit
	log.Println("Shutting down...")
	grpcServer.GracefulStop()
	log.Println("stopping")
}
