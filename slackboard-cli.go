package main

import (
	"./slackboard"
	"bytes"
	"flag"
	"io"
	"log"
	"os"
)

func main() {

	version := flag.Bool("v", false, "slackboard version")
	server := flag.String("s", "", "slackboard server name")
	tag := flag.String("t", "", "slackboard tag name")
	sync := flag.Bool("sync", false, "enable synchronous notification")
	flag.Parse()

	if *version {
		slackboard.PrintVersion()
		os.Exit(0)
	}

	if *server == "" && *tag == "" {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *server == "" {
		log.Fatal("Specify slackboard server name")
	}

	if *tag == "" {
		log.Fatal("Specify slackboard tag name")
	}

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "localhost"
	}

	var text bytes.Buffer
	io.Copy(&text, os.Stdin)
	payload := &slackboard.SlackboardPayload{
		Tag:  *tag,
		Host: hostname,
		Text: text.String(),
		Sync: *sync,
	}

	err = slackboard.SendNotification2Slackboard(*server, payload)
	if err != nil {
		log.Fatal(err.Error())
	}
}
