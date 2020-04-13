BINARY_PATH=bin

CMD_MAIN_PATH=cmd/grpc-demo-product/main.go
CMD_BINARY_PATH=$(BINARY_PATH)/grpc-demo-product
CMD_SERVER_PORT=5000

PROTO_FILES_PATH=../grpc-demo-proto
PROTO_FILE=product/product.proto
PROTO_OUT_PATH=internal/grpc

GOOGLE_APIS=$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis

build:
	go build -o $(CMD_BINARY_PATH) -race $(CMD_MAIN_PATH)

clean:
	go clean $(CMD_MAIN_PATH)
	rm -f $(BINARY_PATH)/*

build-proto:
	mkdir -p $(PROTO_OUT_PATH)
	protoc -I $(PROTO_FILES_PATH) -I $(GOOGLE_APIS) --go_out=plugins=grpc:$(PROTO_OUT_PATH) $(PROTO_FILE)
	protoc -I $(PROTO_FILES_PATH) -I $(GOOGLE_APIS) --grpc-gateway_out=logtostderr=true:$(PROTO_OUT_PATH) $(PROTO_FILE)

clean-proto:
	rm -f $(PROTO_OUT_PATH)/*.pb.go

run:
	$(CMD_BINARY_PATH)

grpcui:
	grpcui -plaintext -import-path $(PROTO_FILES_PATH) -import-path $(GOOGLE_APIS) -proto $(PROTO_FILE) localhost:$(CMD_SERVER_PORT)
