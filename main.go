package main

import (
	"fmt"
	"log"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func main() {
	fmt.Println("Hello World")

	client, err := google.DefaultClient(oauth2.NoContext,
		"https://www.googleapis.com/auth/devstorage.full_control")
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Get("https://nestjs-example-app.handover.d.bike24.dev")

	fmt.Println(resp)
}
