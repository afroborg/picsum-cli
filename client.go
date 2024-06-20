package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Client struct {
	baseUrl string
	options *Flags
	client  *http.Client
}

func NewClient(baseUrl string, options *Flags) *Client {
	return &Client{baseUrl: baseUrl, options: options, client: &http.Client{}}
}

// func (c *Client) get(endpoint string) (*http.Response, error) {
// 	req, err := http.NewRequest("GET", c.baseUrl+endpoint, nil)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return c.client.Do(req)
// }

func (c *Client) buildPath() string {
	var b strings.Builder

	if c.options.Id != "" {
		b.WriteString("/id/" + c.options.Id)
	} else if c.options.Seed != "" {
		b.WriteString("/seed/" + c.options.Seed)
	}

	b.WriteString(fmt.Sprintf("/%d/%d", c.options.Width, c.options.Height))

	return b.String()
}

func (c *Client) buildQueryParams() url.Values {
	q := url.Values{}

	if c.options.Grayscale {
		q.Add("grayscale", "true")
	}

	if c.options.Blur != 0 {
		q.Add("blur", fmt.Sprintf("%d", c.options.Blur))
	}

	return q
}

func (c *Client) GetImage() io.ReadCloser {
	path := c.buildPath()
	query := c.buildQueryParams()

	url := c.baseUrl + path + "?" + query.Encode()

	res, err := c.client.Get(url)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if res.StatusCode != http.StatusOK {
		fmt.Println("Failed to fetch image")
		os.Exit(1)
	}

	fmt.Println("Image fetched successfully")

	return res.Body
}
