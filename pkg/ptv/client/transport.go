package client

import "net/http"

type AuthTransport struct {
	roundTripper http.RoundTripper
	credentials  Credentials
}

func NewAuthTransport(roundTripper http.RoundTripper, credentials Credentials) *AuthTransport {
	newAuthTransport := &AuthTransport{
		roundTripper: roundTripper,
		credentials:  credentials,
	}

	return newAuthTransport
}

func (a *AuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	authUrl, err := a.credentials.AuthRequestUrl(*req.URL)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = authUrl.RawQuery

	return a.roundTripper.RoundTrip(req)
}
