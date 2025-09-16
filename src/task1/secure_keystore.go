package task1

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
	"strings"

	"syscall"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/term"
)

// SecureKeystoreWallet represents a secure keystore-based wallet
type SecureKeystoreWallet struct {
	keystorePath string
	account      *accounts.Account
	// 注意：不存储密码，每次使用时临时输入
}

// NewSecureKeystoreWallet creates a new secure keystore wallet
func NewSecureKeystoreWallet(keystorePath string) *SecureKeystoreWallet {
	return &SecureKeystoreWallet{
		keystorePath: keystorePath,
	}
}

// promptPassword securely prompts for password without storing it
func promptPassword(prompt string) (string, error) {
	fmt.Print(prompt)
	password, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", fmt.Errorf("failed to read password: %v", err)
	}
	fmt.Println() // New line after password input
	return string(password), nil
}

// CreateAccount creates a new account in the keystore
func (kw *SecureKeystoreWallet) CreateAccount() error {
	// Create keystore directory if it doesn't exist
	if err := os.MkdirAll(kw.keystorePath, 0700); err != nil {
		return fmt.Errorf("failed to create keystore directory: %v", err)
	}

	// Prompt for password
	password, err := promptPassword("Enter password for new keystore: ")
	if err != nil {
		return err
	}

	// Confirm password
	confirmPassword, err := promptPassword("Confirm password: ")
	if err != nil {
		return err
	}

	if password != confirmPassword {
		return fmt.Errorf("passwords do not match")
	}

	// Prompt user to input their existing private key (hex string)
	fmt.Print("Enter your existing private key (hex, without 0x): ")
	var privKeyHex string
	_, err = fmt.Scanln(&privKeyHex)
	if err != nil {
		return fmt.Errorf("failed to read private key: %v", err)
	}

	// Decode the hex string to ECDSA private key
	privateKey, err := crypto.HexToECDSA(privKeyHex)
	if err != nil {
		return fmt.Errorf("invalid private key: %v", err)
	}

	// Create new keystore
	ks := keystore.NewKeyStore(kw.keystorePath, keystore.StandardScryptN, keystore.StandardScryptP)

	// Import the user-specified private key into the keystore
	account, err := ks.ImportECDSA(privateKey, password)
	if err != nil {
		return fmt.Errorf("failed to import private key into keystore: %v", err)
	}

	kw.account = &account

	// Print account information
	fmt.Printf("Account created successfully!\n")
	fmt.Printf("Address: %s\n", account.Address.Hex())
	fmt.Printf("Keystore file: %s\n", account.URL.Path)
	fmt.Println("⚠️  IMPORTANT: Remember your password! It cannot be recovered!")

	return nil
}

// ImportKeystore imports an existing keystore file
func (kw *SecureKeystoreWallet) ImportKeystore(keystoreFile string) error {
	// Read keystore file
	data, err := ioutil.ReadFile(keystoreFile)
	if err != nil {
		return fmt.Errorf("failed to read keystore file: %v", err)
	}

	// Parse keystore
	var keystoreData map[string]interface{}
	if err := json.Unmarshal(data, &keystoreData); err != nil {
		return fmt.Errorf("failed to parse keystore file: %v", err)
	}

	// Get address from keystore
	addressHex, ok := keystoreData["address"].(string)
	if !ok {
		return fmt.Errorf("invalid keystore file: missing address")
	}

	address := common.HexToAddress(addressHex)
	fmt.Printf("Imported keystore for address: %s\n", address.Hex())

	// Store account info (without password)
	kw.account = &accounts.Account{
		Address: address,
		URL:     accounts.URL{Path: keystoreFile},
	}

	return nil
}

// GetAddress returns the wallet address
func (kw *SecureKeystoreWallet) GetAddress() (common.Address, error) {
	if kw.account == nil {
		return common.Address{}, fmt.Errorf("no account loaded")
	}
	return kw.account.Address, nil
}

// SignTransaction signs a transaction using the keystore
func (kw *SecureKeystoreWallet) SignTransaction(tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
	if kw.account == nil {
		return nil, fmt.Errorf("no account loaded")
	}

	// Prompt for password each time (never store it)
	password, err := promptPassword("Enter password to sign transaction: ")
	if err != nil {
		return nil, err
	}

	// Load keystore
	ks := keystore.NewKeyStore(kw.keystorePath, keystore.StandardScryptN, keystore.StandardScryptP)

	// Unlock account with temporary password
	if err := ks.Unlock(*kw.account, password); err != nil {
		return nil, fmt.Errorf("failed to unlock account: %v", err)
	}

	// Sign transaction
	signedTx, err := ks.SignTx(*kw.account, tx, chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to sign transaction: %v", err)
	}

	// Lock account immediately after use
	ks.Lock(kw.account.Address)

	return signedTx, nil
}

// TransferETHWithSecureKeystore performs ETH transfer using secure keystore
func TransferETHWithSecureKeystore(
	keystorePath string,
	keystoreFile string,
	toAddress string,
	amount *big.Int,
	rpcURL string,
) error {
	fmt.Println("🔐 开始创建安全 Keystore 钱包...")

	// Create secure keystore wallet
	wallet := NewSecureKeystoreWallet(keystorePath)
	fmt.Printf("✅ 钱包实例创建成功，路径: %s\n", keystorePath)

	fmt.Println("📥 开始导入 Keystore 文件...")
	// Import keystore
	if err := wallet.ImportKeystore(keystorePath + "/" + keystoreFile); err != nil {
		fmt.Printf("❌ 导入 Keystore 失败: %v\n", err)
		return fmt.Errorf("failed to import keystore: %v", err)
	}
	fmt.Printf("✅ Keystore 文件导入成功: %s\n", keystoreFile)

	fmt.Println("🌐 开始连接以太坊网络...")
	// Connect to Ethereum client
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		fmt.Printf("❌ 连接以太坊网络失败: %v\n", err)
		return fmt.Errorf("failed to connect to Ethereum client: %v", err)
	}
	defer client.Close()
	fmt.Printf("✅ 成功连接到以太坊网络: %s\n", rpcURL)

	fmt.Println("📍 获取发送方地址...")
	// Get address from keystore
	fromAddress, err := wallet.GetAddress()
	if err != nil {
		fmt.Printf("❌ 获取发送方地址失败: %v\n", err)
		return fmt.Errorf("failed to get address: %v", err)
	}

	fmt.Printf("✅ 发送方地址: %s\n", fromAddress.Hex())

	fmt.Println("📍 解析接收方地址...")
	// Parse recipient address
	toAddr := common.HexToAddress(toAddress)
	fmt.Printf("✅ 接收方地址: %s\n", toAddr.Hex())

	fmt.Println("🔢 获取账户 Nonce...")
	// Get nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		fmt.Printf("❌ 获取 Nonce 失败: %v\n", err)
		return fmt.Errorf("failed to get nonce: %v", err)
	}
	fmt.Printf("✅ 账户 Nonce: %d\n", nonce)

	fmt.Println("⛽ 获取 Gas 价格...")
	// Get gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Printf("❌ 获取 Gas 价格失败: %v\n", err)
		return fmt.Errorf("failed to get gas price: %v", err)
	}
	gasPriceGwei := new(big.Int).Div(gasPrice, big.NewInt(1e9))
	fmt.Printf("✅ Gas 价格: %s wei (%s Gwei)\n", gasPrice.String(), gasPriceGwei.String())

	// 检查发送方余额
	fmt.Println("💰 检查发送方余额...")
	balance, err := client.BalanceAt(context.Background(), fromAddress, nil)
	if err != nil {
		fmt.Printf("❌ 获取余额失败: %v\n", err)
		return fmt.Errorf("failed to get balance: %v", err)
	}
	balanceEth := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(1e18))
	fmt.Printf("✅ 当前余额: %s ETH (%s wei)\n", balanceEth.Text('f', 18), balance.String())

	// 检查余额是否足够
	gasCost := new(big.Int).Mul(gasPrice, big.NewInt(21000))
	totalCost := new(big.Int).Add(amount, gasCost)
	if balance.Cmp(totalCost) < 0 {
		fmt.Printf("❌ 余额不足！需要 %s wei，当前余额 %s wei\n", totalCost.String(), balance.String())
		return fmt.Errorf("insufficient balance: need %s wei, have %s wei", totalCost.String(), balance.String())
	}
	fmt.Printf("✅ 余额充足，可以执行转账\n")

	fmt.Println("📝 创建交易...")
	// Create transaction
	tx := types.NewTransaction(
		nonce,
		toAddr,
		amount,
		21000, // Standard gas limit for ETH transfer
		gasPrice,
		nil, // No data for simple ETH transfer
	)
	fmt.Printf("✅ 交易创建成功\n")
	fmt.Printf("   Nonce: %d\n", nonce)
	fmt.Printf("   接收方: %s\n", toAddr.Hex())
	fmt.Printf("   金额: %s wei\n", amount.String())
	fmt.Printf("   Gas 限制: 21000\n")
	fmt.Printf("   Gas 价格: %s wei\n", gasPrice.String())

	fmt.Println("🔗 获取链 ID...")
	// Get chain ID
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		fmt.Printf("❌ 获取链 ID 失败: %v\n", err)
		return fmt.Errorf("failed to get chain ID: %v", err)
	}
	fmt.Printf("✅ 链 ID: %s\n", chainID.String())

	fmt.Println("🔐 开始签名交易...")
	// Sign transaction with secure keystore (password prompted here)
	signedTx, err := wallet.SignTransaction(tx, chainID)
	if err != nil {
		fmt.Printf("❌ 交易签名失败: %v\n", err)
		return fmt.Errorf("failed to sign transaction: %v", err)
	}
	fmt.Printf("✅ 交易签名成功\n")

	fmt.Println("📤 发送交易到网络...")
	// Send transaction
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		fmt.Printf("❌ 发送交易失败: %v\n", err)
		return fmt.Errorf("failed to send transaction: %v", err)
	}

	txHash := signedTx.Hash().Hex()
	fmt.Printf("🎉 交易发送成功！\n")
	fmt.Printf("📋 交易哈希: %s\n", txHash)
	fmt.Printf("🔗 Sepolia Etherscan: https://sepolia.etherscan.io/tx/%s\n", txHash)

	// 计算预估 Gas 费用
	estimatedGasCost := new(big.Int).Mul(gasPrice, big.NewInt(21000))
	estimatedGasCostGwei := new(big.Int).Div(estimatedGasCost, big.NewInt(1e9))
	fmt.Printf("⛽ 预估 Gas 费用: %s wei (%s Gwei)\n", estimatedGasCost.String(), estimatedGasCostGwei.String())

	return nil
}

// CreateSecureKeystoreFile creates a new secure keystore file
func CreateSecureKeystoreFile(keystorePath string) (string, error) {
	// Create keystore directory if it doesn't exist
	if err := os.MkdirAll(keystorePath, 0700); err != nil {
		return "", fmt.Errorf("failed to create keystore directory: %v", err)
	}

	// Prompt for password
	password, err := promptPassword("Enter password for new keystore: ")
	if err != nil {
		return "", err
	}

	// Confirm password
	confirmPassword, err := promptPassword("Confirm password: ")
	if err != nil {
		return "", err
	}

	if password != confirmPassword {
		return "", fmt.Errorf("passwords do not match")
	}

	// Create new keystore
	ks := keystore.NewKeyStore(keystorePath, keystore.StandardScryptN, keystore.StandardScryptP)

	// Generate new private key
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "", fmt.Errorf("failed to generate private key: %v", err)
	}

	// Import private key to keystore
	account, err := ks.ImportECDSA(privateKey, password)
	if err != nil {
		return "", fmt.Errorf("failed to import private key: %v", err)
	}

	fmt.Printf("Secure keystore file created successfully!\n")
	fmt.Printf("Address: %s\n", account.Address.Hex())
	fmt.Printf("Keystore file: %s\n", account.URL.Path)
	fmt.Println("⚠️  IMPORTANT: Remember your password! It cannot be recovered!")

	return account.URL.Path, nil
}

// ValidateSecureKeystoreFile validates a keystore file
func ValidateSecureKeystoreFile(keystoreFile string) error {
	data, err := ioutil.ReadFile(keystoreFile)
	if err != nil {
		return fmt.Errorf("failed to read keystore file: %v", err)
	}

	var keystoreData map[string]interface{}
	if err := json.Unmarshal(data, &keystoreData); err != nil {
		return fmt.Errorf("failed to parse keystore file: %v", err)
	}

	// Check required fields
	requiredFields := []string{"version", "address", "crypto"}
	for _, field := range requiredFields {
		if _, exists := keystoreData[field]; !exists {
			return fmt.Errorf("invalid keystore file: missing %s", field)
		}
	}

	// Check crypto fields
	crypto, ok := keystoreData["crypto"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid keystore file: invalid crypto section")
	}

	cryptoFields := []string{"ciphertext", "cipherparams", "cipher", "kdf", "kdfparams", "mac"}
	for _, field := range cryptoFields {
		if _, exists := crypto[field]; !exists {
			return fmt.Errorf("invalid keystore file: missing crypto.%s", field)
		}
	}

	return nil
}

// ListSecureKeystoreFiles lists all keystore files in the directory
func ListSecureKeystoreFiles(keystorePath string) ([]string, error) {
	files, err := filepath.Glob(filepath.Join(keystorePath, "UTC--*"))
	if err != nil {
		return nil, fmt.Errorf("failed to list keystore files: %v", err)
	}

	var validFiles []string
	for _, file := range files {
		filename := filepath.Base(file)
		if isValidKeystoreFilename(filename) {
			validFiles = append(validFiles, file)
		}
	}

	return validFiles, nil
}

// isValidKeystoreFilename checks if a filename follows the keystore naming convention
func isValidKeystoreFilename(filename string) bool {
	if !strings.HasPrefix(filename, "UTC--") {
		return false
	}

	parts := strings.Split(filename, "--")
	if len(parts) != 3 {
		return false
	}

	timestamp := parts[1]
	if len(timestamp) < 20 {
		return false
	}

	address := parts[2]
	if len(address) != 40 {
		return false
	}

	for _, char := range address {
		if !((char >= '0' && char <= '9') || (char >= 'a' && char <= 'f') || (char >= 'A' && char <= 'F')) {
			return false
		}
	}

	return true
}
