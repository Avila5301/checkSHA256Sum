package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"strings"
)

func calculateSHA256Sum(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	sum := hash.Sum(nil)
	return fmt.Sprintf("%x", sum), nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the SHA256 sum: ")
	inputSHA256Sum, _ := reader.ReadString('\n')
	inputSHA256Sum = strings.TrimSpace(inputSHA256Sum)

	fmt.Print("Enter the path of the file to compare: ")
	downloadedFilePath, _ := reader.ReadString('\n')
	downloadedFilePath = strings.TrimSpace(downloadedFilePath)

	// Calculate the SHA256 checksum of the downloaded file
	actualChecksum, err := calculateSHA256Sum(downloadedFilePath)
	if err != nil {
		fmt.Printf("Error calculating SHA256 sum for downloaded file: %v\n", err)
		os.Exit(1)
	}

	// Compare the input SHA256SUM with the actual checksum
	if inputSHA256Sum == actualChecksum {
		fmt.Println("SHA256 sums match. The file is authentic.")
	} else {
		fmt.Println("SHA256 sums do not match. The file may be corrupted or tampered with.")
	}
}
