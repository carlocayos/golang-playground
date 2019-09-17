package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"strconv"
)

func printUsage() {
	fmt.Printf(os.Args[0] + `

Generate a private and public RSA keypair and save as PEM files.
If no key size is provided, a default of 2048 is used.

Usage:
  ` + os.Args[0] + ` <private_key_filename> <public_key_filename> [keysize]

Examples:
  # Store generated private and public key in privkey.pem and pubkey.pem
  ` + os.Args[0] + ` priv.pem pub.pem
  ` + os.Args[0] + ` priv.pem pub.pem 4096`)
}

func checkArgs() (string, string, int) {

	if len(os.Args) < 2 || len(os.Args) > 4 {
		printUsage()
		os.Exit(1)
	}

	// 2048 bits or 256 bytes
	defaultKeySize := 2048

	// If there are 2 args provided, privkey and pubkey filenames
	if len(os.Args) == 3 {
		return os.Args[1], os.Args[2], defaultKeySize
	}

	// If 3 args provided, privkey, pubkey, keysize
	if len(os.Args) == 4 {
		customKeySize, err := strconv.Atoi(os.Args[3])
		if err != nil {
			log.Fatal("invalid key size. ", err)
		}
		return os.Args[1], os.Args[2], customKeySize
	}

	return "", "", 0
}

// Encode keys to PEM format
//privatePem := getPrivatePemFromKey(privateKey)
//publicPem := generatePublicPemFromKey(privateKey.PublicKey)

// Encode the private key as a PEM file
// PEM is a base-64 encoding of the key
func getPrivatePemFromKey(privateKey *rsa.PrivateKey) *pem.Block {

	encodedPrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
	var privatePem = &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: encodedPrivateKey,
	}

	return privatePem
}

// Encode the public key as a PEM file
func generatePublicPemFromKey(publicKey rsa.PublicKey) *pem.Block {

	encodedPublicKey, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		log.Fatal("error converting public key", err)
	}

	var publicPem = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: encodedPublicKey,
	}

	return publicPem
}

func savePemToFile(pemBlock *pem.Block, filename string) {

	// save PEM to output file
	publicPemOutputFile, err := os.Create(filename)
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer publicPemOutputFile.Close()

	err = pem.Encode(publicPemOutputFile, pemBlock)
	if err != nil {
		log.Fatal("unable to encode file", err)
	}
}

func main() {
	privatePemFileName, publicPemFileName, keySize := checkArgs()
	fmt.Println(privatePemFileName, publicPemFileName, keySize)

	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		log.Fatal("error generating key. ", err)
	}

	// Encode keys to PEM format
	privatePem := getPrivatePemFromKey(privateKey)
	publicPem := generatePublicPemFromKey(privateKey.PublicKey)

	savePemToFile(privatePem, "private.pem")
	savePemToFile(publicPem, "public.pem")
}
