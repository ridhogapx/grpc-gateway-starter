protoc -I . \
    --go_out ./server/ --go_opt paths=source_relative \
    --go-grpc_out ./server/ --go-grpc_opt paths=source_relative \
    proto/*.proto

protoc -I . --grpc-gateway_out ./server/ \
    --grpc-gateway_opt paths=source_relative \
    proto/*.proto