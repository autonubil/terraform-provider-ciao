package ciao

import (
	"crypto/tls"
	"fmt"

	resty "github.com/go-resty/resty/v2"
)

// The Check strucure of ciao
type Check struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	URL           string `json:"url"`
	Cron          string `json:"cron"`
	Active        bool   `json:"active"`
	Status        string `json:"status,omitempty"`
	Job           string `json:"job"`
	LastContactAt string `json:"last_contact_at,omitempty"`
	NextContactAt string `json:"next_contact_at,omitempty"`
	CreatedAt     string `json:"created_at,omitempty"`
	UpdatedAt     string `json:"updated_at,omitempty"`
	CheckURL      string `json:"check_url,omitempty"`
}

// The Client for the ciao REST API
type Client struct {
	client   *resty.Client
	baseURL  string
	user     string
	password string
	insecure bool
}

// NewClient returns a new ciao API client
func NewClient(baseURL string, user string, password string) *Client {
	c := &Client{
		client:  resty.New(),
		baseURL: baseURL,
	}
	c.client.SetHostURL(baseURL)

	// Basic Auth for all request
	if (user != "") && (password != "") {
		c.client.SetBasicAuth(user, password)
	}

	// Headers for all request
	c.client.SetHeader("Accept", "application/json")
	c.client.SetHeaders(map[string]string{
		"Content-Type": "application/json",
		"User-Agent":   "terraform",
	})

	// Registering global Error object structure for JSON/XML request
	// c.client.SetError(&Error{}) // or resty.SetError(Error{})
	return c
}

/**
enable debugging
*/
func (c *Client) SetDebug(enable bool) {
	c.client.SetDebug(enable)
}

/**
Allow insecure HTTPS Connections
*/
func (c *Client) SetInsecure(enable bool) {
	c.client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: enable})
}

// NewCheck creates a new check on the ciao server
func (c *Client) NewCheck(chk *Check) (*Check, error) {
	response, err := c.client.R().
		SetBody(chk).
		SetResult(&Check{}).
		Post(fmt.Sprintf("%s/checks.json", c.baseURL))
	if err != nil {
		return nil, err
	}
	if response.StatusCode() > 399 {
		return nil, fmt.Errorf("Request to ciao failed: %s", response.Status())
	}

	return response.Result().(*Check), nil
}

// UpdateCheck updates an existing check on the ciao server
func (c *Client) UpdateCheck(ID string, chk *Check) (*Check, error) {
	response, err := c.client.R().
		SetPathParams(map[string]string{
			"id": ID,
		}).
		SetBody(chk).
		SetResult(&Check{}).
		Put(fmt.Sprintf("%s/checks/{id}.json", c.baseURL))
	if err != nil {
		return nil, err
	}
	if response.StatusCode() > 399 {
		return nil, fmt.Errorf("Request to ciao failed: %s", response.Status())
	}

	return response.Result().(*Check), nil
}

// ReadCheck returns the check for a given ID
func (c *Client) ReadCheck(ID string) (*Check, error) {
	response, err := c.client.R().
		SetPathParams(map[string]string{
			"id": ID,
		}).
		SetResult(&Check{}).
		Get(fmt.Sprintf("%s/checks/{id}.json", c.baseURL))
	if err != nil {
		return nil, err
	}
	if response.StatusCode() > 399 {
		return nil, fmt.Errorf("Request to ciao failed: %s", response.Status())
	}
	return response.Result().(*Check), nil
}

// DeleteCheck deletes the check for a given ID
func (c *Client) DeleteCheck(ID string) error {
	response, err := c.client.R().
		SetPathParams(map[string]string{
			"id": ID,
		}).
		SetResult(&Check{}).
		Delete(fmt.Sprintf("%s/checks/{id}.json", c.baseURL))
	if err != nil {
		return err
	}
	if response.StatusCode() > 399 {
		return fmt.Errorf("Request to ciao failed: %s", response.Status())
	}
	return nil
}
