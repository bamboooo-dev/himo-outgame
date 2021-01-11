default: build

build:
	go build -o bin/outgame cmd/outgame/main.go

gen-grpc:
	protoc \
		--go_out ./pkg \
		--go_opt paths=source_relative \
		--go-grpc_out ./pkg \
		--go-grpc_opt paths=source_relative \
		grpc/*/*/proto/*.proto

install-grpc:
	go install \
	google.golang.org/grpc/cmd/protoc-gen-go-grpc \
	google.golang.org/protobuf/cmd/protoc-gen-go
