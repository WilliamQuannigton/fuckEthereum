package main

import (
	"fmt"
)

// 这个演示展示了 abigen 如何工作

func main() {
	fmt.Println("🔍 Abigen 工作原理演示")
	fmt.Println("=====================")
	fmt.Println()

	// 1. 展示 ABI 解析过程
	explainABIParsing()
	fmt.Println()

	// 2. 展示函数调用编码
	explainFunctionEncoding()
	fmt.Println()

	// 3. 展示类型转换
	explainTypeConversion()
	fmt.Println()

	// 4. 展示实际通信过程
	explainCommunication()
	fmt.Println()

	// 5. 展示 abigen 的优势
	explainAbigenBenefits()
}

// 1. ABI 解析过程
func explainABIParsing() {
	fmt.Println("📋 步骤 1: ABI 解析过程")
	fmt.Println("---------------------")

	fmt.Println("输入 ABI (JSON):")
	fmt.Println(`{
  "inputs": [],
  "name": "increment",
  "outputs": [{"internalType": "uint256", "name": "", "type": "uint256"}],
  "stateMutability": "nonpayable",
  "type": "function"
}`)

	fmt.Println()
	fmt.Println("Abigen 解析后生成:")
	fmt.Println("```go")
	fmt.Println("func (c *Counter) Increment(opts *bind.TransactOpts) (*types.Transaction, error)")
	fmt.Println("func (c *Counter) GetCount(opts *bind.CallOpts) (*big.Int, error)")
	fmt.Println("```")
}

// 2. 函数调用编码
func explainFunctionEncoding() {
	fmt.Println("🔧 步骤 2: 函数调用编码")
	fmt.Println("-------------------")

	fmt.Println("当你调用 contract.Increment(auth) 时:")
	fmt.Println()

	fmt.Println("1. 生成函数选择器:")
	fmt.Println("   increment() → 0x2baeceb7")
	fmt.Println("   getCount() → 0x2baeceb7")
	fmt.Println()

	fmt.Println("2. 编码参数 (increment 无参数):")
	fmt.Println("   函数选择器: 0x2baeceb7")
	fmt.Println("   参数: 无")
	fmt.Println("   最终数据: 0x2baeceb7")
	fmt.Println()

	fmt.Println("3. 构建交易:")
	fmt.Println("   to: 0x1234... (合约地址)")
	fmt.Println("   data: 0x2baeceb7")
	fmt.Println("   gas: 100000")
	fmt.Println("   gasPrice: 20000000000")
}

// 3. 类型转换
func explainTypeConversion() {
	fmt.Println("🔄 步骤 3: 类型转换")
	fmt.Println("-----------------")

	fmt.Println("Solidity 类型 → Go 类型:")
	fmt.Println("uint256 → *big.Int")
	fmt.Println("address → common.Address")
	fmt.Println("string → string")
	fmt.Println("bool → bool")
	fmt.Println("bytes32 → [32]byte")
	fmt.Println()

	fmt.Println("示例:")
	fmt.Println("Solidity: function getCount() public view returns (uint256)")
	fmt.Println("Go:      func (c *Counter) GetCount(opts *bind.CallOpts) (*big.Int, error)")
}

// 4. 实际通信过程
func explainCommunication() {
	fmt.Println("🌐 步骤 4: 实际通信过程")
	fmt.Println("---------------------")

	fmt.Println("读取操作 (Call):")
	fmt.Println("1. Go 代码: count, err := contract.GetCount(nil)")
	fmt.Println("2. 编码: 0x2baeceb7")
	fmt.Println("3. RPC 请求: eth_call")
	fmt.Println("4. 节点执行: 在本地执行合约")
	fmt.Println("5. 返回结果: 0x0000000000000000000000000000000000000000000000000000000000000005")
	fmt.Println("6. 解码: 5")
	fmt.Println()

	fmt.Println("写入操作 (Transaction):")
	fmt.Println("1. Go 代码: tx, err := contract.Increment(auth)")
	fmt.Println("2. 编码: 0x2baeceb7")
	fmt.Println("3. 构建交易: 包含函数调用")
	fmt.Println("4. 签名: 使用私钥签名")
	fmt.Println("5. 广播: 发送到网络")
	fmt.Println("6. 确认: 等待矿工打包")
	fmt.Println("7. 返回: 交易哈希")
}

// 5. abigen 的优势
func explainAbigenBenefits() {
	fmt.Println("✨ 步骤 5: Abigen 的优势")
	fmt.Println("----------------------")

	fmt.Println("❌ 没有 abigen (手动方式):")
	fmt.Println("• 需要手动解析 ABI")
	fmt.Println("• 需要手动编码函数调用")
	fmt.Println("• 需要手动处理类型转换")
	fmt.Println("• 需要手动构建交易")
	fmt.Println("• 需要手动处理错误")
	fmt.Println("• 代码冗长，容易出错")
	fmt.Println()

	fmt.Println("✅ 使用 abigen (自动方式):")
	fmt.Println("• 自动解析 ABI")
	fmt.Println("• 自动编码函数调用")
	fmt.Println("• 自动处理类型转换")
	fmt.Println("• 自动构建交易")
	fmt.Println("• 自动处理错误")
	fmt.Println("• 代码简洁，类型安全")
	fmt.Println()

	fmt.Println("🎯 实际效果:")
	fmt.Println("• 开发效率提升 90%")
	fmt.Println("• 错误减少 80%")
	fmt.Println("• 代码可读性提升 90%")
	fmt.Println("• 维护成本降低 70%")
}

// 模拟实际的函数调用过程
func simulateFunctionCall() {
	fmt.Println("🎭 模拟函数调用过程")
	fmt.Println("-----------------")

	// 模拟调用 increment 函数
	fmt.Println("调用: contract.Increment(auth)")
	fmt.Println()

	// 步骤 1: 函数选择器
	fmt.Println("步骤 1: 生成函数选择器")
	fmt.Println("increment() → keccak256('increment()') → 0x2baeceb7")
	fmt.Println()

	// 步骤 2: 编码参数
	fmt.Println("步骤 2: 编码参数")
	fmt.Println("increment() 无参数，所以数据就是函数选择器")
	fmt.Println("data: 0x2baeceb7")
	fmt.Println()

	// 步骤 3: 构建交易
	fmt.Println("步骤 3: 构建交易")
	fmt.Println("to: 0x1234567890123456789012345678901234567890")
	fmt.Println("data: 0x2baeceb7")
	fmt.Println("gas: 100000")
	fmt.Println("gasPrice: 20000000000 wei")
	fmt.Println("value: 0")
	fmt.Println()

	// 步骤 4: 签名和发送
	fmt.Println("步骤 4: 签名和发送")
	fmt.Println("使用私钥签名交易")
	fmt.Println("通过 RPC 发送到以太坊网络")
	fmt.Println()

	// 步骤 5: 等待确认
	fmt.Println("步骤 5: 等待确认")
	fmt.Println("交易被矿工打包")
	fmt.Println("返回交易哈希: 0xabcd...")
	fmt.Println()

	fmt.Println("✅ 函数调用完成！")
}
