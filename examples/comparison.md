# 传统方式 vs Abigen 方式对比

## 传统方式（手动处理）

```go
// 需要手动解析 ABI
abi, err := abi.JSON(strings.NewReader(contractABI))

// 手动编码函数调用
data, err := abi.Pack("increment")

// 手动构建交易
tx := types.NewTransaction(nonce, contractAddr, value, gasLimit, gasPrice, data)

// 手动发送交易
signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
err = client.SendTransaction(ctx, signedTx)

// 手动解析返回值
var result []interface{}
err = abi.Unpack(&result, "increment", returnData)
```

**问题**：
- ❌ 代码冗长，容易出错
- ❌ 需要手动处理 ABI 解析
- ❌ 类型不安全，运行时才发现错误
- ❌ 维护困难，合约更新需要手动修改

## Abigen 方式（自动生成）

```go
// 直接调用合约方法
count, err := contract.GetCount(nil)

// 直接调用写入方法
tx, err := contract.Increment(auth)

// 类型安全，编译时检查
newCount, err := contract.Increment(auth)
```

**优势**：
- ✅ 代码简洁，易于理解
- ✅ 类型安全，编译时检查
- ✅ 自动处理 ABI 解析
- ✅ 合约更新后重新生成即可
- ✅ 支持事件监听
- ✅ 自动处理 gas 估算
