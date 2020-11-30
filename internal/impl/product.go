package impl

import (
	"context"
	"log"
	"strconv"

	"github.com/Pallinder/go-randomdata"
	"github.com/drhelius/grpc-demo-proto/product"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	product.UnimplementedProductServiceServer
}

func (s *Server) Create(ctx context.Context, in *product.CreateProductReq) (*product.CreateProductResp, error) {

	log.Printf("[Product] Create Req: %v", in.GetProduct())

	r := &product.CreateProductResp{Id: strconv.Itoa(randomdata.Number(1000000))}

	err := failedContext(ctx)
	if err != nil {
		return nil, err
	}

	log.Printf("[Product] Create Res: %v", r.GetId())

	return r, nil
}

func (s *Server) Read(ctx context.Context, in *product.ReadProductReq) (*product.ReadProductResp, error) {

	log.Printf("[Product] Read Req: %v", in.GetId())

	r := &product.ReadProductResp{Product: &product.Product{Id: in.GetId(), Name: randomdata.SillyName(), Description: randomdata.Paragraph(), Price: int32(randomdata.Number(1000))}}

	err := failedContext(ctx)
	if err != nil {
		return nil, err
	}

	log.Printf("[Product] Read Res: %v", r.GetProduct())

	return r, nil
}

func failedContext(ctx context.Context) error {
	if ctx.Err() == context.Canceled {
		log.Printf("[Order] context canceled, stoping server side operation")
		return status.Error(codes.Canceled, "context canceled, stoping server side operation")
	}

	if ctx.Err() == context.DeadlineExceeded {
		log.Printf("[Order] dealine has exceeded, stoping server side operation")
		return status.Error(codes.DeadlineExceeded, "dealine has exceeded, stoping server side operation")
	}

	return nil
}
