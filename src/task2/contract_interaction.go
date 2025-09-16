package task2

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fuckEthereum/contracts"
)

// ContractInteraction demonstrates how to interact with the Counter contract on Sepolia testnet
type ContractInteraction struct {
	client     *ethclient.Client
	privateKey *ecdsa.PrivateKey
	address    common.Address
	instance   *contracts.Counter
}

// NewContractInteraction creates a new contract interaction instance
func NewContractInteraction(rpcURL, privateKeyHex string) (*ContractInteraction, error) {
	// Connect to Sepolia testnet
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum client: %v", err)
	}

	// Parse private key
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %v", err)
	}

	// Get public key and address
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("failed to cast public key to ECDSA")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	return &ContractInteraction{
		client:     client,
		privateKey: privateKey,
		address:    address,
	}, nil
}

// DeployContract deploys the Counter contract to Sepolia testnet
func (ci *ContractInteraction) DeployContract() error {
	fmt.Println("🚀 Deploying Counter contract to Sepolia testnet...")

	// Get the chain ID
	chainID, err := ci.client.NetworkID(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get chain ID: %v", err)
	}

	// Get the nonce
	nonce, err := ci.client.PendingNonceAt(context.Background(), ci.address)
	if err != nil {
		return fmt.Errorf("failed to get nonce: %v", err)
	}

	// Get gas price
	gasPrice, err := ci.client.SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get gas price: %v", err)
	}

	// Create transaction options
	auth, err := bind.NewKeyedTransactorWithChainID(ci.privateKey, chainID)
	if err != nil {
		return fmt.Errorf("failed to create transactor: %v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // No ETH sent
	auth.GasLimit = uint64(300000) // Set gas limit
	auth.GasPrice = gasPrice

	// Deploy the contract
	contractAddress, tx, instance, err := contracts.DeployCounter(auth, ci.client)
	if err != nil {
		return fmt.Errorf("failed to deploy contract: %v", err)
	}

	ci.instance = instance

	fmt.Printf("📝 Contract deployment transaction hash: %s\n", tx.Hash().Hex())
	fmt.Printf("📍 Contract address: %s\n", contractAddress.Hex())

	// Wait for transaction to be mined
	fmt.Println("⏳ Waiting for transaction to be mined...")
	receipt, err := bind.WaitMined(context.Background(), ci.client, tx)
	if err != nil {
		return fmt.Errorf("failed to wait for transaction: %v", err)
	}

	if receipt.Status == 0 {
		return fmt.Errorf("contract deployment failed")
	}

	fmt.Printf("✅ Contract deployed successfully! Gas used: %d\n", receipt.GasUsed)
	return nil
}

// LoadExistingContract loads an existing contract instance
func (ci *ContractInteraction) LoadExistingContract(contractAddress string) error {
	address := common.HexToAddress(contractAddress)
	instance, err := contracts.NewCounter(address, ci.client)
	if err != nil {
		return fmt.Errorf("failed to create contract instance: %v", err)
	}

	ci.instance = instance
	fmt.Printf("📋 Loaded existing contract at address: %s\n", contractAddress)
	return nil
}

// GetCurrentCount retrieves the current count from the contract
func (ci *ContractInteraction) GetCurrentCount() (*big.Int, error) {
	if ci.instance == nil {
		return nil, fmt.Errorf("contract instance not initialized")
	}

	count, err := ci.instance.GetCount(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get count: %v", err)
	}

	return count, nil
}

// IncrementCount increments the counter by 1
func (ci *ContractInteraction) IncrementCount() error {
	if ci.instance == nil {
		return fmt.Errorf("contract instance not initialized")
	}

	fmt.Println("➕ Incrementing counter...")

	// Get the chain ID
	chainID, err := ci.client.NetworkID(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get chain ID: %v", err)
	}

	// Get the nonce
	nonce, err := ci.client.PendingNonceAt(context.Background(), ci.address)
	if err != nil {
		return fmt.Errorf("failed to get nonce: %v", err)
	}

	// Get gas price
	gasPrice, err := ci.client.SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get gas price: %v", err)
	}

	// Create transaction options
	auth, err := bind.NewKeyedTransactorWithChainID(ci.privateKey, chainID)
	if err != nil {
		return fmt.Errorf("failed to create transactor: %v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // No ETH sent
	auth.GasLimit = uint64(100000) // Set gas limit
	auth.GasPrice = gasPrice

	// Call the increment function
	tx, err := ci.instance.Increment(auth)
	if err != nil {
		return fmt.Errorf("failed to call increment: %v", err)
	}

	fmt.Printf("📝 Increment transaction hash: %s\n", tx.Hash().Hex())

	// Wait for transaction to be mined
	fmt.Println("⏳ Waiting for transaction to be mined...")
	receipt, err := bind.WaitMined(context.Background(), ci.client, tx)
	if err != nil {
		return fmt.Errorf("failed to wait for transaction: %v", err)
	}

	if receipt.Status == 0 {
		return fmt.Errorf("increment transaction failed")
	}

	fmt.Printf("✅ Counter incremented successfully! Gas used: %d\n", receipt.GasUsed)
	return nil
}

// DecrementCount decrements the counter by 1
func (ci *ContractInteraction) DecrementCount() error {
	if ci.instance == nil {
		return fmt.Errorf("contract instance not initialized")
	}

	fmt.Println("➖ Decrementing counter...")

	// Get the chain ID
	chainID, err := ci.client.NetworkID(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get chain ID: %v", err)
	}

	// Get the nonce
	nonce, err := ci.client.PendingNonceAt(context.Background(), ci.address)
	if err != nil {
		return fmt.Errorf("failed to get nonce: %v", err)
	}

	// Get gas price
	gasPrice, err := ci.client.SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get gas price: %v", err)
	}

	// Create transaction options
	auth, err := bind.NewKeyedTransactorWithChainID(ci.privateKey, chainID)
	if err != nil {
		return fmt.Errorf("failed to create transactor: %v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // No ETH sent
	auth.GasLimit = uint64(100000) // Set gas limit
	auth.GasPrice = gasPrice

	// Call the decrement function
	tx, err := ci.instance.Decrement(auth)
	if err != nil {
		return fmt.Errorf("failed to call decrement: %v", err)
	}

	fmt.Printf("📝 Decrement transaction hash: %s\n", tx.Hash().Hex())

	// Wait for transaction to be mined
	fmt.Println("⏳ Waiting for transaction to be mined...")
	receipt, err := bind.WaitMined(context.Background(), ci.client, tx)
	if err != nil {
		return fmt.Errorf("failed to wait for transaction: %v", err)
	}

	if receipt.Status == 0 {
		return fmt.Errorf("decrement transaction failed")
	}

	fmt.Printf("✅ Counter decremented successfully! Gas used: %d\n", receipt.GasUsed)
	return nil
}

// ResetCount resets the counter to 0
func (ci *ContractInteraction) ResetCount() error {
	if ci.instance == nil {
		return fmt.Errorf("contract instance not initialized")
	}

	fmt.Println("🔄 Resetting counter...")

	// Get the chain ID
	chainID, err := ci.client.NetworkID(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get chain ID: %v", err)
	}

	// Get the nonce
	nonce, err := ci.client.PendingNonceAt(context.Background(), ci.address)
	if err != nil {
		return fmt.Errorf("failed to get nonce: %v", err)
	}

	// Get gas price
	gasPrice, err := ci.client.SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get gas price: %v", err)
	}

	// Create transaction options
	auth, err := bind.NewKeyedTransactorWithChainID(ci.privateKey, chainID)
	if err != nil {
		return fmt.Errorf("failed to create transactor: %v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // No ETH sent
	auth.GasLimit = uint64(100000) // Set gas limit
	auth.GasPrice = gasPrice

	// Call the reset function
	tx, err := ci.instance.Reset(auth)
	if err != nil {
		return fmt.Errorf("failed to call reset: %v", err)
	}

	fmt.Printf("📝 Reset transaction hash: %s\n", tx.Hash().Hex())

	// Wait for transaction to be mined
	fmt.Println("⏳ Waiting for transaction to be mined...")
	receipt, err := bind.WaitMined(context.Background(), ci.client, tx)
	if err != nil {
		return fmt.Errorf("failed to wait for transaction: %v", err)
	}

	if receipt.Status == 0 {
		return fmt.Errorf("reset transaction failed")
	}

	fmt.Printf("✅ Counter reset successfully! Gas used: %d\n", receipt.GasUsed)
	return nil
}

// GetAccountBalance returns the ETH balance of the account
func (ci *ContractInteraction) GetAccountBalance() (*big.Int, error) {
	balance, err := ci.client.BalanceAt(context.Background(), ci.address, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get balance: %v", err)
	}
	return balance, nil
}

// Close closes the Ethereum client connection
func (ci *ContractInteraction) Close() {
	if ci.client != nil {
		ci.client.Close()
	}
}

// RunContractDemo demonstrates the complete contract interaction workflow
func RunContractDemo() error {
	// Get configuration from environment variables
	rpcURL := os.Getenv("SEPOLIA_RPC_URL")
	if rpcURL == "" {
		rpcURL = "https://sepolia.infura.io/v3/YOUR_INFURA_PROJECT_ID" // Replace with your Infura project ID
	}

	privateKeyHex := os.Getenv("PRIVATE_KEY")
	if privateKeyHex == "" {
		// Use the private key from the keystore file (for demo purposes)
		// In production, use environment variables or secure key management
		return fmt.Errorf("PRIVATE_KEY environment variable not set")
	}

	// Create contract interaction instance
	ci, err := NewContractInteraction(rpcURL, privateKeyHex)
	if err != nil {
		return fmt.Errorf("failed to create contract interaction: %v", err)
	}
	defer ci.Close()

	// Check account balance
	balance, err := ci.GetAccountBalance()
	if err != nil {
		return fmt.Errorf("failed to get balance: %v", err)
	}
	fmt.Printf("💰 Account balance: %s ETH\n", balance.String())

	// Check if we have enough ETH for gas
	gasPrice, err := ci.client.SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get gas price: %v", err)
	}
	requiredBalance := new(big.Int).Mul(gasPrice, big.NewInt(500000)) // Estimate for deployment + operations

	if balance.Cmp(requiredBalance) < 0 {
		fmt.Printf("⚠️  Warning: Low balance. You may need more ETH for gas fees.\n")
		fmt.Printf("   Current balance: %s ETH\n", balance.String())
		fmt.Printf("   Estimated required: %s ETH\n", requiredBalance.String())
		fmt.Println("   Get testnet ETH from: https://sepoliafaucet.com/")
	}

	// Deploy the contract
	err = ci.DeployContract()
	if err != nil {
		return fmt.Errorf("failed to deploy contract: %v", err)
	}

	// Wait a moment for the contract to be fully deployed
	time.Sleep(2 * time.Second)

	// Get initial count
	count, err := ci.GetCurrentCount()
	if err != nil {
		return fmt.Errorf("failed to get initial count: %v", err)
	}
	fmt.Printf("📊 Initial count: %s\n", count.String())

	// Increment the counter
	err = ci.IncrementCount()
	if err != nil {
		return fmt.Errorf("failed to increment: %v", err)
	}

	// Get count after increment
	count, err = ci.GetCurrentCount()
	if err != nil {
		return fmt.Errorf("failed to get count after increment: %v", err)
	}
	fmt.Printf("📊 Count after increment: %s\n", count.String())

	// Increment again
	err = ci.IncrementCount()
	if err != nil {
		return fmt.Errorf("failed to increment again: %v", err)
	}

	// Get count after second increment
	count, err = ci.GetCurrentCount()
	if err != nil {
		return fmt.Errorf("failed to get count after second increment: %v", err)
	}
	fmt.Printf("📊 Count after second increment: %s\n", count.String())

	// Decrement the counter
	err = ci.DecrementCount()
	if err != nil {
		return fmt.Errorf("failed to decrement: %v", err)
	}

	// Get count after decrement
	count, err = ci.GetCurrentCount()
	if err != nil {
		return fmt.Errorf("failed to get count after decrement: %v", err)
	}
	fmt.Printf("📊 Count after decrement: %s\n", count.String())

	// Reset the counter
	err = ci.ResetCount()
	if err != nil {
		return fmt.Errorf("failed to reset: %v", err)
	}

	// Get final count
	count, err = ci.GetCurrentCount()
	if err != nil {
		return fmt.Errorf("failed to get final count: %v", err)
	}
	fmt.Printf("📊 Final count after reset: %s\n", count.String())

	fmt.Println("🎉 Contract interaction demo completed successfully!")
	return nil
}
