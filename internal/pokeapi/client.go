package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/brendenwelch/pokedex/internal/pokecache"
)

const baseUrl string = "https://pokeapi.co/api/v2/location-area/"

type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(60 * time.Second),
	}
}

func (c *Client) ListLocations(pageUrl *string) (RespLocationsBatch, error) {
	url := baseUrl
	if pageUrl != nil {
		url = *pageUrl
	}

	_, exists := c.cache.Get(url)
	if !exists {
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

		entry, err := io.ReadAll(res.Body)
		if err != nil {
			return RespLocationsBatch{}, err
		}
		c.cache.Add(url, entry)
	}

	entry, exists := c.cache.Get(url)
	if !exists {
		return RespLocationsBatch{}, fmt.Errorf("error getting cache entry")
	}

	var data RespLocationsBatch
	if err := json.Unmarshal(entry, &data); err != nil {
		return RespLocationsBatch{}, err
	}

	return data, nil
}

func (c *Client) GetLocation(location *string) (RespLocation, error) {
	if location == nil {
		return RespLocation{}, fmt.Errorf("error getting location. no location provided")
	}

	url := baseUrl + *location

	_, exists := c.cache.Get(url)
	if !exists {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespLocation{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return RespLocation{}, err
		}
		if res.StatusCode > 299 {
			return RespLocation{}, fmt.Errorf("error getting location areas response")
		}
		defer res.Body.Close()

		entry, err := io.ReadAll(res.Body)
		if err != nil {
			return RespLocation{}, err
		}
		c.cache.Add(url, entry)
	}

	entry, exists := c.cache.Get(url)
	if !exists {
		return RespLocation{}, fmt.Errorf("error getting cache entry")
	}

	var data RespLocation
	if err := json.Unmarshal(entry, &data); err != nil {
		return RespLocation{}, err
	}

	return data, nil
}
