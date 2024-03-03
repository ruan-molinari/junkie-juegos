package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ruan-molinari/junkie-juegos/data"
)

const api = "https://www.gamerpower.com/api/giveaways?platform=steam&type=game"

func FetchGames() ([]data.GameDeal, error) {
	var gameDeals []data.GameDeal

	res, err := http.Get(api)
	if err != nil {
		fmt.Println("Error getting from api: ", err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body: ", err)
	}

	if err = json.Unmarshal(body, &gameDeals); err != nil {
		fmt.Println("Error decoding data: ", err)
		return nil, err
	}
	return gameDeals, err
}
