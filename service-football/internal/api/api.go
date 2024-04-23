package api

import (
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
)

type Client struct {
	apiKey  string
	apiHost string
}

func NewClient() (*Client, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("RAPID_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("API_KEY environment variable not set")
	}

	apiHost := os.Getenv("RAPID_API_HOST")
	if apiHost == "" {
		return nil, fmt.Errorf("API_HOST environment variable not set")
	}

	return &Client{
		apiKey:  apiKey,
		apiHost: apiHost,
	}, nil
}

func (c *Client) FetchData(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-RapidAPI-Key", c.apiKey)
	req.Header.Add("X-RapidAPI-Host", c.apiHost)

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Error reading response body:", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		log.Println("Error:", resp.StatusCode)
		return nil, err
	}

	body, _ := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return nil, err
	}

	return body, nil
}
