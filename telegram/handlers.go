package telegram

import (
	"SNiC_MobilIT/network"
	"SNiC_MobilIT/track"
	"fmt"
	"log"
	"os"

	"github.com/yanzay/tbot/v2"
)

//// Send this to BotFather to set commands (with /setcommands)
//help - Get this help message
//update - Get new updates
//info - Get info about all talks
//id - Get your own message ID

// Handle the /start and /help commands here
func (b *Bot) helpHandler(m *tbot.Message) {
	log.Printf("user %s %s with id: %s sent message: '%s'", m.Chat.FirstName, m.Chat.LastName, m.Chat.ID, m.Text)
	msg := `This is a bot whose purpose is to send updates about the capacity of some tracks of the SNiC MobilIT event.
Commands:
1. Use /help to get this help message
2. Use /update to get an update.
3. Use /info to get info about all talks.
4. Use /id to get your message ID.`

	if _, err := b.client.SendMessage(m.Chat.ID, msg); err != nil {
		log.Printf("failed to send unsubscribe message to client: %s", err)
	}
}

// IDHandler Handle the /getID command here
func (b *Bot) IDHandler(m *tbot.Message) {
	if _, err := b.client.SendMessage(m.Chat.ID, "Your ID is: " + m.Chat.ID); err != nil {
		log.Printf("failed to send unsubscribe message to client: %s", err)
	}
}

func (b *Bot) SendMessage(message string) error {
	if _, err := b.client.SendMessage(
		os.Getenv("TELEGRAM_CHAT_ID"),
		message,
	); err != nil {
		return fmt.Errorf("failed to send ad without image to client: %w", err)
	}
	return nil
}

func (b *Bot) updateHandler(m *tbot.Message) {
	err := b.SendMessage("Running...")
	if err != nil {
		log.Printf("Sending the running message went wrong: %s\n", err)
	}

	var count = network.GetUpdate()
	var filtered = track.FilterInterestingTracks(count)
	for _, message := range track.GetInterestingMessages(filtered) {
		err := b.SendMessage(message)
		if err != nil {
			log.Printf("Sending the message went wrong: %s\n", err)
			return
		}
	}

	err = b.SendMessage("Done")
	if err != nil {
		log.Printf("Sending the done message went wrong: %s\n", err)
		return
	}
}

func (b *Bot) infoHandler(m *tbot.Message) {
	err := b.SendMessage("Running...")
	if err != nil {
		log.Printf("Sending the running message went wrong: %s\n", err)
	}

	var count = network.GetUpdate()
	err = b.SendMessage(track.GetAllTracks(count))
	if err != nil {
		log.Printf("Sending the message went wrong: %s\n", err)
		return
	}

	err = b.SendMessage("Done")
	if err != nil {
		log.Printf("Sending the done message went wrong: %s\n", err)
		return
	}
}