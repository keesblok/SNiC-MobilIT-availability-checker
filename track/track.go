package track

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type Track struct {
	Count int
	Limit int
}

type Count struct {
	Success bool
	Content map[string]Track
}

func GetInterestingMessages(count Count) []string {
	var messages []string

	for t, trackInfo := range count.Content {
		message := "There "

		freePlaces := trackInfo.Limit - trackInfo.Count

		if freePlaces == 1 {
			message += "is " + strconv.Itoa(freePlaces) + " place left"
		} else {
			message += "are " + strconv.Itoa(freePlaces) + " places left"
		}
		message += " for the talk: " + t + "\n"

		messages = append(messages, message)
	}

	log.Printf("Got %v new messages", len(messages))

	return messages
}

func GetAllTracks(count Count) string {
	message := ""

	for t, trackInfo := range count.Content {
		message += t + ": " + strconv.Itoa(trackInfo.Count) + " of " + strconv.Itoa(trackInfo.Limit) + " places are gone.\n"
	}

	return message
}

func FilterInterestingTracks(input Count) Count {
	interestingTracks := strings.Split(os.Getenv("INTERESTING_TRACKS"), ",")
	filteredList := &Count{Content: map[string]Track{}}

	for _, track := range interestingTracks {
		if trackInfo, ok := input.Content[track]; ok {
			if trackInfo.Count < trackInfo.Limit {
				filteredList.Content[track] = trackInfo
			}
		}
	}
	return *filteredList
}