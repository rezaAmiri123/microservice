install-tools:
	echo installing tools
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	echo done


mock_gen:
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
