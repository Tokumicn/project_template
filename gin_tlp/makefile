

tidy:
	go mod tidy

build:
	go build -o gin_tlp  main.go

swagger:
	swag init


# 将 swagger-ui 生成为静态文件供服务使用
# 源文件在 third_party/swagger-ui/  生成到 pkg/swagger/data.go 中
bindata-swag:
	go-bindata --nocompress -pkg swagger -o pkg/swagger/data.go third_party/swagger-ui/...

bindata-config:
	go-bindata --nocompress -pkg configs -o configs/config.go configs/...