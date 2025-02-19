// 5. README.md content:
# API Client Package

A simple API client with JWT authentication support.

## Installation

### As a command-line tool
```bash
go install github.com/mohamedtbarka/first-go-package/cmd/apitool@latest
```

### As a package
```bash
go get github.com/mohamedtbarka/first-go-package
```

## Usage

### Command-line tool
```bash
# List all users
apitool -url=http://api.example.com -token=your-jwt-token -action=list

# Get specific user
apitool -url=http://api.example.com -token=your-jwt-token -action=get -id=123

# Create new user
apitool -url=http://api.example.com -token=your-jwt-token -action=create -name="John Doe" -email="john@example.com"
```

### As a package in your code
```go
import "github.com/mohamedtbarka/first-go-package/pkg/client"

func main() {
    // Initialize client
    apiClient := client.NewAPIClient("http://api.example.com", "your-jwt-token")

    // Get all users
    users, err := apiClient.GetUsers()
    if err != nil {
        log.Fatal(err)
    }

    // Get specific user
    user, err := apiClient.GetUserByID("123")
    if err != nil {
        log.Fatal(err)
    }

    // Create new user
    newUser := client.User{
        Name:  "John Doe",
        Email: "john@example.com",
    }
    created, err := apiClient.CreateUser(newUser)
    if err != nil {
        log.Fatal(err)
    }
}
```