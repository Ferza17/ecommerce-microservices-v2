generate:
	protoc \
      --proto_path=model/proto \
      --go_out=model/pb \
      --go_opt=paths=source_relative \
      --go-grpc_out=model/pb \
      --go-grpc_opt=paths=source_relative \
      --gorm_out=model/pb \
      --gorm_opt=paths=source_relative \
      model/proto/*.proto



test:
	@go clean -testcache && go test -v -race -short -failfast -cover ./...

clean:
	rm -r model/pb/*