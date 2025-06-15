package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const baseUrl string = "https://pokeapi.co/api/v2/location-area/"

type RespLocationsBatch struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type Client struct {
	httpClient http.Client
}

func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

func (c Client) ListLocations(pageUrl *string) (RespLocationsBatch, error) {
	url := baseUrl
	if pageUrl != nil {
		url = *pageUrl
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationsBatch{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationsBatch{}, err
	}
	if res.StatusCode > 299 {
		return RespLocationsBatch{}, fmt.Errorf("error getting location areas response")
	}
	defer res.Body.Close()

	var data RespLocationsBatch
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&data); err != nil {
		return RespLocationsBatch{}, err
	}

	return data, nil
}
