package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		println("Error:", err)
		os.Exit(1)
	}

	lw := logWriter{}

	// The Reader interface (io.Reader) is a fundamental interface in Go that defines
	// how data is read from a source. It has a single method:
	//   Read(p []byte) (n int, err error)
	//
	// How it works:
	// - Read takes a byte slice (p) as input and fills it with data from the source
	// - It returns the number of bytes read (n) and any error encountered
	// - When there's no more data, it returns io.EOF error
	//
	// In this code:
	// - resp.Body implements the Reader interface (it's an io.ReadCloser)
	// - logWriter implements the Writer interface
	// - io.Copy uses the Read interface internally: it repeatedly calls Read on
	//   resp.Body to get chunks of data, then writes them to os.Stdout
	// - This pattern allows us to stream data efficiently without loading everything
	//   into memory at once
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
