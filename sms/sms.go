package sms

import (
	"encoding/json"
	"strings"

	"github.com/kingzbauer/africastalking-go/client"
)

// NewRequest creates a new request object with the minimal required fields.
// You can leave `from` as empty string, defaults to AFRICASTKNG
func NewRequest(message string, to []string, from string) *Request {
	return &Request{
		Message: message,
		To:      strings.Join(to, ","),
		From:    from,
	}
}

// SendMessage given a client and request, makes an API call to AT
func SendMessage(cli *client.Client, req *Request) (rep *Response, err error) {
	repBody, err := cli.Do(req, client.V1EndpointMessaging)
	if err != nil {
		return nil, err
	}

	rep = &Response{}
	if err := json.NewDecoder(repBody).Decode(rep); err != nil {
		return nil, err
	}

	return
}

// Service exposes sms based functions
type Service struct {
	cli       *client.Client
	shortCode string
}

// NewService initialises the sms service
func NewService(apiKey, username, defaultShortCode string, sandbox bool) *Service {
	cli := client.New(apiKey, username, sandbox)
	return &Service{
		cli:       cli,
		shortCode: defaultShortCode,
	}
}

// Send makes a send api call and returns the response.
// `to` is an array of phone numbers in the form +2547********
// `shortCode` overrides the default short code if provided
func (s *Service) Send(message string, to []string, shortCode string) (rep *Response, err error) {
	if len(shortCode) == 0 {
		shortCode = s.shortCode
	}

	req := NewRequest(message, to, shortCode)
	return SendMessage(s.cli, req)
}
