package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/hpcloud/tail"
	"github.com/slack-go/slack"
)

func main() {
	logfile, pattern, message, webhook := validargs(os.Args)

	tailConfig := tail.Config{
		Follow: true,
		Location: &tail.SeekInfo{
			Whence: io.SeekEnd,
		},
	}

	t, err := tail.TailFile(logfile, tailConfig)
	if err != nil {
		fmt.Printf("Cant tail the file: %s\n", err)
		os.Exit(1)
	}

	for line := range t.Lines {
		if strings.Contains(line.Text, pattern) {
			sendMsg(webhook, message)
			os.Exit(0)
		}
	}

}

func validargs(args []string) (logfile string, pattern string, message string, webhook string) {
	if len(args) != 4 {
		help()
		os.Exit(1)
	}

	logfile = os.Args[1]
	pattern = os.Args[2]
	message = os.Args[3]
	webhook = os.Getenv("SLACK_WEBHOOK_URL")

	if _, err := os.Stat(logfile); os.IsNotExist(err) {
		fmt.Printf("The file %s doesn't exist.\n", logfile)
		os.Exit(1)
	}

	if webhook == "" {
		fmt.Printf("You need to set $SLACK_WEBHOOK_URL.\n")
		os.Exit(1)
	}

	return
}

func help() {
	fmt.Printf("boink - a tool to notify you when a line is in a log file.\n\n")
	fmt.Printf("USAGE\n")
	fmt.Printf("   export SLACK_WEBHOOK_URL='https://...'\n")
	fmt.Printf("   boink [logfile] '[Substr pattern]' '[message to send]'\n")
}

func sendMsg(webhook string, message string) {

	msg := slack.WebhookMessage{
		Text: message,
	}

	err := slack.PostWebhook(webhook, &msg)
	if err != nil {
		fmt.Println(err)
	}
}
