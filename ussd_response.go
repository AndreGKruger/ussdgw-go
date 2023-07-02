package ussdgwapp

// Enums for Outgoing USSD Response Type
const (
	RESPONSE_TYPE_EXISTING = 2
	RESPONSE_TYPE_END      = 3
	RESPONSE_TYPE_TIMEOUT  = 4
	RESPONSE_TYPE_REDIRECT = 5
)

type UssdResponse struct {
	Type    int    `json:"type"`
	Message string `json:"msg"`
}

func (u *UssdResponse) IsValid() bool {
	return u.ValidateResponseType() && u.ValidateResponseMessage()
}

// ValidateResponseType returns true if the response type is valid
func (u *UssdResponse) ValidateResponseType() bool {
	return u.Type == RESPONSE_TYPE_EXISTING || u.Type == RESPONSE_TYPE_END || u.Type == RESPONSE_TYPE_TIMEOUT || u.Type == RESPONSE_TYPE_REDIRECT
}

// ValidateResponseMessage returns true if the response message is less than 182 characters including whitespaces
func (u *UssdResponse) ValidateResponseMessage() bool {
	return len(u.Message) <= 182
}
