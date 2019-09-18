package main

import (
	"crypto"
	"crypto/rand"
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

Cryptographically sign a message using a private key.
Private key should be a PEM encoded RSA key.
Signature is generated using SHA256 hash.
Output signature is stored in filename provided.

Usage:
  ` + os.Args[0] + ` <privateKeyFilename> <messageFilename>   <signatureFilename>

Example:
  # Use priv.pem to encrypt msg.txt and output to sig.txt.256
  ` + os.Args[0] + ` priv.pem msg.txt sig.txt.256
`)
}

func checkArgs() (string, string, string) {
	if len(os.Args) != 4 {
		printUsage()
		os.Exit(1)
	}

	return os.Args[1], os.Args[2], os.Args[3]
}

// Load the message that will be signed from file
func loadMessageFromFile(messageFile string) []byte {
	message, err := ioutil.ReadFile(messageFile)
	if err != nil {
		log.Fatal("unable to open message file", err)
	}
	return message
}

func loadPrivateKeyFromPemFile(privateKeyFile string) *rsa.PrivateKey {
	// load file to memory
	fileData, err := ioutil.ReadFile(privateKeyFile)
	if err != nil {
		log.Fatal("unable to open private key file", err)
	}

	// Get the block data from the PEM encoded file
	block, _ := pem.Decode(fileData)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		log.Fatal("Unable to load a valid private key.")
	}

	// Parse the bytes and put it in to a proper privateKey struct
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Fatal("error loading private key. ", err)
	}

	//-------------------------------------------------
	// START of PUBLIC KEY PEM Creation
	//
	// NOTE: This part is added just to get the public key from existing PEM file
	//  and save it to a separate Public Key PEM file. This pubilc key PEM file
	//  is needed to verify the signature in the next lesson.
	// See "rsa_sha256_digital_signature_verification"
	//-------------------------------------------------
	encodedPublicKey, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		log.Fatal("error converting public key. ", err)
	}

	publicKeyPemblock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: encodedPublicKey,
	}

	// save pem to file
	// write to file
	publicPemFile, err := os.OpenFile("public.pem", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal("error creating public PEM file", err)
	}
	defer publicPemFile.Close()

	err = pem.Encode(publicPemFile, publicKeyPemblock)
	if err != nil {
		log.Fatal("error saving public pem file. ", err)
	}

	//-------------------------------------------------
	// END of PUBLIC KEY PEM Creation
	//-------------------------------------------------

	return privateKey
}

// Cryptographically sign a message= creating a digital signature
// of the original message. Uses SHA-256 hashing.
func signMessage(privateKey *rsa.PrivateKey, message []byte) []byte {

	hashedData := sha256.Sum256(message)
	fmt.Printf("SHA256 hash = %x\n", hashedData)

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashedData[:])
	if err != nil {
		log.Fatal("error signing message. ", err)
	}

	return signature
}

// Save data to file
func saveToFile(signatureFile string, signedMessage []byte) error {

	file, err := os.OpenFile(
		signatureFile,
		os.O_RDWR|os.O_CREATE,
		0666,
	)

	if err != nil {
		return err
	}
	defer file.Close()

	// Write bytes to file
	_, err = file.Write(signedMessage)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	privateKeyFile, messageFile, signatureFile := checkArgs()
	fmt.Printf("%s %s %s\n", privateKeyFile, messageFile, signatureFile)

	// Load message and private key files from disk
	message := loadMessageFromFile(messageFile)
	privateKey := loadPrivateKeyFromPemFile(privateKeyFile)
	fmt.Printf("%s\n", message)
	fmt.Printf("%x\n", privateKey)

	// Cryptographically sign the message
	signedMessage := signMessage(privateKey, message)

	// Write the signed message to file
	saveToFile(signatureFile, signedMessage)

	//------------------------------------
	// Try to verify signature here
	//------------------------------------
	hashedMessage := sha256.Sum256(message)
	err := rsa.VerifyPKCS1v15(&privateKey.PublicKey, crypto.SHA256, hashedMessage[:], signedMessage)
	if err != nil {
		fmt.Println("signature is NOT verified")
	} else {
		fmt.Println("signature is verified")
	}

}
