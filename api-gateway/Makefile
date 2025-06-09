generate-proto:
	buf genereate


GQLGEN_VERSION = v0.17.73
generate-gql:
	rm -rf model/graph/gen/*.go && \
	go get github.com/99designs/gqlgen@v0.17.73 && \
	go get github.com/99designs/gqlgen/codegen/config@v0.17.73 && \
	go get github.com/99designs/gqlgen/internal/imports@v0.17.73 && \
	go get github.com/99designs/gqlgen/api@v0.17.73 && \
	go get github.com/99designs/gqlgen@v0.17.73 && \
	go run github.com/99designs/gqlgen generate && \
	go mod tidy


test:
	@go clean -testcache && go test -v -race -short -failfast -cover ./...

clean:
	rm -r model/pb/*