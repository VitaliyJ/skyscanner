package skyscanner

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

const (
	BaseURL    = "https://partners.api.skyscanner.net/apiservices/v3"
	AuthHeader = "x-api-key"
)

type client struct {
	cfg *Config
}

// NewClient returns new SkyScanner client instance
func NewClient(cfg *Config) Client {
	if cfg.QueriesTimeout == 0 {
		cfg.QueriesTimeout = time.Second * 15
	}

	return &client{
		cfg: cfg,
	}
}

// Create does a create request
func (c client) Create(req *CreateRequest) (*CreatePollResponse, error) {
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	r, err := c.do("POST", "/flights/live/search/create", jsonData)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = r.Body.Close()
	}()

	if r.StatusCode != http.StatusOK {
		var errREsp ErrorResponse
		if err := json.NewDecoder(r.Body).Decode(&errREsp); err != nil {
			return nil, badResponseStatus("create", err, r)
		}

		return nil, errors.New(errREsp.Message)
	}

	var resp CreatePollResponse
	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Poll does a poll request
func (c client) Poll(req *PollRequest) (*CreatePollResponse, error) {
	uri := "/flights/live/search/poll/" + req.SessionToken
	r, err := c.do("POST", uri, []byte{})
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = r.Body.Close()
	}()

	if r.StatusCode != http.StatusOK {
		var errREsp ErrorResponse
		if err := json.NewDecoder(r.Body).Decode(&errREsp); err != nil {
			return nil, badResponseStatus("poll", err, r)
		}

		return nil, errors.New(errREsp.Message)
	}

	var resp CreatePollResponse
	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Locales retrieves the locales that we support to translate your content
func (c client) Locales() (*LocalesResponse, error) {
	r, err := c.do("GET", "/culture/locales", []byte{})
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = r.Body.Close()
	}()

	if r.StatusCode != http.StatusOK {
		var errREsp ErrorResponse
		if err := json.NewDecoder(r.Body).Decode(&errREsp); err != nil {
			return nil, badResponseStatus("locales", err, r)
		}

		return nil, errors.New(errREsp.Message)
	}

	var resp LocalesResponse
	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Currencies retrieves the currencies that Skyscanner support and information about format
func (c client) Currencies() (*CurrenciesResponse, error) {
	r, err := c.do("GET", "/culture/currencies", []byte{})
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = r.Body.Close()
	}()

	if r.StatusCode != http.StatusOK {
		var errREsp ErrorResponse
		if err := json.NewDecoder(r.Body).Decode(&errREsp); err != nil {
			return nil, badResponseStatus("poll", err, r)
		}

		return nil, errors.New(errREsp.Message)
	}

	var resp CurrenciesResponse
	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Markets retrieves the market countries that we support
func (c client) Markets(locale string) (*MarketsResponse, error) {
	uri := "/culture/markets/" + locale
	r, err := c.do("GET", uri, []byte{})
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = r.Body.Close()
	}()

	if r.StatusCode != http.StatusOK {
		var errREsp ErrorResponse
		if err := json.NewDecoder(r.Body).Decode(&errREsp); err != nil {
			return nil, badResponseStatus("poll", err, r)
		}

		return nil, errors.New(errREsp.Message)
	}

	var resp MarketsResponse
	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// NearestCulture retrieves the most relevant culture information for a user, based on an IP address
func (c client) NearestCulture(ip string) (*NearestCultureResponse, error) {
	uri := "/culture/nearestculture?ipAddress=" + ip
	r, err := c.do("GET", uri, []byte{})
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = r.Body.Close()
	}()

	if r.StatusCode != http.StatusOK {
		var errREsp ErrorResponse
		if err := json.NewDecoder(r.Body).Decode(&errREsp); err != nil {
			return nil, badResponseStatus("nearestculture", err, r)
		}

		return nil, errors.New(errREsp.Message)
	}

	var resp NearestCultureResponse
	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c client) do(method, uri string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest(method, c.getURL(uri), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(AuthHeader, c.cfg.APIKey)

	httpClient := http.Client{Timeout: c.cfg.QueriesTimeout}
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c client) getURL(uri string) string {
	b := strings.Builder{}
	b.WriteString(strings.TrimRight(BaseURL, "/"))
	b.WriteString("/")
	b.WriteString(strings.TrimLeft(uri, "/"))

	return b.String()
}

func badResponseStatus(action string, err error, resp *http.Response) error {
	var fullRes map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&fullRes); err != nil {
		return err
	}

	return fmt.Errorf("action: %s; status: %d; error: %s; full response: %+v", action, resp.StatusCode, err, fullRes)
}
