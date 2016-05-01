package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type Config struct {
	Consumer_key        string
	Consumer_secret     string
	Access_token        string
	Access_token_secret string
}

func main() {
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	config_s := Config{}
	err := decoder.Decode(&config_s)

	if err != nil {
		log.Fatal(err)
	}

	config := oauth1.NewConfig(config_s.Consumer_key, config_s.Consumer_secret)
	token := oauth1.NewToken(config_s.Access_token, config_s.Access_token_secret)
	httpClient := config.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)
	tweets, _, err := client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{})
	if err != nil {
		log.Fatal(err)
	}

	for _, t := range tweets {
		fmt.Println(t.Text)
	}
}
