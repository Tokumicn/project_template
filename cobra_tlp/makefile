.PHONY: echo tidy build run test-word test-time test-sql

echo:
	@echo "cobra-tlp"

tidy:
	go mod tidy

build:
	go mod tidy
	go build -ldflags "-X main.version=0.0.1" -o cobra_tlp main.go

run:
	go run main.go help word

# 测试 word: 单词格式转换
test-word:
	go run main.go word -s=tokumi -m=1 # 输出结果: TOKUMI
	go run main.go word -s=TOKUMI -m=2 # 输出结果: tokumi
	go run main.go word -s=tokumi -m=3 # 输出结果: Tokumi
	go run main.go word -s=TOKUMI -m=4 # 输出结果: tOKUMI
	go run main.go word -s=ToKuMi -m=5 # 输出结果: to_ku_mi


# 测试 time: 时间计算工具
test-time:
	go run main.go time now                                  # 输出结果: 以当前时间为准
	go run main.go time calc -c="2024-09-09 12:02:33" -d=5m  # 输出结果: 2024-09-09 12:07:33, 1725883653
	go run main.go time calc -c="2024-09-09 12:02:33" -d=-2h # 输出结果: 2024-09-09 10:02:33, 1725876153


# 测试 sql: 数据库表结构和Go结构体转换工具
test-sql:
	go run main.go sql sql2struct -u 'root' -p 'rootroot' --db 'GinTlpService' -t 'blog_auth'
	# 全名使用双横杠 "--db"
	# 短名使用单横杠 "-p"     -p == --password