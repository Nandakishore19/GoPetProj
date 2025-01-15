package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

// "fmt"
// "log"
// "petProj/src/repositories"

// "google.golang.org/grpc"
// // "time"
// "petProj/src/attendanceProtos"
// "petProj/src/server"
// "net"
// "petProj/src/client"


func main() {

	bcFile, err := os.Open("bc.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer bcFile.Close()
	rpbcFile, err := os.Open("rpbc.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer rpbcFile.Close()

	macRegex := regexp.MustCompile(`([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})`)

	extractedMacAddresses := []string{}
	// Read the contents of the bc.txt file
	scanner := bufio.NewScanner(bcFile)
	for scanner.Scan() {
		line := scanner.Text()
		macAddresses := macRegex.FindAllString(line, -1)

		if len(macAddresses) > 0 {
			extractedMacAddresses = append(extractedMacAddresses, macAddresses...)
		}
	}

	// Read the contents of the rpbc.txt file
	scanner = bufio.NewScanner(rpbcFile)
	for scanner.Scan() {
		line := scanner.Text()
		macAddresses := macRegex.FindAllString(line, -1)

		if len(macAddresses) > 0 {
			extractedMacAddresses = append(extractedMacAddresses, macAddresses...)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Extracted MAC addresses:", extractedMacAddresses)
}