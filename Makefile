fmt:
	go fmt ./...

run:
	make fmt
	go run cmd/main.go

build:
	make fmt
	mkdir -p target
	go build -o "./target/dnd-template-creator" ./cmd/main.go
