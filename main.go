package main

import (
	"SNiC_MobilIT/network"
	"SNiC_MobilIT/telegram"
	"SNiC_MobilIT/track"
	"context"
	"log"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	bot := telegram.NewBot()

	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	stopNotifier := make(chan bool)
	stopServer := make(chan bool)

	go func() {
		for {
			select {
			case <- stopNotifier:
				return
			case <- ticker.C:
				var count = network.GetUpdate()
				var filtered = track.FilterInterestingTracks(count)
				for _, message := range track.GetInterestingMessages(filtered) {
					err := bot.SendMessage(message)
					if err != nil {
						log.Printf("Sending the message went wrong: %s\n", err)
						return
					}
				}
			}
		}
	}()

	bot.Start(ctx, stopNotifier, stopServer)

	network.StartServer(stopServer)
}

