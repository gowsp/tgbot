package tgbot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Config struct {
	Url   string `json:"url,omitempty"`
	Token string `json:"token,omitempty"`
}

func (b *Config) url(method string) string {
	if b.Url == "" {
		b.Url = "https://api.telegram.org"
	}
	return fmt.Sprintf("%s/bot%s/%s", b.Url, b.Token, method)
}

type Client interface {
	Do(req *http.Request) (*http.Response, error)
}
type Cmd interface {
	Method() string
}

func New(config *Config) *Bot {
	return &Bot{config: config, client: http.DefaultClient}
}
func NewWithClient(config *Config, client Client) *Bot {
	return &Bot{config: config, client: client}
}

type Error struct {
	ErrorCode   int
	Description string
}

func (e *Error) Error() string {
	return fmt.Sprintf("error code: %d, description: %s", e.ErrorCode, e.Description)
}

type result struct {
	Ok     bool            `json:"ok"`
	Result json.RawMessage `json:"result"`

	ErrorCode   int    `json:"error_code"`
	Description string `json:"description"`
}

func (r *result) Success() error {
	if r.Ok {
		return nil
	}
	return &Error{ErrorCode: r.ErrorCode, Description: r.Description}
}

type Bot struct {
	config *Config
	client Client
}

func (b *Bot) do(req *http.Request, rsp any) error {
	res, err := b.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	result := new(result)
	err = json.NewDecoder(res.Body).Decode(result)
	if err != nil {
		return err
	}
	if err = result.Success(); err != nil {
		return err
	}
	return json.Unmarshal(result.Result, rsp)
}

// Execute the request
func (b *Bot) Run(cmd Cmd, rsp any) error {
	body := new(bytes.Buffer)
	if err := json.NewEncoder(body).Encode(cmd); err != nil {
		return err
	}
	url := b.config.url(cmd.Method())
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	return b.do(req, rsp)
}
