package network

import (
	"SNiC_MobilIT/track"

	"encoding/json"
	"log"
	"net/http"
)

func GetUpdate() track.Count {
	resp, err := http.Get("https://mobilit.snic.nl/api/talk/count/")
	if err != nil {
		log.Printf("Getting results from snic went wrong: %v", err)
	}
	defer resp.Body.Close()

	countResult := &track.Count{Content: map[string]track.Track{}}

	err = json.NewDecoder(resp.Body).Decode(&countResult)
	if err != nil {
		log.Printf("Decoding json went wrong: %s\n", err)
		return track.Count{}
	}

	return *countResult
}