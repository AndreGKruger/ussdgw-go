package ussdgwapp

// Enums for Incoming USSD Request Type
const (
	REQUEST_TYPE_NEW      = 1
	REQUEST_TYPE_EXISTING = 2
	REQUEST_TYPE_END      = 3
	REQUEST_TYPE_TIMEOUT  = 4
	REQUEST_TYPE_REDIRECT = 5
	REQUEST_TYPE_USSD_NI  = 6
)

// UssdRequest struct
type UssdRequest struct {
	SessionId string `json:"sessionid"`
	Msisdn    uint64 `json:"msisdn"`
	Type      int    `json:"type"` // 1: New, 2: Existing, 3: End, 4: Timeout, 5: Redirect, 6: USSD NI
	Message   string `json:"msg"`  // urlencoded
}

// IsNew returns true if the request is a new session
func (u *UssdRequest) IsNew() bool {
	return u.Type == REQUEST_TYPE_NEW
}

// IsExisting returns true if the request is an existing session
func (u *UssdRequest) IsExisting() bool {
	return u.Type == REQUEST_TYPE_EXISTING
}

// IsEnd returns true if the request is an end session
func (u *UssdRequest) IsEnd() bool {
	return u.Type == REQUEST_TYPE_END
}

// IsTimeout returns true if the request is a timeout session
func (u *UssdRequest) IsTimeout() bool {
	return u.Type == REQUEST_TYPE_TIMEOUT
}

// GetSessionId returns the session id
func (u *UssdRequest) GetSessionId() string {
	return u.SessionId
}
