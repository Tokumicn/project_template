.PHONY: echo build protoc-sample protoc bindata-swag


echo:
	@echo "grpc-tlp"

build:
	go mod tidy
	go build -ldflags "-X main.version=0.0.1" -o grpc_tlp_svc ./cmd/main.go
	go build -ldflags "-X main.version=0.0.1" -o grpc_tlp_cli ./cli/main.go

protoc-sample:
	protoc --go_out=plugins=grpc:. ./proto/*.proto

# 为使用 google/api/annotations.proto 因此需要额外引入 googleapis
# --proto_path 或者 -I
# --grpc-gateway_out=logtostderr=true:.             网关反向代理支持
# --go_out=paths=source_relative:./proto            go代码  文件写到 proto 文件夹中
# --go-http_out=paths=source_relative:./proto       http代码
# --go-grpc_out=paths=source_relative:./proto       grpc代码
# --swagger_out=logtostderr=true:./proto            swagger文档
protoc:
	protoc --proto_path=./proto \
		   --proto_path=./third_party \
		   --grpc-gateway_out=logtostderr=true:. \
		   --go_out=paths=source_relative:./proto \
           --go-http_out=paths=source_relative:./proto \
           --go-grpc_out=paths=source_relative:./proto \
           --swagger_out=logtostderr=true:./proto \
		   ./proto/*.proto

# 将 swagger-ui 生成为静态文件供服务使用
# 源文件在 third_party/swagger-ui/  生成到 pkg/swagger/data.go 中
bindata-swag:
	go-bindata --nocompress -pkg swagger -o pkg/swagger/data.go third_party/swagger-ui/...

