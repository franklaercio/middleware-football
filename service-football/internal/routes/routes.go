package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"service-football/internal/api"
)

type CurrentRoundResponse struct {
	Response []string `json:"response"`
}

type Client struct {
	api.Client
}

type Match struct {
	Info        MatchInfo   `json:"fixture"`
	TeamDetails TeamDetails `json:"teams"`
	GoalStats   GoalStats   `json:"goals"`
}

type MatchInfo struct {
	Date   int64 `json:"timestamp"`
	Status struct {
		Long string `json:"long"`
	} `json:"status"`
}

type GoalStats struct {
	Home int `json:"home"`
	Away int `json:"away"`
}

type TeamDetails struct {
	Home struct {
		Name   string `json:"name"`
		Winner string `json:"winner"`
	} `json:"home"`
	Away struct {
		Name   string `json:"name"`
		Winner string `json:"winner"`
	} `json:"away"`
}

type MatchDay struct {
	Match []Match `json:"response"`
}

func NewClient() (*Client, error) {
	c, err := api.NewClient()
	if err != nil {
		return nil, err
	}

	return &Client{*c}, nil
}

func (c *Client) GetCurrentRound(league, season string) (*CurrentRoundResponse, error) {
	baseURL := "https://api-football-v1.p.rapidapi.com/v3/fixtures/rounds?league=%s&season=%s&current=true"
	if league == "" {
		return nil, fmt.Errorf("tournamentID cannot be empty")
	}
	if season == "" {
		return nil, fmt.Errorf("seasonID cannot be empty")
	}

	url := fmt.Sprintf(baseURL, league, season)
	res, err := c.FetchData(url)
	if err != nil {
		log.Println("Error reading response body:", err)
		return nil, err
	}

	var round CurrentRoundResponse

	err = json.Unmarshal(res, &round)
	if err != nil {
		log.Println("Error parsing response body:", err)
		return nil, err
	}

	return &round, nil
}

func (c *Client) GetEvents(league, season, round string) (*MatchDay, error) {
	baseURL := "https://api-football-v1.p.rapidapi.com/v3/fixtures?league=%s&season=%s&round=%s"
	if league == "" || season == "" || round == "" {
		return nil, fmt.Errorf("error when validate parameters")
	}

	url := fmt.Sprintf(baseURL, league, season, round)
	res, err := c.FetchData(url)
	if err != nil {
		log.Println("Error reading response body:", err)
		return nil, err
	}

	var matchDay MatchDay

	err = json.Unmarshal(res, &matchDay)
	if err != nil {
		log.Println("Error parsing response body:", err)
		return nil, err
	}

	return &matchDay, nil
}
