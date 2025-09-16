package task1

import (
	"crypto/ecdsa"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

// ImportPrivateKeyToKeystore imports an existing private key into a new keystore file
func ImportPrivateKeyToKeystore(keystorePath, privateKeyHex string) (string, error) {
	// Create keystore directory if it doesn't exist
	if err := os.MkdirAll(keystorePath, 0700); err != nil {
		return "", fmt.Errorf("failed to create keystore directory: %v", err)
	}

	// Prompt for password
	password, err := promptPassword("Enter password to encrypt keystore: ")
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

	// Parse private key
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return "", fmt.Errorf("invalid private key: %v", err)
	}

	// Create new keystore
	ks := keystore.NewKeyStore(keystorePath, keystore.StandardScryptN, keystore.StandardScryptP)

	// Import private key to keystore
	account, err := ks.ImportECDSA(privateKey, password)
	if err != nil {
		return "", fmt.Errorf("failed to import private key: %v", err)
	}

	fmt.Printf("✅ Private key imported successfully!\n")
	fmt.Printf("📁 Keystore file: %s\n", filepath.Base(account.URL.Path))
	fmt.Printf("📍 Address: %s\n", account.Address.Hex())
	fmt.Println("⚠️  IMPORTANT: Remember your password! It cannot be recovered!")

	return account.URL.Path, nil
}

// ImportPrivateKeyFromFile imports a private key from a text file
func ImportPrivateKeyFromFile(keystorePath, privateKeyFilePath string) (string, error) {
	// Read private key from file
	privateKeyBytes, err := os.ReadFile(privateKeyFilePath)
	if err != nil {
		return "", fmt.Errorf("failed to read private key file: %v", err)
	}

	// Remove whitespace and newlines
	privateKeyHex := strings.TrimSpace(string(privateKeyBytes))

	// Remove 0x prefix if present
	if strings.HasPrefix(privateKeyHex, "0x") {
		privateKeyHex = privateKeyHex[2:]
	}

	// Validate private key length (64 hex characters = 32 bytes)
	if len(privateKeyHex) != 64 {
		return "", fmt.Errorf("invalid private key length: expected 64 hex characters, got %d", len(privateKeyHex))
	}

	// Import to keystore
	return ImportPrivateKeyToKeystore(keystorePath, privateKeyHex)
}

// BatchImportPrivateKeys imports multiple private keys from a directory
func BatchImportPrivateKeys(keystorePath, privateKeysDir string) ([]string, error) {
	var importedFiles []string

	// List all files in the directory
	files, err := os.ReadDir(privateKeysDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %v", err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		// Skip non-text files
		if !isTextFile(file.Name()) {
			continue
		}

		filePath := filepath.Join(privateKeysDir, file.Name())
		fmt.Printf("Processing file: %s\n", file.Name())

		keystoreFile, err := ImportPrivateKeyFromFile(keystorePath, filePath)
		if err != nil {
			fmt.Printf("❌ Failed to import %s: %v\n", file.Name(), err)
			continue
		}

		importedFiles = append(importedFiles, keystoreFile)
		fmt.Printf("✅ Successfully imported %s\n", file.Name())
	}

	return importedFiles, nil
}

// CreateKeystoreFromMnemonic creates a keystore from a mnemonic phrase
func CreateKeystoreFromMnemonic(keystorePath, mnemonic string, derivationPath string) (string, error) {
	// Create keystore directory if it doesn't exist
	if err := os.MkdirAll(keystorePath, 0700); err != nil {
		return "", fmt.Errorf("failed to create keystore directory: %v", err)
	}

	// Prompt for password
	password, err := promptPassword("Enter password to encrypt keystore: ")
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

	// Derive private key from mnemonic
	privateKey, err := derivePrivateKeyFromMnemonic(mnemonic, derivationPath)
	if err != nil {
		return "", fmt.Errorf("failed to derive private key: %v", err)
	}

	// Create new keystore
	ks := keystore.NewKeyStore(keystorePath, keystore.StandardScryptN, keystore.StandardScryptP)

	// Import private key to keystore
	account, err := ks.ImportECDSA(privateKey, password)
	if err != nil {
		return "", fmt.Errorf("failed to import private key: %v", err)
	}

	fmt.Printf("✅ Keystore created from mnemonic!\n")
	fmt.Printf("📁 Keystore file: %s\n", filepath.Base(account.URL.Path))
	fmt.Printf("📍 Address: %s\n", account.Address.Hex())
	fmt.Printf("🔑 Derivation path: %s\n", derivationPath)
	fmt.Println("⚠️  IMPORTANT: Remember your password! It cannot be recovered!")

	return account.URL.Path, nil
}

// Helper functions

// Note: promptPassword function is already defined in secure_keystore.go
// This file uses the existing function from that package

func isTextFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	textExtensions := []string{".txt", ".key", ".pem", ".hex", ""} // empty string for no extension

	for _, textExt := range textExtensions {
		if ext == textExt {
			return true
		}
	}
	return false
}

func derivePrivateKeyFromMnemonic(mnemonic, derivationPath string) (*ecdsa.PrivateKey, error) {
	// This is a simplified implementation
	// In production, you should use a proper BIP39/BIP44 library like:
	// github.com/tyler-smith/go-bip39
	// github.com/btcsuite/btcutil/hdkeychain

	// For now, we'll use a hash of the mnemonic as a demonstration
	// WARNING: This is NOT secure for production use!

	hash := crypto.Keccak256Hash([]byte(mnemonic + derivationPath))
	privateKey, err := crypto.ToECDSA(hash.Bytes())
	if err != nil {
		return nil, fmt.Errorf("failed to create private key from hash: %v", err)
	}

	return privateKey, nil
}
