
# Twitter API Interaction Assignment

## Introduction
This assignment demonstrates how to interact with the Twitter API by creating and deleting tweets using a go program. The goal of this project is to help students understand how to work with external APIs, use OAuth authentication, and manage API responses. Students will perform tasks such as posting a tweet and deleting a tweet programmatically while handling errors and rate limiting.

By completing this assignment, students will:
- Learn how to authenticate with the Twitter API using OAuth.
- Gain experience making POST requests to create a new tweet.
- Understand how to make DELETE requests to remove an existing tweet.
- Practice writing well-documented code.
  
## Setup Instructions

### 1. Set Up a Twitter Developer Account
Before you can interact with the Twitter API, you need to create a developer account and generate API keys:
- Create a [Twitter Developer Account](https://developer.twitter.com/).
- Once your account is approved, create a new project and an app under your Twitter Developer Dashboard.
  
### 2. Generate API Keys
Follow these steps to generate the necessary API keys:
1. Navigate to **Projects & Apps** → **Your App** → **Keys and Tokens**.
2. Generate the following credentials:
   - **API Key** (Consumer Key)
   - **API Secret Key** (Consumer Secret)
   - **Bearer Token**
   - **Access Token**
   - **Access Token Secret**

Make sure to store these credentials securely as you will need them in your program.

### 3. Configure OAuth
- Set the **Callback URL** to `http://localhost:3000` (or any local testing environment).
- Enable **User Authentication Settings** under the App settings to allow you to interact with your Twitter account.

### 4. Install Go and Required Libraries
1. In your project directory, initialize a Go module:
  
   go mod init tweet

Install all the necessary pakckages
- go get github.com/dghubble/go-twitter/twitter
- go get github.com/dghubble/oauth1


## Program Details

### Posting a New Tweet
The program uses Tweepy, a go library for accessing the Twitter API, to post a tweet. The function `post_tweet` handles this functionality by sending a POST request to the **statuses/update** endpoint.

### To run the program 
To run program Add  Twitter API credentials (API Key, API Secret Key, Access Token, and Access Token Secret) to the go script and run go script.

go run post2.go

![image](https://github.com/user-attachments/assets/6aabd455-4e63-4eb5-a778-a5d9194075f9)

We can see the post on the twitter as below

![image](https://github.com/user-attachments/assets/d5466614-29b3-44e6-915f-9b6f1174ee70)


### Deleting an existing Account 
To delete an existing tweet, the program uses the POST statuses/destroy/:id endpoint. For deleting required to give the tweet ID .

To delete the existing tweet post we need to give the tweet ID and i run below command

go run delete.go

![image](https://github.com/user-attachments/assets/63b3e668-2d87-4880-9a17-44c54fbe560f)

In the twitter account you can see the post is deleted as shown in the below

![image](https://github.com/user-attachments/assets/ec7475bc-6b70-43f8-a065-eec243d55765)


### API Response 
Posting a Tweet: A JSON response with details of the newly created tweet  
Deleting a Tweet: A JSON response confirming the deletion of the tweet.

### Error handling 
The program includes error handling to manage various issues that may arise when interacting with the Twitter API, such as:

- Invalid credentials: The program checks if your API keys and tokens are correct.
- Rate limiting: Twitter has rate limits for API requests. If you exceed these limits, the program will catch the error and handle it gracefully.
- Invalid tweet ID: When deleting a tweet, the program ensures that a valid tweet ID is provided. If an invalid ID is used, the program catches the error and prints an appropriate message.

![image](https://github.com/user-attachments/assets/835c528d-f11f-4c80-8136-de5f958a66d6)

