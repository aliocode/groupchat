generate:
	mkdir -p ./generated/groupchatpb
	protoc --proto_path=api/groupchat \
			--go_out=generated/groupchatpb --go_opt=paths=source_relative \
  			--go-grpc_out=generated/groupchatpb --go-grpc_opt=paths=source_relative \
   			groupchat.proto

run-server:
	go run ./cmd/server/main.go

run-client:
	go run ./cmd/client/main.go

deps:
	brew install golangci-lint
	brew install protobuf
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

lint:
	golangci-lint run --out-format progress --config golangci.yaml

run:
	go run ./cmd/server/main.go
	go run ./cmd/client/main.go