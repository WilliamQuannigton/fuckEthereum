package task1

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func QueryBlock(client *ethclient.Client, blockNumber *uint64) (*types.Block, error) {
	var blockNum *big.Int
	if blockNumber == nil {
		blockNum = nil
	} else {
		blockNum = big.NewInt(int64(*blockNumber))
	}
	block, err := client.BlockByNumber(context.Background(), blockNum)
	if err != nil {
		return nil, err
	}
	return block, nil
}

// TransferETH performs ETH transfer using secure keystore
func TransferETH() error {
	fmt.Println("🚀 开始执行 ETH 转账...")

	// 转账参数
	keystorePath := "./credentials"
	keystoreFile := "UTC--2025-08-19T04-11-33.145529000Z--ed2026d04ed4c5ae27d4b460b72030054f85d86e"
	toAddress := "0x5691ab974191673eFe1ce2090f2404b26E2f7D9d"
	amount := big.NewInt(10000000000000000) // 0.01 ETH
	rpcURL := "https://eth-sepolia.g.alchemy.com/v2/vG2GE3gIGxYnxU5KzF3kN79qhQAvl2mS"

	fmt.Printf("📁 Keystore 路径: %s\n", keystorePath)
	fmt.Printf("📄 Keystore 文件: %s\n", keystoreFile)
	fmt.Printf("📍 接收地址: %s\n", toAddress)
	fmt.Printf("💰 转账金额: %s wei (%.8f ETH)\n", amount.String(), new(big.Float).Quo(new(big.Float).SetInt(amount), big.NewFloat(1e18)))
	fmt.Printf("🌐 RPC URL: %s\n", rpcURL)

	err := TransferETHWithSecureKeystore(
		keystorePath,
		keystoreFile,
		toAddress,
		amount,
		rpcURL,
	)

	if err != nil {
		fmt.Printf("❌ 转账失败: %v\n", err)
		return err
	}

	fmt.Println("✅ 转账执行完成！")
	return nil
}
