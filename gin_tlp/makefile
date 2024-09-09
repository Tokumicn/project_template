

tidy:
	go mod tidy

build:
	go build -o gin_tlp  main.go

swagger:
	swag init


# 将 config.yaml 生成为静态文件供服务使用
# 源文件在 configs/config.yaml  生成到 configs/config.go 中
bindata-config:
	go-bindata --nocompress -pkg configs -o configs/config.go configs/...