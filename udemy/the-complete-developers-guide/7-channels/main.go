package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// List of URLs to check for availability
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.com",
		"http://amazon.com",
	}

	// Create a channel for communicating between goroutines
	// Channels are typed conduits through which goroutines can send and receive values
	// This channel will carry string values (the URLs)
	c := make(chan string)

	// Launch a goroutine for each link to check them concurrently
	// The 'go' keyword runs the function in a separate goroutine (lightweight thread)
	// This allows all links to be checked simultaneously rather than sequentially
	for _, link := range links {
		go checkLink(link, c)
	}

	// Infinite loop that receives messages from the channel
	// 'range c' waits for values to be sent to the channel and receives them one at a time
	// This loop blocks until a value is available, then processes it
	for l := range c {
		// Launch a new goroutine to re-check the link after a delay
		// This creates a continuous monitoring loop for each URL
		go func() {
			// Pause for 5 seconds before re-checking the link
			// This prevents hammering the servers with constant requests
			time.Sleep(5 * time.Second)
			// Re-check the link and send the result back to the channel
			checkLink(l, c)
		}()
	}
}

// checkLink performs an HTTP GET request to verify if a URL is accessible
// It takes a link to check and a channel to communicate the result back
func checkLink(link string, c chan string) {
	// Attempt to fetch the URL
	// We ignore the response body (using _) since we only care if the request succeeds
	_, err := http.Get(link)
	if err != nil {
		// If there's an error (timeout, DNS failure, connection refused, etc.)
		fmt.Println(link, "might be down!")
		// Send the link back to the channel so it can be re-checked
		// The '<-' operator sends a value into the channel
		c <- link
		return
	}
	// If the request was successful, the site is up
	fmt.Println(link, "is up!")
	// Send the link back to the channel regardless of success or failure
	// This ensures the monitoring loop continues
	c <- link
}
