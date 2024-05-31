package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
)

func main() {
	// Configure the SSH client
	config := &ssh.ClientConfig{
		User: "pansu",
		Auth: []ssh.AuthMethod{
			ssh.Password("Shivam@098"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Connect to the remote server
	client, err := ssh.Dial("tcp", "192.168.152.206:22", config)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer client.Close()

	// Create a new session
	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}
	defer session.Close()

	// Run the mpstat command on the remote server
	output, err := session.Output(`ipconfig | findstr "IPv4"`)
	if err != nil {
		log.Fatalf("Failed to run mpstat: %v", err)
	}

	// Print the output to the local terminal
	fmt.Printf("%s", output)
}
