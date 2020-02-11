package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kingzbauer/africastalking-go/sms"
)

var (
	apiKey    = flag.String("k", "", "apiKey provided by AT")
	username  = flag.String("u", "", "username provided by AT")
	shortCode = flag.String("s", "", "Short code registered with your AT app")
	live      = flag.Bool("l", false, "Whether to make a live api call. Default is sandbox")
	message   = flag.String("m", "This is a test message", "Message to send")
	number    = flag.String("p", "", "Phone number to receive the message")
)

// parseFromEnv fills in the apiKey and username if they are blank from environmental variables
func parseFromEnv(apiKey, username *string) {
	if len(*apiKey) == 0 {
		*apiKey = os.Getenv("AT_API_KEY")
	}

	if len(*username) == 0 {
		*username = os.Getenv("AT_USERNAME")
	}
}

func main() {
	flag.Parse()

	parseFromEnv(apiKey, username)
	// apiKey and username are compulsory values
	if len(*apiKey) == 0 || len(*username) == 0 || len(*number) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	srv := sms.NewService(*apiKey, *username, *shortCode, *live)
	rep, err := srv.Send(*message, []string{*number}, "")

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Response: %v\n", rep)
}
