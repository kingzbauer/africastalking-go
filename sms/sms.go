package sms

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"strings"

	"github.com/kingzbauer/africastalking-go/client"
)

// ErrJSONDecode thrown when an error is encountered when decoding a payload
// Include the payload
type ErrJSONDecode struct {
	parent  error
	Payload []byte
}

// Error implements the error interface
func (err ErrJSONDecode) Error() string {
	return err.parent.Error()
}

// Unwrap for compatibility with errors.Unwrap method
func (err ErrJSONDecode) Unwrap() error {
	return err.parent
}

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
	bodyContent, err := ioutil.ReadAll(repBody)
	if err != nil {
		return nil, err
	}
	// Close the body if is supports it
	if closer, ok := repBody.(io.Closer); ok {
		closer.Close()
	}

	if err := json.Unmarshal(bodyContent, rep); err != nil {
		// Use ErrJSONDecode to provide the caller with the opportunity to inspect the payload
		err = &ErrJSONDecode{
			parent:  err,
			Payload: bodyContent,
		}
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
