#!/bin/bash

# Setup script for abigen smart contract interaction demo

echo "🚀 Setting up Abigen Smart Contract Interaction Demo"
echo "=================================================="

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go 1.19+ first."
    echo "   Download from: https://golang.org/doc/install"
    exit 1
fi

echo "✅ Go is installed: $(go version)"

# Check if solc is installed
if ! command -v solc &> /dev/null; then
    echo "📦 Installing Solidity compiler..."
    if command -v brew &> /dev/null; then
        brew install solidity
    else
        echo "❌ Homebrew not found. Please install solc manually:"
        echo "   Download from: https://github.com/ethereum/solidity/releases"
        exit 1
    fi
fi

echo "✅ Solidity compiler is installed: $(solc --version | head -n1)"

# Install abigen
echo "📦 Installing abigen..."
go install github.com/ethereum/go-ethereum/cmd/abigen@latest

# Check if abigen was installed successfully
if ! command -v abigen &> /dev/null; then
    # Try to find abigen in GOPATH
    ABIGEN_PATH=$(find $(go env GOPATH) -name abigen 2>/dev/null | head -n1)
    if [ -n "$ABIGEN_PATH" ]; then
        echo "✅ Abigen found at: $ABIGEN_PATH"
        echo "   Add to PATH: export PATH=\$PATH:$(dirname $ABIGEN_PATH)"
    else
        echo "❌ Failed to install abigen"
        exit 1
    fi
else
    echo "✅ Abigen is installed: $(abigen --version)"
fi

# Compile the smart contract
echo "🔨 Compiling smart contract..."
chmod +x scripts/compile.sh
./scripts/compile.sh

if [ $? -ne 0 ]; then
    echo "❌ Contract compilation failed"
    exit 1
fi

# Generate Go bindings
echo "📝 Generating Go bindings..."
ABIGEN_CMD="abigen"
if ! command -v abigen &> /dev/null; then
    ABIGEN_CMD="$(go env GOPATH)/bin/abigen"
fi

$ABIGEN_CMD --abi build/Counter.abi --bin build/Counter.bin --pkg contracts --type Counter --out contracts/counter.go

if [ $? -ne 0 ]; then
    echo "❌ Failed to generate Go bindings"
    exit 1
fi

echo "✅ Go bindings generated successfully"

# Check if Go modules are initialized
if [ ! -f "go.mod" ]; then
    echo "📦 Initializing Go modules..."
    go mod init github.com/fuckEthereum
fi

# Download dependencies
echo "📦 Downloading Go dependencies..."
go mod tidy

# Create environment file template
echo "📝 Creating environment template..."
cat > .env.example << EOF
# Ethereum Configuration
PRIVATE_KEY=your_private_key_here
SEPOLIA_RPC_URL=https://sepolia.infura.io/v3/YOUR_PROJECT_ID

# Optional: Custom RPC endpoints
# ALCHEMY_URL=https://eth-sepolia.g.alchemy.com/v2/YOUR_API_KEY
# QUICKNODE_URL=https://your-endpoint.quiknode.pro/YOUR_API_KEY/
EOF

echo ""
echo "🎉 Setup completed successfully!"
echo ""
echo "Next steps:"
echo "1. Copy .env.example to .env and fill in your values:"
echo "   cp .env.example .env"
echo ""
echo "2. Get your private key from the keystore:"
echo "   go run src/task1/keystore_import.go"
echo ""
echo "3. Get Sepolia testnet ETH from:"
echo "   https://sepoliafaucet.com/"
echo ""
echo "4. Run the demo:"
echo "   go run main.go task2"
echo ""
echo "For detailed instructions, see ABIGEN_README.md"
