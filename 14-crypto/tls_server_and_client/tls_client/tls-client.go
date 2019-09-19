package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"os"
)

func printUsage() {
	fmt.Println(os.Args[0] + ` - Send and receive a message to a TLS server

Usage:
  ` + os.Args[0] + ` <hostString>

Example:
  ` + os.Args[0] + ` localhost:9999
`)
}

func checkArgs() string {
	if len(os.Args) != 2 {
		printUsage()
		os.Exit(1)
	}

	// Host string e.g. localhost:9999
	return os.Args[1]
}

func main() {
	hostString := checkArgs()
	messageToSend := "Hello\n"

	// Configure TLS settings
	tlsConfig := &tls.Config{
		// Required to accept self-signed certs
		InsecureSkipVerify: true,
		// Provide your client certificate if necessary
		// Certificates: []Certificate

		// ServerName is used to verify the hostname (unless you are skipping verification)
		// It is also included in the handshake in case the server uses virtual hosts Can also just be an IP address instead of a hostname.
		// ServerName: string,

		// RootCAs that you are willing to accept
		// If RootCAs is nil, the host's default root CAs are used
		// RootCAs: *x509.CertPool
	}

	serverConnection, err := tls.Dial("tcp", hostString, tlsConfig)
	if err != nil {
		log.Fatal("error connecting to server. ", err)
	}
	defer serverConnection.Close()

	// Write Data to socket
	numBytesWritten, err := serverConnection.Write([]byte(messageToSend))
	if err != nil {
		log.Fatal("error writing to socket. ", err)
	}
	fmt.Printf("Wrote %d bytes to the socket.\n", numBytesWritten)

	// Read data from socket and print to STDOUT
	buffer := make([]byte, 100)
	numBytesRead, err := serverConnection.Read(buffer)
	if err != nil {
		log.Fatal("error reading bytes from socket. ", err)
	}
	fmt.Printf("Read %d bytes to the socket.\n", numBytesRead)
	fmt.Printf("Message Received %s", buffer)
}
