package main

import (
	"fmt"
	"net/http"

	"github.com/dghubble/oauth1"
)

func main() {
	// OAuth1 config with your API keys
	config := oauth1.NewConfig("uI2mvGNYW1JyP2BseMUe5rlxy", "93gNn4M8JW1uVi9YUpIzKYJMQDESWTnNTzKjWqqK3NPGllKb4K")
	token := oauth1.NewToken("1844361366070505482-Datzvvo16RMBpHQgaUT4haZChotqfA", "i7U4lwEuj1etD8jZcX83qAEqQKBbZ2El18sDfKM5Lhm7k")

	// Create an HTTP client with OAuth
	httpClient := config.Client(oauth1.NoContext, token)

	// Replace with the Tweet ID you want to delete
	tweetID := "1844444433774764514" // Example tweet ID, replace with your own tweet's ID

	// URL for deleting the tweet
	deleteUrl := fmt.Sprintf("https://api.twitter.com/1.1/statuses/destroy/%s.json", tweetID)
	req, err := http.NewRequest("POST", deleteUrl, nil)
	if err != nil {
		fmt.Println("Error creating delete request:", err)
		return
	}

	// Perform the delete request
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Println("Error deleting tweet:", err)
		return
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != 200 {
		fmt.Printf("Failed to delete the tweet. Status: %s\n", resp.Status)
		return
	}

	fmt.Println("Tweet deleted successfully!")
}
