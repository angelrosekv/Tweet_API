package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dghubble/oauth1"
)

// Response structure for Twitter API v2 (only relevant fields)
type TwitterResponse struct {
	Data struct {
		ID string `json:"id"`
	} `json:"data"`
}

func main() {
	// OAuth1 config with your API keys
	config := oauth1.NewConfig("uI2mvGNYW1JyP2BseMUe5rlxy", "93gNn4M8JW1uVi9YUpIzKYJMQDESWTnNTzKjWqqK3NPGllKb4K")
	token := oauth1.NewToken("1844361366070505482-Datzvvo16RMBpHQgaUT4haZChotqfA", "i7U4lwEuj1etD8jZcX83qAEqQKBbZ2El18sDfKM5Lhm7k")

	// Create an HTTP client with OAuth
	httpClient := config.Client(oauth1.NoContext, token)

	// Prepare the tweet JSON payload
	payload := map[string]string{
		"text": "Testing tweet posting with the Twitter API v2 and printing tweet ID!",
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// POST request to create a new tweet using API v2
	tweetUrl := "https://api.twitter.com/2/tweets"
	req, err := http.NewRequest("POST", tweetUrl, bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	// Perform POST request
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Println("Error posting tweet:", err)
		return
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != 201 { // 201 Created is expected in v2
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("Error response from Twitter:", resp.Status, string(body))
		return
	}

	// Parse response to get the tweet ID
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var twitterResponse TwitterResponse
	if err := json.Unmarshal(bodyBytes, &twitterResponse); err != nil {
		fmt.Println("Error unmarshalling response:", err)
		return
	}

	// Print the Tweet ID
	fmt.Printf("Tweet posted successfully! Tweet ID: %s\n", twitterResponse.Data.ID)
}
