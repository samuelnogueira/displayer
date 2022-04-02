package main

import (
	"fmt"
	"github.com/alteamc/minequery/ping"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

func grabEnvConfiguration() (discordWebhookURL string, minecraftHost string, minecraftPort uint16) {
	discordWebhookURL, ok := os.LookupEnv("DISCORD_WEBHOOK_URL")
	if !ok {
		panic("DISCORD_WEBHOOK_URL is mandatory")
	}

	minecraftHost, ok = os.LookupEnv("MINECRAFT_HOST")
	if !ok {
		panic("MINECRAFT_HOST is mandatory")
	}

	minecraftPortStr, ok := os.LookupEnv("MINECRAFT_PORT")
	if !ok {
		panic("MINECRAFT_PORT is mandatory")
	}

	minecraftPort64, err := strconv.ParseUint(minecraftPortStr, 10, 16)
	if err != nil {
		panic("MINECRAFT_PORT is not valid")
	}

	minecraftPort = uint16(minecraftPort64)

	return
}

func main() {
	discordWebhookURL, minecraftHost, minecraftPort := grabEnvConfiguration()

	var lastPlayers *stringList
	client := http.Client{}
	ticker := time.NewTicker(1 * time.Minute)
	for {
		log.Printf("Pinging %s:%d...", minecraftHost, minecraftPort)
		res, err := ping.Ping(minecraftHost, minecraftPort)
		if err != nil {
			panic(err)
		}
		log.Printf("Ping OK: %v", res)

		currentPlayers := newStringListFromPlayerNames(*res)
		if lastPlayers == nil {
			lastPlayers = &currentPlayers
		}

		newPlayers := currentPlayers.diff(*lastPlayers)
		oldPlayers := lastPlayers.diff(currentPlayers)
		content := ""

		if len(newPlayers) > 0 {
			content = fmt.Sprintf("%v entered, now standing: %v", newPlayers, currentPlayers)
		}

		if len(oldPlayers) > 0 {
			content = fmt.Sprintf("%v left, currently standing: %v", oldPlayers, currentPlayers)
		}

		if content != "" {
			log.Printf("Posting '%s' to discord...", content)
			resp, err := client.PostForm(
				discordWebhookURL,
				url.Values{
					"content": {content},
				},
			)
			err = resp.Body.Close()
			if err != nil {
				panic(err)
			}
		}

		lastPlayers = &currentPlayers

		<-ticker.C
	}
}
