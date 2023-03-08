.PHONY: start

start-parser:
	(cd ./cmd/parser && go run .)

start-gateway:
	(cd ./cmd/gateway && go run .)

lint:
	golangci-lint run

deploy-parser:
#	(go build ./cmd/parser && scp main user@192.168.0.1:~/lft-backend)
	(go build ./cmd/parser)

deploy-gateway:
#	(go build ./cmd/parser && scp main user@192.168.0.1:~/lft-backend)
	(go build ./cmd/gateway)

deploy-parser-env:
	scp ./cmd/parser/.env user@192.168.0.1:~/lft-backend

deploy-gateway-env:
	scp ./cmd/gateway/.env user@192.168.0.1:~/lft-backend