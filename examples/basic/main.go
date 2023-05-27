package main

import (
	"log"
	"strings"

	"github.com/gempir/go-twitch-irc/v4"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	c := twitch.NewAnonymousClient()

	c.OnPrivateMessage(func(msg twitch.PrivateMessage) {
		if strings.Contains(strings.ToLower(msg.Message), "ping") {
			log.Println(msg.User.Name, "PONG", msg.Message)
		}
	})

	c.Join("testaccount_420")

	must(c.Connect())
}
