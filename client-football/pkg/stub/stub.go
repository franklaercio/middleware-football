package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/franklaercio/middleware-football/stub-middleware"
	"github.com/gookit/color"
)

type Match struct {
	Info struct {
		Date   int64  `json:"timestamp"`
		Status string `json:"status.long"`
	} `json:"fixture"`
	Teams struct {
		Home struct {
			Name   string `json:"name"`
			Winner string `json:"winner"`
		} `json:"home"`
		Away struct {
			Name   string `json:"name"`
			Winner string `json:"winner"`
		} `json:"away"`
	} `json:"teams"`
	Goals struct {
		Home int `json:"home"`
		Away int `json:"away"`
	} `json:"goals"`
}

type MatchDay struct {
	Matches []Match `json:"response"`
}

type MatchService struct {
	service stub.Response
}

func NewMatchService() *MatchService {
	return &MatchService{}
}

func (m *MatchService) GetMatchDay(league string) {
	res, err := m.service.GetStubFromService(league, "2024")
	if err != nil {
		fmt.Println("Error when getting response from service:", err)
	}

	var matchDay MatchDay

	err = json.Unmarshal([]byte(res.Message), &matchDay)
	if err != nil {
		fmt.Println("Error parsing response body:", err)
	}

	fmt.Println("\n" + color.FgGreen.Sprintf("---- Upcoming Matches ----") + "\n")

	for i, match := range matchDay.Matches {
		fmt.Printf("%d. %s x %s\n", i, match.Teams.Home.Name, match.Teams.Away.Name)
	}

	fmt.Println("\n" + color.FgGreen.Sprintf("--------------------------") + "\n")
}
