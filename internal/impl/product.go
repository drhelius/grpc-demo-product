package impl

import (
	"context"
	"log"

	"github.com/drhelius/grpc-demo-proto/product"
)

type Server struct {
	product.UnimplementedProductServiceServer
}

func (s *Server) Create(ctx context.Context, in *product.CreateProductReq) (*product.CreateProductResp, error) {

	log.Printf("[Product] Received: %s", in.GetProduct())

	return &product.CreateProductResp{Id: "testid"}, nil
}

func (s *Server) Read(ctx context.Context, in *product.ReadProductReq) (*product.ReadProductResp, error) {

	log.Printf("[Product] Received: %v", in.GetId())

	return &product.ReadProductResp{Product: &product.Product{Id: "demoid", Name: "demoname", Description: "demodesc", Price: 100}}, nil
}
