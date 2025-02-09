help:
	@echo "HELP: make clean|gqlgen|build|run"
clean:
	@echo "-cleaning-"
	rm -rfv cmd/product-gql/binary
	rm -rfv cmd/tenant-gql/binary
gqlgen:
	@echo "-generate-"
	cd tenant; go run github.com/99designs/gqlgen generate; cd ..
	cd product; go run github.com/99designs/gqlgen generate; cd ..
build:
	@echo "-building-"
	go build -o cmd/tenant-gql/binary cmd/tenant-gql/main.go
	go build -o cmd/product-gql/binary cmd/product-gql/main.go
dev:
	@echo "-running-"
	@echo "./cmd/tenant-gql/binary"
	@echo "./cmd/product-gql/binary"
	@echo "node gateway/index.js"