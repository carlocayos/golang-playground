package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
)

func printUsage() {
	fmt.Println("Usage: " + os.Args[0] + "<filepath>")
	fmt.Println("Example: " + os.Args[0] + " document.txt")
}

func checkArgs() string {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
	return os.Args[1]
}

// Run "hash_small_file document_large.txt"
func main() {
	filename := checkArgs()

	// Openfile
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	hasher := md5.New()

	_, err = io.Copy(hasher, file)
	if err != nil {
		log.Fatal(err)
	}

	checksum := hasher.Sum(nil)
	fmt.Printf("MD5 checksum: %x\n", checksum)
}
