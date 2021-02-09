package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"./cmds"
	dg "github.com/bwmarrin/discordgo"
)

var (
	config = make(map[string]string)
)

func init() {
	arr, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatalf("Error while reading configuration file")
	}

	if err := json.Unmarshal(arr, &config); err != nil {
		log.Fatalf("Incorrect configuration file format")
	}
}

func main() {
	s, err := dg.New(config["token"])
	if err != nil {
		log.Fatalf("Error creating Discord session")
	}

	s.AddHandler(messageCreate)

	//s.Identify.Intents = dg.IntentsGuildMessages

	if err := s.Open(); err != nil {
		log.Fatalf("Error openning connection")
	}
	defer s.Close()

	log.Printf("Bot is running...")

	wait := make(chan os.Signal)
	signal.Notify(wait, syscall.SIGINT)
	<-wait
}

func messageCreate(s *dg.Session, m *dg.MessageCreate) {
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case *dg.MessageEmbed:
				s.ChannelMessageSendEmbed(m.ChannelID, err.(*dg.MessageEmbed))
			default:
				log.Printf("%v", err)
			}
		}
	}()

	if !strings.HasPrefix(m.Content, config["prefix"]) || m.Author.Bot {
		for _, user := range m.Mentions {
			if user.ID == s.State.User.ID {
				s.ChannelMessageSend(m.ChannelID, "Че надо? Мой префикс: "+config["prefix"])
				break
			}
		}

		return
	}

	args := strings.Fields(m.Content[len(config["prefix"]):])
	if cmd, ok := cmds.Map[args[0]]; ok {
		cmd.Exec(s, m, args...)
	}
}
