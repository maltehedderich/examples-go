package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Check if a filename argument was provided on the command line
	// os.Args[0] is the program name, so we need at least 2 arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		fmt.Println("Example: go run main.go testfile.txt")
		os.Exit(1)
	}

	// Open the file specified by the first command line argument
	// Returns a file handle and any error encountered
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}

	// Copy the contents of the file to standard output
	// io.Copy reads from the file (io.Reader) and writes to os.Stdout (io.Writer)
	io.Copy(os.Stdout, file)
}
