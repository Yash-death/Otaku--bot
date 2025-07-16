package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type KitsuResponse struct {
	Data []struct {
		Attributes struct {
			PosterImage struct {
				Original string `json:"original"`
			} `json:"posterImage"`
		} `json:"attributes"`
	} `json:"data"`
}

func FetchAnimeImage(animeTitle string) string {
	// Normalize anime name: remove special chars, accents
	cleanedName := strings.NewReplacer(
		":", "",
		"!", "",
		"?", "",
		".", "",
		"é", "e",
		",", "",
	).Replace(animeTitle)

	query := fmt.Sprintf("https://kitsu.io/api/edge/anime?filter[text]=%s", url.QueryEscape(cleanedName))
	defaultImage := "https://tenor.com/bUiwO.gif"

	resp, err := http.Get(query)
	if err != nil {
		log.Println("Kitsu Api Error: ", err)
		return defaultImage
	}
	defer resp.Body.Close()

	var data KitsuResponse

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Println("Kistu decode error:", err)
		return defaultImage
	}

	if len(data.Data) == 0 {
		log.Printf("Kitsu returned no results for anime: %s\n", animeTitle)
		return defaultImage
	}

	return data.Data[0].Attributes.PosterImage.Original
}
