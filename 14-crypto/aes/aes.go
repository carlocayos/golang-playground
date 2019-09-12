package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func printUsage() {
	fmt.Printf(os.Args[0] + `

Encrypt or decrypt a file using AES with a 256-bit key file.
This program can also generate 256-bit keys.

Usage:
  ` + os.Args[0] + ` [-h|--help]
  ` + os.Args[0] + ` [-g|--genkey]
  ` + os.Args[0] + ` <keyFile> <file> [-d|--decrypt]

Examples:
  # Generate a 32-byte (256-bit) key
  ` + os.Args[0] + ` --genkey

  # Encrypt with secret key. Output to STDOUT
  ` + os.Args[0] + ` --genkey > secret.key

  # Encrypt message using secret key. Output to ciphertext.dat
  ` + os.Args[0] + ` secret.key message.txt > ciphertext.dat

  # Decrypt message using secret key. Output to STDOUT
  ` + os.Args[0] + ` secret.key ciphertext.dat -d

  # Decrypt message using secret key. Output to message.txt
  ` + os.Args[0] + ` secret.key ciphertext.dat -d > cleartext.txt
`)
}

func generateKey() []byte {
	// generate random bytes
	randomBytes := make([]byte, 32) // 32 bytes is 256 bits
	numberOfBytes, err := rand.Read(randomBytes)
	if err != nil {
		log.Fatal("error generating random key", err)
	}

	if numberOfBytes != 32 {
		log.Fatal("incorrect number of bytes", err)
	}
	return randomBytes
}

// AES Encryption
func encrypt(key, message []byte) ([]byte, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// create the byte slice that will hold the encrypted message
	cipherText := make([]byte, aes.BlockSize+len(message))

	// Generate the Initialization Vector (IV) nonce
	// which is stored at the beginning of the byte slice
	// The IV is the same length as the AES blocksize
	iv := cipherText[:aes.BlockSize]
	_, err = io.ReadFull(rand.Reader, iv)
	if err != nil {
		return nil, err
	}

	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], message)

	return cipherText, nil
}

func checkArgs() (string, string, bool) {

	if len(os.Args) < 2 || len(os.Args) > 4 {
		printUsage()
		os.Exit(1)
	}

	// if one arg provided
	if len(os.Args) == 2 {
		// only --help and --genkey are valid one-argument uses
		if os.Args[1] == "-h" || os.Args[1] == "--help" {
			printUsage()
			os.Exit(0) // exit gracefully no error
		}
		if os.Args[1] == "-g" || os.Args[1] == "--genkey" {
			// TODO: func generateKey
			key := generateKey()
			fmt.Printf("Generated Key: %x\n", key)
			os.Exit(0) // graceful exit
		}
	}

	// The only use options left is
	// encrypt <keyFile> <file> [-d|--decrypt]
	// If there are only 2 args provided, they must be the
	// keyFile and file without a decrypt flag.
	if len(os.Args) == 3 {
		// keyfile, File, decryptFlag
		return os.Args[1], os.Args[2], false
	}

	// If 3 args are provided,
	// check that the last one is -d or --decrypt
	if len(os.Args) == 4 {
		if os.Args[3] != "-d" || os.Args[3] != "--decrypt" {
			fmt.Println("Error: unknown usage.")
			printUsage()
			os.Exit(1) // exit with error
		}
		return os.Args[1], os.Args[2], true
	}
	return "", "", false
}

func main() {
	log.Println("TEST ME")
	keyFile, file, decryptFlag := checkArgs()
	fmt.Println(keyFile, file, decryptFlag)

	// load key from file
	keyFileData, err := ioutil.ReadFile(keyFile)
	if err != nil {
		log.Fatal("invalid key", err)
	}

	fmt.Printf("keyFileData = %s", keyFileData)

}
