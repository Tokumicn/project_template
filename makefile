
# 通过 go build main.go 检查修改是否造成了不良影响
check-build:
	@cd grpc_tlp && $(MAKE) build
	@cd gin_tlp && $(MAKE) build
	@cd cobra_tlp && $(MAKE) build
	@cd cmd_tlp && $(MAKE) build

# 清理 check-build 产生的二进制文件
check-build-clear:
	@cd grpc_tlp && rm grpc_tlp_svc && rm grpc_tlp_cli
	@cd gin_tlp && rm gin_tlp
	@cd cobra_tlp && rm cobra_tlp
	@cd cmd_tlp && rm cmd_tlp