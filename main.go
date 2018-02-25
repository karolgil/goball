package main

import (
	"fmt"
	"os"

	"github.com/nlopes/slack"
)

func main() {
	api := slack.New(os.Getenv("SLACK_TOKEN"))
	params := slack.PostMessageParameters{}
	_, _, err := api.PostMessage("#general",
		"tryout", params)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
}
