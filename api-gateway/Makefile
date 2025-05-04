generate-proto:
	buf genereate


GQLGEN_VERSION = v0.17.73
generate-gql:
	go get github.com/99designs/gqlgen@$(GQLGEN_VERSION) && \
	go get github.com/99designs/gqlgen/codegen/config@$(GQLGEN_VERSION) && \
	go get github.com/99designs/gqlgen/internal/imports@$(GQLGEN_VERSION) && \
	go get github.com/99designs/gqlgen/api@$(GQLGEN_VERSION) && \
	go get github.com/99designs/gqlgen@$(GQLGEN_VERSION) && \
	go run github.com/99designs/gqlgen generate && \
	go mod tidy


test:
	@go clean -testcache && go test -v -race -short -failfast -cover ./...

clean:
	rm -r model/pb/*