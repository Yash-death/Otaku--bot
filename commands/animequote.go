package commands

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/slack-go/slack"
	"github.com/yash-death/otaku-bot/utils"
)

type AnimeQuote struct {
	Status string `json:"status"`
	Data   struct {
		Content string `json:"content"`
		Anime   struct {
			Name string `json:"name"`
		} `json:"anime"`
		Character struct {
			Name string `json:"name"`
		} `json:"character"`
	} `json:"data"`
}

type KitsuResponse struct {
	Data []struct {
		Attributes struct {
			CanonicalTitle string `json:"canonicalTitle"`
			PosterImage    struct {
				Original string `json:"original"`
			} `json:"posterImage"`
		} `json:"attributes"`
	} `json:"data"`
}

func FetchAnimeQuoteBlock() []slack.Block {
	resp, err := http.Get("https://api.animechan.io/v1/quotes/random")
	if err != nil {
		log.Println("API error:", err)
		return fallbackBlock("😢 Failed to fetch quote...")
	}
	defer resp.Body.Close()

	var result AnimeQuote
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Println("Decode error:", err)
		return fallbackBlock("😵 Couldn't decode quote...")
	}

	animeName := result.Data.Anime.Name
	quote := result.Data.Content
	character := result.Data.Character.Name

	imgURL := utils.FetchAnimeImage(animeName)

	emojis := []string{"🎴", "⚔️", "🔥", "✨", "💥", "🌀", "👺"}
	emoji := emojis[rand.Intn(len(emojis))]

	header := slack.NewSectionBlock(
		slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("%s *%s*", emoji, animeName), false, false), nil, nil,
	)
	imageBlock := slack.NewImageBlock(imgURL, animeName+" cover", "", nil)

	quoteBlock := slack.NewSectionBlock(
		slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("> _%s_", quote), false, false), nil, nil,
	)
	authorBlock := slack.NewContextBlock("",
		slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("— *%s*", character), false, false))
	divider := slack.NewDividerBlock()

	return []slack.Block{header, imageBlock, quoteBlock, authorBlock, divider}

}

func fallbackBlock(message string) []slack.Block {
	return []slack.Block{
		slack.NewSectionBlock(slack.NewTextBlockObject("mrkdwn", message, false, false), nil, nil),
	}
}
