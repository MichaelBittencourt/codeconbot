package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {

	token := "<TOKEN DO DISCORD>"
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalln(err)
	}

	dg.AddHandler(messageCreate)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("O bot esta em execução, pressione CTRL+C para sair")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()

}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	if m.Content == "!quemsoueu" {
		s.ChannelMessageSend(m.ChannelID, "O bot mais lindo da Codecon 2021!")
	}

	if m.Content == "/cmd" {
		s.ChannelMessageSend(m.ChannelID, "Novo comando!")
	}

}
