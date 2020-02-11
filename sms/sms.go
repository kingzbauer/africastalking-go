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
