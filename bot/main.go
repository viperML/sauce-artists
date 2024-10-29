package bot

import (
	"log"
	"os"
	"github.com/bwmarrin/discordgo"
)

func Execute() {
	token, found := os.LookupEnv("DISCORD_BOT_TOKEN")
	if !found {
		log.Fatal("DISCORD_BOT_TOKEN is not set")
	}

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err)
	}

	// dg.AddHandler(handler)

	dg.Identify.Intents = discordgo.IntentGuildMessages

	err = dg.Open()
	if err != nil {
		log.Fatal(err)
	}

	channel, found := os.LookupEnv("DISCORD_CHANNEL_ID")
	if !found {
		log.Fatal("DISCORD_CHANNEL_ID not set")
	}

    dg.ChannelMessageSend(channel, "https://danbooru.donmai.us/posts")

	// log.Print("Bot running")
	// sc := make(chan os.Signal, 1)
	// signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	// <-sc

	dg.Close()
}

// func handler(s *discordgo.Session, m *discordgo.MessageCreate) {
//     log.Print(m)
//
//
//     s.ChannelMessageSend(channel, "ping")
// }
