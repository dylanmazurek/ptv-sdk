package ptv

import "net/http"

type addAuthTransport struct {
	roundTripper http.RoundTripper
	credentials  *Credentials
}

func (a *addAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	authUrl, err := a.credentials.AuthRequestUrl(*req.URL)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = authUrl.RawQuery

	return a.roundTripper.RoundTrip(req)
}
