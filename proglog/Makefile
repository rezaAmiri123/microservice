install-tools:
	echo installing tools
#	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
#	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
#	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
#	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
#	go install github.com/bufbuild/buf/cmd/buf@latest
	go install github.com/vektra/mockery/v2@latest
	#go install github.com/go-swagger/go-swagger/cmd/swagger@latest
	echo done

generate:
	mockery --quiet --dir ./internal -r --all --inpackage --case underscore
	mockery --quiet --dir ./api -r --all --inpackage --case underscore
	#swagger generate client -q -f ./internal/ports/rest/api.swagger.json -c storesclient -m storesclient/models --with-flatten=remove-unused


compile:
	protoc internal/domain/*.proto \
		--go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		--proto_path=.
	protoc api/v1/*.proto \
    		--go_out=. --go_opt=paths=source_relative \
    		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
    		--proto_path=.


