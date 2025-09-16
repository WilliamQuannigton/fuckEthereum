package task2

import (
	"fmt"
	"os"
)

// RunTask2 demonstrates abigen usage for smart contract interaction
func RunTask2() error {
	fmt.Println("🚀 Starting Task 2: Abigen Smart Contract Interaction Demo")
	fmt.Println("============================================================")

	// Check if required environment variables are set
	if os.Getenv("PRIVATE_KEY") == "" {
		fmt.Println("❌ Error: PRIVATE_KEY environment variable not set")
		fmt.Println("Please set your private key:")
		fmt.Println("  export PRIVATE_KEY=your_private_key_here")
		fmt.Println("")
		fmt.Println("You can get your private key from the keystore file:")
		fmt.Println("  go run src/task1/keystore_import.go")
		return fmt.Errorf("PRIVATE_KEY not set")
	}

	if os.Getenv("SEPOLIA_RPC_URL") == "" {
		fmt.Println("⚠️  Warning: SEPOLIA_RPC_URL not set, using default Infura URL")
		fmt.Println("For better performance, set your own RPC URL:")
		fmt.Println("  export SEPOLIA_RPC_URL=https://sepolia.infura.io/v3/YOUR_PROJECT_ID")
		fmt.Println("")
	}

	// Run the contract interaction demo
	err := RunContractDemo()
	if err != nil {
		return fmt.Errorf("contract demo failed: %v", err)
	}

	return nil
}
