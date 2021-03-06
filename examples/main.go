package main

import (
	"fmt"
	"log"

	"github.com/ardikabs/shorty/kutt"
)

func main() {
	api := kutt.API{
		BaseURL:      "https://kutt.it",
		Token:        "YOUR_SECRET_API_TOKEN_NICE",
		CustomDomain: "custom.example.com",
	}

	customURL := "rgoogler"

	url, err := api.SubmitURL("https://google.com", customURL, "", false)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Target: %s\n", url.Target)
	fmt.Printf("ShortURL: %s\n", url.ShortURL)

	urls, err := api.GetListURL()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(urls)

	err = api.DeleteURL(customURL)
	if err != nil {
		log.Fatalln(err)
	}

}
