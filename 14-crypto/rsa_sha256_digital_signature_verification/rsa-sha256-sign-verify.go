package main

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func printUsage() {
	fmt.Println(os.Args[0] + `

Verify an RSA signature of a message using SHA-256 hashing.
Public key is expected to be a PEM file.

Usage:
  ` + os.Args[0] + ` <publicKeyFilename> <signatureFilename> <messageFilename>

Example:
  ` + os.Args[0] + ` public.pem signature.txt message.txt
`)
}

func checkArgs() (string, string, string) {

	if len(os.Args) != 4 {
		printUsage()
		os.Exit(1)
	}

	return os.Args[1], os.Args[2], os.Args[3]
}

func loadPublicKeyFromPem(publicKeyPemFile string) *rsa.PublicKey {

	fileData, err := ioutil.ReadFile(publicKeyPemFile)
	if err != nil {
		log.Fatal("error loading the public key pem file. ", err)
	}

	block, _ := pem.Decode(fileData)
	if block == nil || block.Type != "PUBLIC KEY" {
		log.Fatal("Unable to load valid public key. ")
	}

	//publicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Fatal("unable to parse pem public key. ", err)
	}

	return publicKey.(*rsa.PublicKey)
}

func loadFile(fileName string) []byte {
	fileData, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal("unable to load file. ", err)
	}

	return fileData
}

func verifySignature(publicKey *rsa.PublicKey, signature, message []byte) bool {
	hashed := sha256.Sum256(message)

	err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed[:], signature)
	if err != nil {
		log.Println(err)
		return false
	}

	return true // if no error then hash is matching
}

func main() {
	publicKeyFile, signatureFile, messageFile := checkArgs()
	fmt.Println(publicKeyFile, signatureFile, messageFile)

	// load the PEM file public key
	publicKey := loadPublicKeyFromPem(publicKeyFile)

	// load files
	signature := loadFile(signatureFile)
	message := loadFile(messageFile)

	// decrypt the digital signature using the public key result SHOULD be the hash of the plain text message
	// hash the plain text message - message.txt compare the result with the decrypted signature
	valid := verifySignature(publicKey, signature, message)

	if valid {
		fmt.Println("Signature Verified")
	} else {
		fmt.Println("Signature could not be verified")
	}
}
