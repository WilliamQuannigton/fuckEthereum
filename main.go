package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/fuckEthereum/src/task1"
	"github.com/fuckEthereum/src/task2"
)

func main() {
	// Check command line arguments
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "task1":
			runTask1()
		case "task2":
			runTask2()
		case "setup":
			runSetup()
		default:
			printUsage()
		}
	} else {
		printUsage()
	}
}

func runTask1() {
	fmt.Println("🚀 开始执行 Task 1: ETH 转账测试...")

	// Execute ETH transfer
	err := task1.TransferETH()
	if err != nil {
		log.Printf("转账失败: %v", err)
		return
	}

	fmt.Println("✅ Task 1 转账测试完成！")
}

func runTask2() {
	fmt.Println("🚀 开始执行 Task 2: Abigen 智能合约交互...")

	// Execute contract interaction demo
	err := task2.RunTask2()
	if err != nil {
		log.Printf("智能合约交互失败: %v", err)
		return
	}

	fmt.Println("✅ Task 2 智能合约交互完成！")
}

func runSetup() {
	fmt.Println("🚀 开始执行 Abigen 设置...")

	// Run the setup script
	cmd := exec.Command("./scripts/setup_abigen.sh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Printf("设置失败: %v", err)
		return
	}

	fmt.Println("✅ Abigen 设置完成！")
}

func printUsage() {
	fmt.Println("🚀 Ethereum Go 学习项目")
	fmt.Println("========================")
	fmt.Println("")
	fmt.Println("使用方法:")
	fmt.Println("  go run main.go task1    - 执行 ETH 转账测试")
	fmt.Println("  go run main.go task2    - 执行 Abigen 智能合约交互")
	fmt.Println("  go run main.go setup    - 设置 Abigen 环境")
	fmt.Println("")
	fmt.Println("Task 1: ETH 转账")
	fmt.Println("  - 演示基本的 ETH 转账功能")
	fmt.Println("  - 使用 keystore 文件管理私钥")
	fmt.Println("")
	fmt.Println("Task 2: Abigen 智能合约交互")
	fmt.Println("  - 使用 abigen 工具生成 Go 绑定代码")
	fmt.Println("  - 与 Sepolia 测试网络上的智能合约交互")
	fmt.Println("  - 演示合约部署、调用和状态查询")
	fmt.Println("")
	fmt.Println("设置:")
	fmt.Println("  - 自动安装所需工具 (solc, abigen)")
	fmt.Println("  - 编译智能合约")
	fmt.Println("  - 生成 Go 绑定代码")
}
