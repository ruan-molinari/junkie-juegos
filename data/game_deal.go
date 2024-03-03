package data

import (
	"encoding/json"
	"fmt"
	"time"
)

type GameDeal struct {
	Id              int        `json:"id"`
	Title           string     `json:"title"`
	Worth           string     `json:"worth"`
	Thumbnail       string     `json:"thumbnail"`
	Image           string     `json:"image"`
	Description     string     `json:"description"`
	Instructions    string     `json:"instructions"`
	OpenGiveawayUrl string     `json:"open_giveaway_url"`
	PublishedDate   *time.Time `json:"published_date"`
	Type            string     `json:"type"`
	Platforms       string     `json:"platforms"`
	EndDate         *time.Time `json:"end_date"`
	Users           int        `json:"users"`
	Status          string     `json:"status"`
	GamerpowerUrl   string     `json:"gamerpower_url"`
	OpenGiveaway    string     `json:"open_giveaway"`
}

// Custom unmarshal function for type `GameDeal`.
//
// Needed because the Gamerpower API returns a Datetime
// format different from time.Time.
//
// Also, non pre-defined datetime format on go time library
// matches this specific format
func (g *GameDeal) UnmarshalJSON(data []byte) error {
	// Definid an alias for our GameDeal struct with the
	// datetime fields as tring.
	type GameDealAlias GameDeal
	tmp := struct {
		PublishedDate string `json:"published_date"`
		EndDate       string `json:"end_date"`
		*GameDealAlias
	}{
		GameDealAlias: (*GameDealAlias)(g),
	}

	err := json.Unmarshal(data, &tmp)
	if err != nil {
		fmt.Println("Error unmarshalling json: ", err)
		return err
	}

	// Datetime format returned by the Gamerpower API
	const dateTimeFormat = "2006-01-02 15:04:05"

	// Converting datetime strings to time.Time in the original GameDeal struct
	publishedDate, err := time.Parse(dateTimeFormat, tmp.PublishedDate)
	if err != nil {
		g.PublishedDate = nil
	} else {
		g.PublishedDate = &publishedDate
	}

	endDate, err := time.Parse(dateTimeFormat, tmp.EndDate)
	if err != nil {
		g.EndDate = nil
	} else {
		g.EndDate = &endDate
	}

	return nil
}
