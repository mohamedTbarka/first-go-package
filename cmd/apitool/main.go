// 4. cmd/apitool/main.go content:
package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/mohamedtbarka/first-go-package/pkg/client"
)

func main() {
	var (
		baseURL   = flag.String("url", "", "Base URL of the API")
		token     = flag.String("token", "", "JWT token for authentication")
		action    = flag.String("action", "list", "Action to perform: list, get, create")
		userID    = flag.String("id", "", "User ID for get action")
		userName  = flag.String("name", "", "User name for create action")
		userEmail = flag.String("email", "", "User email for create action")
	)

	flag.Parse()

	if *baseURL == "" || *token == "" {
		flag.Usage()
		os.Exit(1)
	}

	apiClient := client.NewAPIClient(*baseURL, *token)

	switch *action {
	case "list":
		users, err := apiClient.GetUsers()
		if err != nil {
			log.Fatalf("Error getting users: %v", err)
		}
		printJSON(users)

	case "get":
		if *userID == "" {
			log.Fatal("User ID is required for get action")
		}
		user, err := apiClient.GetUserByID(*userID)
		if err != nil {
			log.Fatalf("Error getting user: %v", err)
		}
		printJSON(user)

	case "create":
		if *userName == "" || *userEmail == "" {
			log.Fatal("Name and email are required for create action")
		}
		newUser := client.User{
			Name:  *userName,
			Email: *userEmail,
		}
		created, err := apiClient.CreateUser(newUser)
		if err != nil {
			log.Fatalf("Error creating user: %v", err)
		}
		printJSON(created)

	default:
		log.Fatalf("Unknown action: %s", *action)
	}
}

func printJSON(v interface{}) {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(v); err != nil {
		log.Fatalf("Error encoding JSON: %v", err)
	}
}
