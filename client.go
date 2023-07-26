package skyscanner

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
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
func (c client) Create(ctx context.Context, req *CreateRequest) (*CreatePollResponse, *ErrorResponse) {
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, internalErrorResponse("request marshalling error: " + err.Error())
	}

	r, err := c.do(ctx, http.MethodPost, "/flights/live/search/create", jsonData)
	if err != nil {
		return nil, internalErrorResponse("request doing error: " + err.Error())
	}
	defer func() {
		_ = r.Body.Close()
	}()

	if r.StatusCode != http.StatusOK {
		return nil, badResponseStatus(r)
	}

	var resp CreatePollResponse
	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
		return nil, internalErrorResponse("response decoding error: " + err.Error())
	}

	return &resp, nil
}

// Poll does a poll request
func (c client) Poll(ctx context.Context, req *PollRequest) (*CreatePollResponse, *ErrorResponse) {
	uri := "/flights/live/search/poll/" + req.SessionToken
	r, err := c.do(ctx, http.MethodPost, uri, []byte{})
	if err != nil {
		return nil, internalErrorResponse("request doing error: " + err.Error())
	}
	defer func() {
		_ = r.Body.Close()
	}()

	if r.StatusCode != http.StatusOK {
		return nil, badResponseStatus(r)
	}

	var resp CreatePollResponse
	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
		return nil, internalErrorResponse("response decoding error: " + err.Error())
	}

	return &resp, nil
}

// Locales retrieves the locales that we support to translate your content
func (c client) Locales(ctx context.Context) (*LocalesResponse, *ErrorResponse) {
	r, err := c.do(ctx, http.MethodGet, "/culture/locales", []byte{})
	if err != nil {
		return nil, internalErrorResponse("request doing error: " + err.Error())
	}
	defer func() {
		_ = r.Body.Close()
	}()

	if r.StatusCode != http.StatusOK {
		return nil, badResponseStatus(r)
	}

	var resp LocalesResponse
	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
		return nil, internalErrorResponse("response decoding error: " + err.Error())
	}

	return &resp, nil
}

// Currencies retrieves the currencies that Skyscanner support and information about format
func (c client) Currencies(ctx context.Context) (*CurrenciesResponse, *ErrorResponse) {
	r, err := c.do(ctx, http.MethodGet, "/culture/currencies", []byte{})
	if err != nil {
		return nil, internalErrorResponse("request doing error: " + err.Error())
	}
	defer func() {
		_ = r.Body.Close()
	}()

	if r.StatusCode != http.StatusOK {
		return nil, badResponseStatus(r)
	}

	var resp CurrenciesResponse
	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
		return nil, internalErrorResponse("response decoding error: " + err.Error())
	}

	return &resp, nil
}

// Markets retrieves the market countries that we support
func (c client) Markets(ctx context.Context, locale string) (*MarketsResponse, *ErrorResponse) {
	uri := "/culture/markets/" + locale
	r, err := c.do(ctx, http.MethodGet, uri, []byte{})
	if err != nil {
		return nil, internalErrorResponse("request doing error: " + err.Error())
	}
	defer func() {
		_ = r.Body.Close()
	}()

	if r.StatusCode != http.StatusOK {
		return nil, badResponseStatus(r)
	}

	var resp MarketsResponse
	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
		return nil, internalErrorResponse("response decoding error: " + err.Error())
	}

	return &resp, nil
}

// NearestCulture retrieves the most relevant culture information for a user, based on an IP address
func (c client) NearestCulture(ctx context.Context, ip string) (*NearestCultureResponse, *ErrorResponse) {
	uri := "/culture/nearestculture?ipAddress=" + ip
	r, err := c.do(ctx, http.MethodGet, uri, []byte{})
	if err != nil {
		return nil, internalErrorResponse("request doing error: " + err.Error())
	}
	defer func() {
		_ = r.Body.Close()
	}()

	if r.StatusCode != http.StatusOK {
		return nil, badResponseStatus(r)
	}

	var resp NearestCultureResponse
	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
		return nil, internalErrorResponse("response decoding error: " + err.Error())
	}

	return &resp, nil
}

// AutoSuggestFlights returns a list of places that match a specified searchTerm
func (c client) AutoSuggestFlights(ctx context.Context, req *AutoSuggestFlightsRequest) (*AutoSuggestFlightsResponse, *ErrorResponse) {
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, internalErrorResponse("request marshalling error: " + err.Error())
	}

	r, err := c.do(ctx, http.MethodPost, "/autosuggest/flights", jsonData)
	if err != nil {
		return nil, internalErrorResponse("request doing error: " + err.Error())
	}
	defer func() {
		_ = r.Body.Close()
	}()

	if r.StatusCode != http.StatusOK {
		return nil, badResponseStatus(r)
	}

	var resp AutoSuggestFlightsResponse
	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
		return nil, internalErrorResponse("response decoding error: " + err.Error())
	}

	return &resp, nil
}

func (c client) do(ctx context.Context, method, uri string, body []byte) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, c.getURL(uri), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(AuthHeader, c.cfg.APIKey)
	req.Close = true

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

func internalErrorResponse(msg string) *ErrorResponse {
	return &ErrorResponse{
		Code:    http.StatusInternalServerError,
		Message: msg,
	}
}

func badResponseStatus(resp *http.Response) *ErrorResponse {
	errResp := &ErrorResponse{}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		errResp.Code = resp.StatusCode
		errResp.Message = "response reading error:" + err.Error()
		return errResp
	}

	if err := json.Unmarshal(b, &errResp); err != nil {
		errResp.Code = resp.StatusCode
		errResp.Message = string(b)
		return errResp
	}

	return errResp
}
