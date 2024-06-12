package core

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
)

// EDE is an Encrypt and Decrypt structure
// Use Secret Key and Origin Data
type EDE struct {
	Key  string `json:"key,omitempty"`
	Data string `json:"data,omitempty"`
}

func Encrypt(e EDE) (encrypted string, err error) {
	key := []byte(e.Key)
	origData := []byte(e.Data)
	// Convert the byte key to a block
	block, err := des.NewCipher(key)
	// Perform a complement operation on the plaintext first
	origData = pkcs5Padding(origData, block.BlockSize())
	// Set the encryption method
	blockMode := cipher.NewCBCEncrypter(block, key)
	// Create an array of bytes of plaintext length
	encryptData := make([]byte, len(origData))
	// Encrypt plaintext, and put the encrypted data into an array
	blockMode.CryptBlocks(encryptData, origData)
	// Convert byte arrays to strings
	encrypted = base64.StdEncoding.EncodeToString(encryptData)
	return
}

func Decrypt(e EDE) (decrypted string, err error) {
	key := []byte(e.Key)
	// Flashback performs the encryption method once
	// Convert strings to byte arrays
	decryptData, _ := base64.StdEncoding.DecodeString(e.Data)
	// Convert the byte key to a block
	block, err := des.NewCipher(key)
	// Set the encryption method
	blockMode := cipher.NewCBCDecrypter(block, key)
	// Create an array variable of ciphertext size
	origData := make([]byte, len(decryptData))
	// Decrypt the ciphertext into the origData array
	blockMode.CryptBlocks(origData, decryptData)
	// Go to the complement
	origData = pkcs5UnPadding(origData)
	decrypted = string(origData)
	return
}

// pkcs5Padding implement the complement of the plaintext
func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	repeat := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, repeat...)
}

// pkcs5UnPadding remove the complement
func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}
