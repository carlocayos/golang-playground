package main

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
)

// Print the usage of executable file
func printUsage() {
	fmt.Printf("Usage: " + os.Args[0] + " <filename>\n")
	fmt.Printf("Example: " + os.Args[0] + " document.txt\n")
}

// Check args and return parameter
func checkArgs() string {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
	return os.Args[1]
}

// generate salt using random bytes
func generateSalt() string {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		log.Fatal(err)
	}
	return base64.URLEncoding.EncodeToString(randomBytes)
}

var secretKey = "this should be a secret"

// hash password with the salt
func hashPassword(plainText, salt string) string {

	h := hmac.New(sha256.New, []byte(secretKey))
	_, err := io.WriteString(h, plainText+salt)
	if err != nil {
		log.Fatal(err)
	}
	hashedVal := h.Sum(nil)
	fmt.Printf("Hashed Value = %x\n", hashedVal)
	return hex.EncodeToString(hashedVal)
}

// run as "hmac Password!"
func main() {
	password := checkArgs()
	fmt.Println("Password = " + password)

	salt := generateSalt()
	fmt.Println("Salt = " + salt)

	result := hashPassword(password, salt)
	fmt.Println("Hashed Value returned from hashPassword() = " + result)
}
