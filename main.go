package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
	"github.com/yash-death/otaku-bot/commands"
)

func main() {
	_ = godotenv.Load()

	appToken := os.Getenv("SLACK_APP_TOKEN")
	botToken := os.Getenv("SLACK_BOT_TOKEN")

	api := slack.New(
		botToken,
		slack.OptionDebug(true),
		slack.OptionLog(log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)),
		slack.OptionAppLevelToken(appToken),
	)

	client := socketmode.New(
		api,
		socketmode.OptionDebug(true),
		socketmode.OptionLog(log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)),
	)

	go func() {
		for evt := range client.Events {
			switch evt.Type {

			case socketmode.EventTypeSlashCommand:
				cmd, ok := evt.Data.(slack.SlashCommand)
				if !ok {
					log.Println("invalid Slash Command")
					continue
				}

				client.Ack(*evt.Request)

				if cmd.Command == "/animequote" {
					blocks := commands.FetchAnimeQuoteBlock()

					api.PostMessage(cmd.ChannelID, slack.MsgOptionBlocks(blocks...))
				}

			case socketmode.EventTypeEventsAPI:
				event, _ := evt.Data.(slackevents.EventsAPIEvent)

				// Acknowledge the event to Slack
				client.Ack(*evt.Request)

				switch event.Type {
				case slackevents.CallbackEvent:
					innerEvent := event.InnerEvent

					switch ev := innerEvent.Data.(type) {
					case *slackevents.AppMentionEvent:
						_, _, err := api.PostMessage(ev.Channel, slack.MsgOptionText("🎌 Nani?! You called me?", false))
						if err != nil {
							log.Println("error posting message:", err)
						}
					}
				}

			default:
				log.Printf("Ignored: %s\n", evt.Type)
			}
		}
	}()

	log.Println("Otaku Bot is live 🎌")
	client.Run()
}
