package main

import (
	"fmt"
	"log"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	os.Setenv("Slack-Bot-Token", "xoxb-5222282569271-5239366232677-ZDsAf5vk0cRdpZU1i55TJiaB")
	// os.Setenv("Slack-App-Token", "")
	os.Setenv("Channel-Id", "C056J8AJYT1")

	api := slack.New(os.Getenv("Slack-Bot-Token"))
	channelArr := []string{os.Getenv("Channel-id")}
	fileArr := []string{"QB"}

	// we could have more than one file so we will be looping through it
	for i := 0; i < len(fileArr); i++ {
		// it has two parameters channels and file
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
			// we can have multiple files
		}
		file, err := api.UploadFile(params)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Name %s , URL: %s \n", file.Name, file.URL)

	}
}
