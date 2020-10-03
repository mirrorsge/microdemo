.PHONY: proto
proto:
	protoc -I ./proto \
	-I /home/jianjunli/Document/go/pkg/mod/github.com/googleapis/googleapis@v0.0.0-20201002000029-7e837ae45f7b \
	-I /home/jianjunli/Document/go/pkg/mod/github.com/mwitkow/go-proto-validators@v0.3.2 \
	--grpc-gateway_out ./proto \
	--grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
	--go_out ./proto \
	--go_opt paths=source_relative \
	--go-grpc_out ./proto \
	--go-grpc_opt paths=source_relative \
	--govalidators_out ./proto \
	--govalidators_opt paths=source_relative \
	./proto/*.proto
	
	