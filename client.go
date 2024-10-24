package passkey

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type requester struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
}

type Client struct {
	tenantID  string
	requester *requester
}

func New(baseURL, apiKey, tenantID string) *Client {
	r := &requester{
		baseURL: baseURL,
		apiKey:  apiKey,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}

	return &Client{requester: r, tenantID: tenantID}
}

func (r *requester) do(
	ctx context.Context,
	method string,
	path string,
	body interface{},
	result interface{},
) error {
	var payload []byte
	if body != nil {
		if jb, err := json.Marshal(body); nil != err {
			return err
		} else {
			payload = jb
		}
	}

	fullURL := fmt.Sprintf("%s/%s",
		r.baseURL,
		path)

	req, err := http.NewRequestWithContext(
		ctx,
		method,
		fullURL,
		bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("ApiKey", r.apiKey)

	resp, err := r.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	code := resp.StatusCode
	if code < 200 || code >= 300 {
		var errResp string
		if err := dec.Decode(&errResp); nil != err {
			return err
		}
		return errors.New(errResp)
	}

	return dec.Decode(result)
}
