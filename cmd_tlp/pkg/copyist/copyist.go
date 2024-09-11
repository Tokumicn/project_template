package copyist

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

var recordFilePath = "output.txt"

// RecordEverything 记录开始
//
// 参数 filePath： 不填默认 ./output.txt
//
// 参数 fn： 用户交互逻辑
func RecordEverything(filePath string, fn func()) {

	if len(filePath) > 0 {
		recordFilePath = filePath
	}

	// 创建一个同步等待组
	var wg sync.WaitGroup

	// 创建一个管道，用于同步输出
	r, w, err := os.Pipe()
	if err != nil {
		log.Fatalf("error creating pipe: %v", err)
		return
	}
	defer r.Close()

	// 创建一个 goroutine 来从管道读取数据并写入文件
	wg.Add(1)
	go func() {
		defer wg.Done() // 该方法退出时，结束记录工作

		// 打开文件
		file, err := os.OpenFile(recordFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
			return
		}
		defer file.Close()

		// 从管道读取数据并写入文件
		// 程序会在这里永久等待EOF，因此标准输入的所有内容都会写入文件中
		_, err = io.Copy(file, r)
		if err != nil {
			log.Fatalf("error writing to file: %v", err)
			return
		}
	}()

	// 将标准输出重定向到管道写端[核心逻辑：劫持标准输出]
	stdoutOrig := os.Stdout
	os.Stdout = w

	fn() // 用户交互逻辑

	// 关闭管道写端，触发 goroutine 退出
	err = w.Close()
	if err != nil {
		log.Fatalf("error close pipe witer to file: %v", err)
		return
	}

	// 等待后台 goroutine 完成
	wg.Wait()

	// 恢复标准输出
	os.Stdout = stdoutOrig
	fmt.Println("程序结束，内容已写入文件")
}
