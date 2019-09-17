package main

import (
	"fmt"
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

func main() {
	privateKeyFile, messageFile, signatureFile := checkArgs()
	fmt.Printf("%s %s %s", privateKeyFile, messageFile, signatureFile)

}
