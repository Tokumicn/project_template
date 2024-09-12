# .PHONY: build clean tool lint help
# 声明为伪目标后：在执行对应的命令时，make 就不会去检查是否存在
# build / clean / tool / lint / help 其对应的文件，而是每次都会运行标签对应的命令
# 若不声明：恰好存在对应的文件，则 make 将会认为 xx 文件已存在，没有重新构建的必要了
.PHONY: check-build check-build-clear

# 通过 go build main.go 检查修改是否造成了不良影响
check-build:
	@echo "检查服务改动是否影响正常编译..."
	@cd grpc_tlp && $(MAKE) build
	@cd gin_tlp && $(MAKE) build
	@cd cobra_tlp && $(MAKE) build
	@cd cmd_tlp && $(MAKE) build

# 清理 check-build 产生的二进制文件
check-build-clear:
	@echo "清理检查过程中编译产生的二进制文件..."
	@cd grpc_tlp && rm grpc_tlp_svc && rm grpc_tlp_cli
	@cd gin_tlp && rm gin_tlp
	@cd cobra_tlp && rm cobra_tlp
	@cd cmd_tlp && rm cmd_tlp