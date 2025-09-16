#!/bin/bash

# Script to compile Solidity contracts and generate ABI and bytecode files

echo "🔨 Compiling Solidity contracts..."

# Create output directory if it doesn't exist
mkdir -p build

# Check if solc is installed
if ! command -v solc &> /dev/null; then
    echo "❌ solc (Solidity compiler) is not installed."
    echo "Please install it using:"
    echo "  - macOS: brew install solidity"
    echo "  - Ubuntu: sudo apt-get install solc"
    echo "  - Or download from: https://github.com/ethereum/solidity/releases"
    exit 1
fi

# Compile the Counter contract
echo "📝 Compiling Counter.sol..."
solc --abi --bin --overwrite -o build contracts/Counter.sol

# Check if compilation was successful
if [ $? -eq 0 ]; then
    echo "✅ Compilation successful!"
    echo "📁 Generated files in build/ directory:"
    ls -la build/
else
    echo "❌ Compilation failed!"
    exit 1
fi

# Create a combined JSON output for easier abigen usage
echo "📦 Creating combined JSON output..."
solc --combined-json abi,bin,metadata -o build contracts/Counter.sol

echo "🎉 Compilation complete!"
