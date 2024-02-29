swag:
	swag init -g cmd/main.go

test:
	go test -v -cover ./test/...

run:
	go run ./cmd/main.go

build:
	rm app
	go build -o app ./cmd/main.go

.PHONY: swag test run build