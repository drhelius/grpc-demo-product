package grpc

import (
	"log"
	"net"
	"sync"

	"github.com/drhelius/grpc-demo-product/internal/impl"
	"github.com/drhelius/grpc-demo-proto/product"
	"google.golang.org/grpc"
)

func Serve(wg *sync.WaitGroup, port string) {
	defer wg.Done()

	lis, err := net.Listen("tcp", ":"+port)

	if err != nil {
		log.Fatalf("[Product] GRPC failed to listen: %v", err)
	}

	s := grpc.NewServer()

	product.RegisterProductServiceServer(s, &impl.Server{})

	log.Printf("[Product] Serving GRPC on localhost:%s ...", port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("[Product] GRPC failed to serve: %v", err)
	}
}
