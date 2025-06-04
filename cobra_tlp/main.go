package main

import (
	"github.com/Tokumicn/cobra_tlp/cmd"
	agent "lego_lib/gops_agent"
	"log"
)

func main() {
	// 启动 gops agent
	agent.StartGopsAgent("")

	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
