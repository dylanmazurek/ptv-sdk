package ptv

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/dylanmazurek/ptv-sdk/pkg/ptv/client"
	"github.com/dylanmazurek/ptv-sdk/pkg/ptv/constants"
)

type Client struct {
	baseURL        string
	internalClient *http.Client
	timezone       *time.Location
}

func New(ctx context.Context, opts ...Option) *Client {
	clientOptions := DefaultOptions()
	for _, opt := range opts {
		opt(&clientOptions)
	}

	creds := client.Credentials{
		Key:    clientOptions.AccessKey,
		UserID: clientOptions.UserID,
	}

	newTransport := client.NewAuthTransport(
		http.DefaultTransport,
		creds,
	)

	httpClient := &http.Client{
		Transport: newTransport,
		Timeout:   constants.DEFAULT_TIMEOUT,
	}

	return &Client{
		baseURL:        clientOptions.BaseURL,
		internalClient: httpClient,
		timezone:       clientOptions.Timezone,
	}
}

func (c *Client) NewRequest(method string, path string, body io.Reader, params *url.Values) (*http.Request, error) {
	reqUrlStr := fmt.Sprintf("%s%s", c.baseURL, path)
	requestUrl, err := url.Parse(reqUrlStr)
	if err != nil {
		return nil, err
	}

	if params != nil && len(*params) > 0 {
		query := requestUrl.Query()
		for key, values := range *params {
			for _, value := range values {
				query.Add(key, value)
			}
		}

		requestUrl.RawQuery = query.Encode()
	}

	req, err := http.NewRequest(method, requestUrl.String(), body)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) Do(req *http.Request, respObj any) error {
	resp, err := c.internalClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return errors.New("unexpected status code: " + resp.Status)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bodyBytes, &respObj)
	if err != nil {
		return err
	}

	return nil
}
