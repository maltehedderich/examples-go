package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		println("Error:", err)
		os.Exit(1)
	}
	// The Read interface (io.Reader) is a fundamental interface in Go that defines
	// how data is read from a source. It has a single method:
	//   Read(p []byte) (n int, err error)
	//
	// How it works:
	// - We pass in a byte slice (p) that acts as a container
	// - Read fills this slice with data from the source (up to len(p) bytes)
	// - It returns the number of bytes read (n) and any error encountered
	//
	// Here, resp.Body implements the Reader interface, allowing us to read
	// the HTTP response body data into our byte slice (bs)
	bs := make([]byte, 99999)
	resp.Body.Read(bs)

	fmt.Println(string(bs))
}
