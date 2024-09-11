package main

import (
	"bufio"
	"fmt"
	agent "lego_lib/gops_agent"
	"lego_lib/utils/copyist"
	"log"
	"os"
)

func main() {
	// 启动gops angent
	agent.StartGopsAgent("")

	// 用户与控制台交互的过程中产生的标准输出都会写入到文件中保存
	copyist.RecordEverything("", func() {
		// 用户交互
		log.Println("开始输入内容，输入 'exit' 退出程序。")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
			if scanner.Text() == "exit" {
				break
			}
		}
	})
}
