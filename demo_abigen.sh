#!/bin/bash

# Complete demonstration of abigen workflow for smart contract interaction

echo "🚀 Abigen Smart Contract Interaction Demo"
echo "========================================"
echo ""

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    echo -e "${GREEN}✅ $1${NC}"
}

print_warning() {
    echo -e "${YELLOW}⚠️  $1${NC}"
}

print_error() {
    echo -e "${RED}❌ $1${NC}"
}

print_info() {
    echo -e "${BLUE}ℹ️  $1${NC}"
}

# Step 1: Check prerequisites
echo "📋 Step 1: Checking Prerequisites"
echo "---------------------------------"

# Check Go
if command -v go &> /dev/null; then
    print_status "Go is installed: $(go version | cut -d' ' -f3)"
else
    print_error "Go is not installed. Please install Go 1.19+"
    exit 1
fi

# Check solc
if command -v solc &> /dev/null; then
    print_status "Solidity compiler is installed: $(solc --version | head -n1 | cut -d' ' -f2)"
else
    print_error "Solidity compiler is not installed. Please install solc"
    exit 1
fi

# Check abigen
if command -v abigen &> /dev/null; then
    print_status "Abigen is installed"
else
    ABIGEN_PATH=$(find $(go env GOPATH) -name abigen 2>/dev/null | head -n1)
    if [ -n "$ABIGEN_PATH" ]; then
        print_status "Abigen found at: $ABIGEN_PATH"
        export PATH=$PATH:$(dirname $ABIGEN_PATH)
    else
        print_error "Abigen is not installed. Please run: go install github.com/ethereum/go-ethereum/cmd/abigen@latest"
        exit 1
    fi
fi

echo ""

# Step 2: Show project structure
echo "📁 Step 2: Project Structure"
echo "---------------------------"
print_info "Smart Contract: contracts/Counter.sol"
print_info "Generated Bindings: contracts/counter.go"
print_info "ABI File: build/Counter.abi"
print_info "Bytecode: build/Counter.bin"
print_info "Go Code: src/task2/contract_interaction.go"
echo ""

# Step 3: Show smart contract
echo "📝 Step 3: Smart Contract Overview"
echo "--------------------------------"
print_info "Counter Contract Features:"
echo "  • increment() - Increments counter by 1"
echo "  • decrement() - Decrements counter by 1 (with underflow protection)"
echo "  • reset() - Resets counter to 0"
echo "  • getCount() - Returns current count value"
echo "  • Events: CountIncremented, CountDecremented, CountReset"
echo ""

# Step 4: Show compilation process
echo "🔨 Step 4: Compilation Process"
echo "-----------------------------"
print_info "Compiling Solidity contract..."
if [ -f "build/Counter.abi" ] && [ -f "build/Counter.bin" ]; then
    print_status "Contract already compiled"
    echo "  ABI size: $(wc -c < build/Counter.abi) bytes"
    echo "  Bytecode size: $(wc -c < build/Counter.bin) bytes"
else
    print_info "Running compilation..."
    ./scripts/compile.sh
    if [ $? -eq 0 ]; then
        print_status "Compilation successful"
    else
        print_error "Compilation failed"
        exit 1
    fi
fi
echo ""

# Step 5: Show Go bindings generation
echo "📦 Step 5: Go Bindings Generation"
echo "--------------------------------"
print_info "Generated Go bindings include:"
echo "  • contracts.Counter - Main contract struct"
echo "  • contracts.DeployCounter() - Contract deployment function"
echo "  • Method bindings for all contract functions"
echo "  • Event handling structures"
echo "  • Type-safe parameter handling"
echo ""

# Show generated bindings size
if [ -f "contracts/counter.go" ]; then
    print_status "Go bindings generated: $(wc -l < contracts/counter.go) lines"
else
    print_warning "Go bindings not found. Run: abigen --abi build/Counter.abi --bin build/Counter.bin --pkg contracts --type Counter --out contracts/counter.go"
fi
echo ""

# Step 6: Show contract interaction features
echo "🔗 Step 6: Contract Interaction Features"
echo "---------------------------------------"
print_info "The Go application provides:"
echo "  • Contract deployment to Sepolia testnet"
echo "  • Read operations (getCount)"
echo "  • Write operations (increment, decrement, reset)"
echo "  • Transaction management (gas estimation, nonce handling)"
echo "  • Event monitoring capabilities"
echo "  • Error handling and logging"
echo "  • Account balance checking"
echo ""

# Step 7: Show usage instructions
echo "🚀 Step 7: Usage Instructions"
echo "----------------------------"
print_info "To run the complete demo:"
echo ""
echo "1. Set up environment variables:"
echo "   export PRIVATE_KEY=your_private_key_here"
echo "   export SEPOLIA_RPC_URL=https://sepolia.infura.io/v3/YOUR_PROJECT_ID"
echo ""
echo "2. Get testnet ETH:"
echo "   Visit: https://sepoliafaucet.com/"
echo ""
echo "3. Run the demo:"
echo "   go run main.go task2"
echo ""
echo "4. Or use the binary:"
echo "   ./ethereum-demo task2"
echo ""

# Step 8: Show generated files
echo "📄 Step 8: Generated Files"
echo "-------------------------"
if [ -d "build" ]; then
    print_info "Build directory contents:"
    ls -la build/ | grep -E "\.(abi|bin|json)$" | while read line; do
        echo "  $line"
    done
fi

if [ -f "contracts/counter.go" ]; then
    print_info "Generated Go bindings:"
    echo "  contracts/counter.go ($(wc -l < contracts/counter.go) lines)"
fi
echo ""

# Step 9: Show next steps
echo "🎯 Step 9: Next Steps"
echo "--------------------"
print_info "You can now:"
echo "  • Deploy the contract to Sepolia testnet"
echo "  • Interact with the deployed contract"
echo "  • Monitor contract events"
echo "  • Extend the contract with more functionality"
echo "  • Create a web frontend for contract interaction"
echo ""

# Step 10: Show resources
echo "📚 Step 10: Resources"
echo "--------------------"
print_info "Useful resources:"
echo "  • Documentation: ABIGEN_README.md"
echo "  • Sepolia Faucet: https://sepoliafaucet.com/"
echo "  • Etherscan Sepolia: https://sepolia.etherscan.io/"
echo "  • Go Ethereum Docs: https://pkg.go.dev/github.com/ethereum/go-ethereum"
echo "  • Solidity Docs: https://docs.soliditylang.org/"
echo ""

print_status "Demo completed! Ready to interact with smart contracts on Sepolia testnet."
echo ""
echo "🎉 Happy coding with Ethereum and Go!"
