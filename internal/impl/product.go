package impl

import (
	"context"
	"log"

	pb "github.com/drhelius/grpc-demo-proto/product"
)

type Server struct {
	pb.UnimplementedProductServiceServer
}

func (s *Server) Create(ctx context.Context, in *pb.CreateProductReq) (*pb.CreateProductResp, error) {

	log.Printf("Received: %s", in.GetProduct())

	return &pb.CreateProductResp{Id: "testid"}, nil
}

func (s *Server) Read(ctx context.Context, in *pb.ReadProductReq) (*pb.ReadProductResp, error) {

	log.Printf("Received: %v", in.GetId())

	return &pb.ReadProductResp{Product: &pb.Product{Id: "demoid", Name: "demoname", Description: "demodesc", Price: 100}}, nil
}
