package main

import (
	"discraft/collection"
	"discraft/discord"
	"discraft/minecraft"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func requireEnvVariable(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("%s is mandatory", key))
	}

	return value
}

func stringToURL(s string) url.URL {
	u, err := url.Parse(s)
	if err != nil {
		panic(fmt.Sprintf("'%s' is not a valid URL", s))
	}

	return *u
}

func stringToMinecraftServerURL(s string) url.URL {
	if !strings.Contains(s, "//") {
		s = "//" + s
	}

	u := stringToURL(s)
	if u.Port() == "" {
		u.Host = u.Host + ":25565"
	}

	return u
}

func main() {
	discordWebhookURL := stringToURL(requireEnvVariable("DISCORD_WEBHOOK_URL"))
	minecraftServerUrl := stringToMinecraftServerURL(requireEnvVariable("MINECRAFT_SERVER_URL"))

	discordServer := discord.Server{
		Client:  http.Client{},
		Webhook: discordWebhookURL,
	}
	minecraftServer := minecraft.NewServerFromURL(minecraftServerUrl)

	var lastPlayers *collection.Strings
	ticker := time.NewTicker(1 * time.Minute)
	for {
		currentPlayers := minecraftServer.CurrentPlayerList()
		if lastPlayers == nil {
			lastPlayers = &currentPlayers
		}

		newPlayers := currentPlayers.Diff(*lastPlayers)
		oldPlayers := lastPlayers.Diff(currentPlayers)

		for _, player := range newPlayers {
			discordServer.SendMessage(fmt.Sprintf("%s has entered. Currently playing: %v", player, currentPlayers))
		}

		for _, player := range oldPlayers {
			discordServer.SendMessage(fmt.Sprintf("%s left. Currently playing: %v", player, currentPlayers))
		}

		lastPlayers = &currentPlayers

		<-ticker.C
	}
}
