package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/dghubble/oauth1"
)

func main() {
	// OAuth1 config with your API keys
	config := oauth1.NewConfig("uI2mvGNYW1JyP2BseMUe5rlxy", "93gNn4M8JW1uVi9YUpIzKYJMQDESWTnNTzKjWqqK3NPGllKb4K")
	token := oauth1.NewToken("1844361366070505482-Datzvvo16RMBpHQgaUT4haZChotqfA", "i7U4lwEuj1etD8jZcX83qAEqQKBbZ2El18sDfKM5Lhm7k")

	// Create an HTTP client with OAuth
	httpClient := config.Client(oauth1.NoContext, token)

	reader := bufio.NewReader(os.Stdin)
	for {
		// Get user input for the Tweet ID
		fmt.Print("Enter the Tweet ID you want to delete (or type 'exit' to quit): ")
		tweetIDStr, _ := reader.ReadString('\n')

		// Trim the newline character and any extra spaces
		tweetIDStr = strings.TrimSpace(tweetIDStr)

		// Exit condition
		if tweetIDStr == "exit" {
			fmt.Println("Exiting the program.")
			break
		}

		// Check for empty Tweet ID
		if tweetIDStr == "" {
			fmt.Println("Tweet ID cannot be empty. Please try again.")
			continue
		}

		// Convert tweetID from string to int64
		tweetID, err := strconv.ParseInt(tweetIDStr, 10, 64)
		if err != nil {
			fmt.Println("Error: Invalid Tweet ID format. Please enter a numeric Tweet ID.")
			continue
		}

		// URL for deleting the tweet
		deleteUrl := fmt.Sprintf("https://api.twitter.com/1.1/statuses/destroy/%d.json", tweetID)
		req, err := http.NewRequest("POST", deleteUrl, nil)
		if err != nil {
			fmt.Println("Error creating delete request:", err)
			continue
		}

		// Perform the delete request
		resp, err := httpClient.Do(req)
		if err != nil {
			fmt.Println("Error deleting tweet:", err)
			continue
		}
		defer resp.Body.Close()

		// Handle different response status codes
		body, _ := ioutil.ReadAll(resp.Body)
		switch resp.StatusCode {
		case 200:
			fmt.Println("Tweet deleted successfully!")
		case 401:
			fmt.Println("Error: Unauthorized. Please check your API credentials.")
		case 403:
			fmt.Println("Error: Forbidden. You may not have the required permissions.")
		case 404:
			fmt.Println("Error: Tweet not found. Please check the tweet ID.")
		case 429:
			fmt.Println("Error: Rate limit exceeded. Please wait before trying again.")
		default:
			fmt.Printf("Error: Failed to delete tweet. Status: %d. Response: %s\n", resp.StatusCode, string(body))
		}
	}
}
