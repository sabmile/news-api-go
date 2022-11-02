package news

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	http     *http.Client
	apiKey   string
	PageSize int
}

type Article struct {
	Source struct {
		ID   interface{} `json:"id"`
		Name string      `json:"name"`
	} `json:"source"`
	Author      string    `json:"author"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	URLToImage  string    `json:"urlToImage"`
	PublishedAt time.Time `json:"publishedAt"`
	Content     string    `json:"content"`
}

type Results struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
}

func NewClient(httpClient *http.Client, apiKey string, pageSize int) *Client {
	return &Client{httpClient, apiKey, pageSize}
}

func (c *Client) FetchEverything(query string) (*Results, error) {
	q := "https://newsapi.org/v2/everything?q=%s&apiKey=%s&pageSize=%d"
	endPoint := fmt.Sprintf(q, url.QueryEscape(query), c.apiKey, c.PageSize)
	resp, err := c.http.Get(endPoint)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}

	res := &Results{}

	return res, json.Unmarshal(body, res)
}
