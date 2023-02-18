install-tools:
	echo installing tools
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/bufbuild/buf/cmd/buf@latest
	go install github.com/vektra/mockery/v2@latest
	go install github.com/go-swagger/go-swagger/cmd/swagger@latest
	go install github.com/cucumber/godog/cmd/godog@latest
	echo done

generate:
	echo running code generation
	go generate ./...
	echo done

evans:
	evans --host localhost --port 5080 -r repl
	# evans --host 192.168.39.213 --port 30000 -r repl

evans_tls:
	evans --host localhost --port 5080 -t --servername localhost --cacert tls/ca.pem --cert tls/client.pem --certkey tls/client-key.pem -r repl

# set evanse header
# header Authorization='Bearer xxxxx'
