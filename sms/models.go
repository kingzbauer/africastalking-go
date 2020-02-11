package sms

// StatusCode corresponds to the status of the request
type StatusCode int

const (
	StatusProcessed             StatusCode = 100
	StatusSent                  StatusCode = 101
	StatusQueued                StatusCode = 102
	StatusRiskHold              StatusCode = 401
	StatusInvalidSenderID       StatusCode = 402
	StatusInvalidPhoneNumber    StatusCode = 403
	StatusUnsupportedNumberType StatusCode = 404
	StatusInsufficientBalance   StatusCode = 405
	StatusUserInBlacklist       StatusCode = 406
	StatusCouldNotRoute         StatusCode = 407
	StatusInternalServerError   StatusCode = 500
	StatusGatewayError          StatusCode = 501
	StatusRejectedByGateway     StatusCode = 502
)

// Request the body of the request to send an SMS
type Request struct {
	Username             string `form:"username"`
	To                   string `form:"to"`
	Message              string `form:"message"`
	From                 string `form:"from"`
	BulkSMSMode          int    `form:"bulkSMSMode"`
	Enqueue              int    `form:"enqueue"`
	Keyword              string `form:"keyword"`
	LinkID               string `form:"linkId"`
	RetryDurationInHours int    `form:"retryDurationInHours"`
}

// Response body of the response, a JSON object
type Response struct {
	Message    string
	Recipients []Recipient
}

// Recipient recipient contained in the original request data showing the status of the sent message
type Recipient struct {
	StatusCode StatusCode `json:"statusCode"`
	Number     string     `json:"number"`
	Status     string     `json:"success"`
	Cost       string     `json:"cost"`
	MessageID  string     `json:"messageId"`
}
