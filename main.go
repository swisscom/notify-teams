package main

import (
	"github.com/alexflint/go-arg"
	goteamsnotify "github.com/atc0005/go-teams-notify/v2"
	"github.com/sirupsen/logrus"
)

var args struct {
	WebhookUrl string `arg:"-u,required"`
	Message    string `arg:"-m,required"`
	Title      string `arg:"-t,required"`
}

func main() {
	arg.MustParse(&args)
	mstClient := goteamsnotify.NewClient()
	err := mstClient.Send(args.WebhookUrl, goteamsnotify.MessageCard{
		Type:       "MessageCard",
		Title:      args.Title,
		Text:       args.Message,
		ThemeColor: "ff0000",
	})
	if err != nil {
		logrus.Fatal(err)
		return
	}
}
