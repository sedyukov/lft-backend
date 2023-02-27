.PHONY: start

start-parser:
	(cd ./cmd/parser && go run .)

start-gateway:
	(cd ./cmd/gateway && go run .)

lint:
	golangci-lint run

deploy:
	(go build cmd/sync/main.go && scp main user@192.168.0.1:~/lft-backend)