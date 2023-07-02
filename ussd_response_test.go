package ussdgwapp

import (
	"testing"
)

func TestValidateResponseType(t *testing.T) {
	//Test if the response type is valid
	want := true
	u := UssdResponse{
		Type: RESPONSE_TYPE_EXISTING,
	}
	got := u.ValidateResponseType()
	if got != want {
		t.Errorf("UssdResponse.ValidateResponseType() = %v, want %v", got, want)
	}

	//Test if the response type is invalid
	want = false
	u = UssdResponse{
		Type: 7,
	}
	got = u.ValidateResponseType()
	if got != want {
		t.Errorf("UssdResponse.ValidateResponseType() = %v, want %v", got, want)
	}
}

func TestValidateResponseMessage(t *testing.T) {
	//Test if the response message is valid
	want := true
	u := UssdResponse{
		Message: "This is a valid response message",
	}
	got := u.ValidateResponseMessage()
	if got != want {
		t.Errorf("UssdResponse.ValidateResponseMessage() = %v, want %v", got, want)
	}

	//Test if the response message is invalid
	want = false
	u = UssdResponse{
		Message: "This is an invalid reponse message that is way too long ...............................................................................................................................",
	}
	got = u.ValidateResponseMessage()
	if got != want {
		t.Errorf("UssdResponse.ValidateResponseMessage() = %v, want %v", got, want)
	}
}
