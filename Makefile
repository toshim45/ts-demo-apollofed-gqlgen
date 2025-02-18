help:
	@echo "HELP: make clean|gqlgen|build|dev-start|dev-stop"
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
tmux-check:
	@which tmux > /dev/null
dev-start: tmux-check
	@echo "-running-"
	@tmux new -d -s tenant-gql-api ./cmd/tenant-gql/binary
	@tmux new -d -s product-gql-api ./cmd/product-gql/binary
	@tmux new -d -s gql-gateway 'node --env-file=gateway/.env gateway/index.js'

dev-status: tmux-check
	@echo "-status-"
	@tmux ls
	@lsof -nP -iTCP -sTCP:LISTEN

dev-stop: tmux-check
	@echo "-stopping-"
	@tmux kill-session -t tenant-gql-api
	@tmux kill-session -t product-gql-api
	@tmux kill-session -t gql-gateway