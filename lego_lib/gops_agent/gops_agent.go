package lego_lib

import (
	"github.com/google/gops/agent"
	"log"
)

// StartGopsAgent 启动gops agent 方便性能分析
func StartGopsAgent(addr string) {
	if len(addr) <= 0 {
		addr = "127.0.0.1:6060"
	}
	// 创建并监听 gops agent，gops 命令会通过连接 agent 来读取进程信息
	// 若需要远程访问，可配置 agent.Options{Addr: "0.0.0.0:6060"}，否则默认仅允许本地访问
	if err := agent.Listen(agent.Options{Addr: addr}); err != nil {
		log.Fatalf("agent.Listen err: %v", err)
	}
}
