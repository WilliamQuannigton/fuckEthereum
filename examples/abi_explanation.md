# ABI 详解：智能合约的"接口说明书"

## 什么是 ABI？

ABI（Application Binary Interface）就像是智能合约的"接口说明书"，它告诉外部程序：
- 合约有哪些函数
- 每个函数需要什么参数
- 函数返回什么数据
- 合约有哪些事件

## ABI 示例

```json
{
  "inputs": [],
  "name": "increment",
  "outputs": [{"internalType": "uint256", "name": "", "type": "uint256"}],
  "stateMutability": "nonpayable",
  "type": "function"
}
```

这个 ABI 告诉我们：
- 函数名：`increment`
- 输入参数：无
- 输出：一个 uint256 类型的值
- 状态可变性：nonpayable（不接收 ETH）

## 为什么需要 ABI？

智能合约编译后变成字节码，就像这样：
```
0x6080604052348015600e575f5ffd5b505f5f819055506103aa806100225f395ff3fe...
```

这串字节码对人类来说完全不可读！ABI 就是"翻译器"，告诉我们：
- 如何调用 `increment()` 函数
- 如何解析返回的数据
- 如何监听合约事件
