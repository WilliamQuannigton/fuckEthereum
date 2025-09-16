# Abigen 工作流程详解

## 输入：ABI + 字节码

### 1. ABI 文件 (Counter.abi)
```json
[
  {
    "inputs": [],
    "name": "increment",
    "outputs": [{"internalType": "uint256", "name": "", "type": "uint256"}],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "getCount",
    "outputs": [{"internalType": "uint256", "name": "", "type": "uint256"}],
    "stateMutability": "view",
    "type": "function"
  }
]
```

### 2. 字节码文件 (Counter.bin)
```
0x6080604052348015600e575f5ffd5b505f5f819055506103aa806100225f395ff3fe...
```

## Abigen 的处理过程

### 步骤 1：解析 ABI
```go
// abigen 解析 ABI，提取函数信息
type Function struct {
    Name    string
    Inputs  []Input
    Outputs []Output
    Type    string
}

// 生成函数签名
func (c *Counter) Increment(opts *bind.TransactOpts) (*types.Transaction, error)
func (c *Counter) GetCount(opts *bind.CallOpts) (*big.Int, error)
```

### 步骤 2：生成类型安全的绑定
```go
// 自动生成类型安全的参数和返回值
func (c *Counter) Increment(opts *bind.TransactOpts) (*types.Transaction, error) {
    // 自动处理：
    // 1. 函数选择器编码
    // 2. 参数编码
    // 3. 交易构建
    // 4. 签名和发送
}
```

### 步骤 3：处理复杂的数据类型
```go
// 自动处理 big.Int 类型
func (c *Counter) GetCount(opts *bind.CallOpts) (*big.Int, error) {
    // 自动处理：
    // 1. 函数调用编码
    // 2. 返回值解码
    // 3. 类型转换
}
```

## 输出：完整的 Go 绑定代码

### 生成的代码包含：
1. **合约结构体**：`type Counter struct`
2. **函数绑定**：每个合约函数对应一个 Go 方法
3. **事件处理**：自动生成事件监听代码
4. **类型转换**：自动处理 Solidity 和 Go 类型转换
5. **错误处理**：统一的错误处理机制
6. **交易管理**：自动处理 nonce、gas 等
