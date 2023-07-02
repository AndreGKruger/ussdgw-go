package ussdgwapp

import (
	"testing"
)

func TestUssdRequest_IsNew(t *testing.T) {
	//Test if the request is a new session
	want := true
	u := UssdRequest{
		Type: REQUEST_TYPE_NEW,
	}
	got := u.IsNew()
	if got != want {
		t.Errorf("UssdRequest.IsNew() = %v, want %v", got, want)
	}

	//Test if the request is not a new session
	want = false
	u = UssdRequest{
		Type: REQUEST_TYPE_EXISTING,
	}
	got = u.IsNew()
	if got != want {
		t.Errorf("UssdRequest.IsNew() = %v, want %v", got, want)
	}
}

func TestUssdRequest_IsExisting(t *testing.T) {
	//Test if the request is an existing session
	want := true
	u := UssdRequest{
		Type: REQUEST_TYPE_EXISTING,
	}
	got := u.IsExisting()
	if got != want {
		t.Errorf("UssdRequest.IsExisting() = %v, want %v", got, want)
	}

	//Test if the request is not an existing session
	want = false
	u = UssdRequest{
		Type: REQUEST_TYPE_NEW,
	}
	got = u.IsExisting()
	if got != want {
		t.Errorf("UssdRequest.IsExisting() = %v, want %v", got, want)
	}
}

func TestUssdRequest_IsEnd(t *testing.T) {
	//Test if the request is an end session
	want := true
	u := UssdRequest{
		Type: REQUEST_TYPE_END,
	}
	got := u.IsEnd()
	if got != want {
		t.Errorf("UssdRequest.IsEnd() = %v, want %v", got, want)
	}

	//Test if the request is not an end session
	want = false
	u = UssdRequest{
		Type: REQUEST_TYPE_NEW,
	}
	got = u.IsEnd()
	if got != want {
		t.Errorf("UssdRequest.IsEnd() = %v, want %v", got, want)
	}
}

func TestUssdRequest_IsTimeout(t *testing.T) {
	//Test if the request is a timeout session
	want := true
	u := UssdRequest{
		Type: REQUEST_TYPE_TIMEOUT,
	}
	got := u.IsTimeout()
	if got != want {
		t.Errorf("UssdRequest.IsTimeout() = %v, want %v", got, want)
	}

	//Test if the request is not a timeout session
	want = false
	u = UssdRequest{
		Type: REQUEST_TYPE_NEW,
	}
	got = u.IsTimeout()
	if got != want {
		t.Errorf("UssdRequest.IsTimeout() = %v, want %v", got, want)
	}
}

func TestUssdRequest_GetSessionId(t *testing.T) {
	//Test if the request returns the session id
	want := "1234567890"
	u := UssdRequest{
		SessionId: "1234567890",
	}
	got := u.GetSessionId()
	if got != want {
		t.Errorf("UssdRequest.GetSessionId() = %v, want %v", got, want)
	}

	//Test if the request does not return the session id
	want = "1234567890"
	u = UssdRequest{
		SessionId: "0987654321",
	}
	got = u.GetSessionId()
	if got == want {
		t.Errorf("UssdRequest.GetSessionId() = %v, want %v", got, want)
	}
}
