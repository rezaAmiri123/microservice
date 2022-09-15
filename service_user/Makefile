test_coverage:
	go test --cover ./...

user_proto:
	protoc \
		--go_out=proto/grpc --go_opt=paths=source_relative \
		--go-grpc_out=proto/grpc --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=proto/grpc --grpc-gateway_opt=paths=source_relative \
        --openapiv2_out=docs/swagger --openapiv2_opt=allow_merge=true,merge_file_name=service_user \
		--proto_path=proto/grpc proto/grpc/*.proto