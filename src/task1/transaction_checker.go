package task1

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// TransactionStatus represents the status of a transaction
type TransactionStatus struct {
	Hash              string
	Status            string
	BlockNumber       *big.Int
	GasUsed           uint64
	EffectiveGasPrice *big.Int
	From              common.Address
	To                *common.Address
	Value             *big.Int
	Network           string
	Error             string
}

// CheckTransactionStatus checks the status of a transaction
func CheckTransactionStatus(txHash string, rpcURL string) (*TransactionStatus, error) {
	// Connect to Ethereum client
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum client: %v", err)
	}
	defer client.Close()

	// Get network information
	networkID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get network ID: %v", err)
	}

	// Determine network name
	networkName := getNetworkName(networkID)

	// Parse transaction hash
	hash := common.HexToHash(txHash)

	// Get transaction receipt
	receipt, err := client.TransactionReceipt(context.Background(), hash)
	if err != nil {
		// Check if transaction is pending
		_, isPending, err := client.TransactionByHash(context.Background(), hash)
		if err != nil {
			return nil, fmt.Errorf("transaction not found: %v", err)
		}

		if isPending {
			return &TransactionStatus{
				Hash:    txHash,
				Status:  "PENDING",
				Network: networkName,
				Error:   "Transaction is still pending",
			}, nil
		}

		return nil, fmt.Errorf("failed to get transaction receipt: %v", err)
	}

	// Get transaction details
	tx, _, err := client.TransactionByHash(context.Background(), hash)
	if err != nil {
		return nil, fmt.Errorf("failed to get transaction details: %v", err)
	}

	// Determine transaction status
	status := "SUCCESS"
	if receipt.Status == 0 {
		status = "FAILED"
	}

	// Get sender address from transaction
	fromAddr := common.Address{}
	if tx.To() != nil {
		// This is a simplified approach - in production you'd need to recover the sender
		// For now, we'll use a placeholder
		fromAddr = common.HexToAddress("0x0000000000000000000000000000000000000000")
	}

	return &TransactionStatus{
		Hash:              txHash,
		Status:            status,
		BlockNumber:       receipt.BlockNumber,
		GasUsed:           receipt.GasUsed,
		EffectiveGasPrice: receipt.EffectiveGasPrice,
		From:              fromAddr,
		To:                tx.To(),
		Value:             tx.Value(),
		Network:           networkName,
	}, nil
}

// WaitForTransaction waits for a transaction to be mined
func WaitForTransaction(txHash string, rpcURL string, maxWaitTime time.Duration) (*TransactionStatus, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum client: %v", err)
	}
	defer client.Close()

	hash := common.HexToHash(txHash)
	fmt.Println("txHash", txHash)
	startTime := time.Now()

	for time.Since(startTime) < maxWaitTime {
		// Check if transaction is mined
		_, err := client.TransactionReceipt(context.Background(), hash)
		if err == nil {
			// Transaction is mined
			return CheckTransactionStatus(txHash, rpcURL)
		}

		// Wait before checking again
		time.Sleep(5 * time.Second)
	}

	return nil, fmt.Errorf("transaction not mined within %v", maxWaitTime)
}

// GetAccountBalance gets the balance of an account
func GetAccountBalance(address string, rpcURL string) (*big.Int, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum client: %v", err)
	}
	defer client.Close()

	addr := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), addr, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get balance: %v", err)
	}

	return balance, nil
}

// GetGasPrice gets the current gas price
func GetGasPrice(rpcURL string) (*big.Int, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum client: %v", err)
	}
	defer client.Close()

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get gas price: %v", err)
	}

	return gasPrice, nil
}

// GetNetworkInfo gets information about the current network
func GetNetworkInfo(rpcURL string) (map[string]interface{}, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum client: %v", err)
	}
	defer client.Close()

	// Get network ID
	networkID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get network ID: %v", err)
	}

	// Get latest block
	latestBlock, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get latest block: %v", err)
	}

	// Get gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get gas price: %v", err)
	}

	return map[string]interface{}{
		"networkID":    networkID.String(),
		"networkName":  getNetworkName(networkID),
		"latestBlock":  latestBlock.Number().String(),
		"gasPrice":     gasPrice.String(),
		"gasPriceGwei": new(big.Int).Div(gasPrice, big.NewInt(1e9)).String(),
	}, nil
}

// getNetworkName converts network ID to human-readable name
func getNetworkName(networkID *big.Int) string {
	switch networkID.String() {
	case "1":
		return "Ethereum Mainnet"
	case "3":
		return "Ropsten Testnet (Deprecated)"
	case "4":
		return "Rinkeby Testnet (Deprecated)"
	case "5":
		return "Goerli Testnet (Deprecated)"
	case "11155111":
		return "Sepolia Testnet"
	case "137":
		return "Polygon Mainnet"
	case "80001":
		return "Polygon Mumbai Testnet"
	case "56":
		return "BSC Mainnet"
	case "97":
		return "BSC Testnet"
	default:
		return fmt.Sprintf("Unknown Network (ID: %s)", networkID.String())
	}
}

// ValidateTransaction checks if a transaction can be sent
func ValidateTransaction(fromAddress, toAddress string, amount *big.Int, rpcURL string) error {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return fmt.Errorf("failed to connect to Ethereum client: %v", err)
	}
	defer client.Close()

	fromAddr := common.HexToAddress(fromAddress)
	_ = common.HexToAddress(toAddress) // Use toAddress parameter

	// Check sender balance
	balance, err := client.BalanceAt(context.Background(), fromAddr, nil)
	if err != nil {
		return fmt.Errorf("failed to get balance: %v", err)
	}

	// Get gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get gas price: %v", err)
	}

	// Estimate gas
	gasLimit := uint64(21000) // Standard ETH transfer
	gasCost := new(big.Int).Mul(gasPrice, big.NewInt(int64(gasLimit)))
	totalCost := new(big.Int).Add(amount, gasCost)

	if balance.Cmp(totalCost) < 0 {
		return fmt.Errorf("insufficient balance: have %s wei, need %s wei", balance.String(), totalCost.String())
	}

	fmt.Printf("✅ Transaction validation passed\n")
	fmt.Printf("💰 Balance: %s wei\n", balance.String())
	fmt.Printf("💸 Transfer amount: %s wei\n", amount.String())
	fmt.Printf("⛽ Gas cost: %s wei\n", gasCost.String())
	fmt.Printf("💳 Total cost: %s wei\n", totalCost.String())

	return nil
}
