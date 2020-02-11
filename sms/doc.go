/*
Package sms provides SMS sending and fetching API

Example Send SMS:

	package main

	import (
		"fmt"
		"os"

		"github.com/kingzbauer/africastalking-go/sms"
	)

	var (
		apiKey string = "yourapikey"
		username string = "yourusername"
		shortCode string = "yourshortcode"
		live bool = true
		number = "+2547********"
	)

	func main() {
		srv := sms.NewService(apiKey, username, shortCode, live)
		rep, err := srv.Send("Test message", []string{number}, shortCode)

		if err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}

		fmt.Printf("Response: %s\n", rep)
	}
*/
package sms
