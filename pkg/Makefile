install-tools:
	echo installing tools
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/vektra/mockery/v2@latest
	echo done


generate:
	go generate ./...

proto:
	protoc \
		--go_out=am --go_opt=paths=source_relative \
		--go-grpc_out=am --go-grpc_opt=paths=source_relative \
		--proto_path=am am/*.proto
	protoc \
		--go_out=jetstream --go_opt=paths=source_relative \
		--go-grpc_out=jetstream --go-grpc_opt=paths=source_relative \
		--proto_path=jetstream jetstream/*.proto
	protoc \
		--go_out=kafka/kafkastream --go_opt=paths=source_relative \
		--go-grpc_out=kafka/kafkastream --go-grpc_opt=paths=source_relative \
		--proto_path=kafka/kafkastream kafka/kafkastream/*.proto
