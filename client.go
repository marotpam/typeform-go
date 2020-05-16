package typeform

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	defaultTypeformAPIAddress = "https://api.typeform.com"
	defaultRetries            = 5
	defaultTimeoutSeconds     = 30
	defaultUserAgent          = "typeform-go"
)

// client will be used to interact with the typeform API.
type Client struct {
	// AccessToken to access the typeform API.
	AccessToken string
	// HttpClient that will be used to do http calls to the API.
	HttpClient http.Client
	// BaseURL is the base URL to be used in all requests to the API.
	BaseURL url.URL
	// MaxRetries is the maximum number of retries that will be done to the API on transient errors.
	MaxRetries int
	// BackoffDuration is the duration off the backoff period between retries.
	BackoffDuration BackoffDuration
}

type Config struct {
	// BaseAddress is the base address of the API, which can be configured to point to fake implementations of the
	// API for testing purposes. Defaults to https://api.typeform.com
	BaseAddress string
	// HttpTimeoutSeconds is the timeout to be used in http calls to the API, 30 seconds by default.
	HttpTimeoutSeconds int
	// MaxRetries is the maximum number of retries that will be done to the API on transient errors, 5 by default.
	MaxRetries int
	// BackoffDuration is the duration off the backoff period between retries, defaults to TwoSecondsIncrementBackoff.
	BackoffDuration BackoffDuration
	// HttpClient can be used to override the HTTP client that will calls to the typeform API, defaults
	// to http.DefaultClient with a Timeout of 30seconds. Setting it will override HttpTimeoutSeconds.
	HttpClient *http.Client
}

// NewDefaultClient will return a production client will all its config values set to defaults.
func NewDefaultClient(accessToken string) Client {
	u, _ := url.Parse(defaultTypeformAPIAddress)
	return Client{
		AccessToken:     accessToken,
		HttpClient:      defaultHTTPClient(),
		BaseURL:         *u,
		MaxRetries:      defaultRetries,
		BackoffDuration: TwoSecondsIncrementBackoff,
	}
}

// NewClient will return a default client with values overridden by the provided config.
func NewClient(token string, c Config) (Client, error) {
	client := NewDefaultClient(token)
	address := defaultTypeformAPIAddress

	if c.BaseAddress != "" {
		u, err := url.Parse(c.BaseAddress)
		if err != nil || u.Scheme == "" || u.Host == "" {
			return Client{}, fmt.Errorf("invalid base URL: %s", c.BaseAddress)
		}

		address = c.BaseAddress
	}

	u, _ := url.Parse(address)
	client.BaseURL = *u

	timeoutInSeconds := c.HttpTimeoutSeconds
	if timeoutInSeconds != 0 {
		if timeoutInSeconds < 0 {
			return Client{}, fmt.Errorf("http timeout cannot be lower than 0")
		}
		client.HttpClient.Timeout = time.Duration(timeoutInSeconds) * time.Second
	}

	maxRetries := c.MaxRetries
	if maxRetries > 0 {
		if timeoutInSeconds < 0 {
			return Client{}, fmt.Errorf("max retries cannot be lower than 0")
		}
		client.MaxRetries = maxRetries
	}

	if c.BackoffDuration != nil {
		client.BackoffDuration = c.BackoffDuration
	}

	if c.HttpClient != nil {
		client.HttpClient = *c.HttpClient
	}

	return client, nil
}

func defaultHTTPClient() http.Client {
	return http.Client{
		Timeout: defaultTimeoutSeconds * time.Second,
	}
}

// Do will do a request to the typeform API, and unmarshal its response into v. An error will be returned upon failure.
func (c Client) Do(req *http.Request, v interface{}) error {
	res, err := c.doWithRetries(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusNoContent || v == nil {
		return nil
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &v)
	if err != nil {
		return err
	}

	return nil
}

// BackoffDuration will determine for how long the client should wait before retrying a request after a given `try`.
type BackoffDuration func(try int) time.Duration

// NoBackoff is the no-op implementation of BackoffDuration. This means that there will be no pause between retries,
// which is not recommended for production usage.
func NoBackoff(try int) time.Duration {
	return 0
}

// TwoSecondsIncrementBackoff will retry after increases of 2 seconds between retries (2 seconds after first try,
// 4 seconds after the second, and so forth)
func TwoSecondsIncrementBackoff(try int) time.Duration {
	return time.Duration((try+1)*2) * time.Second
}

// doWithRetries will attempt to do a request to the typeform API, and retry for up to c.MaxRetries
// retries when applicable. If the request is successfully processed by typeform before running out of retries, the corresponding http.Response and
// no error will be returned. In case there's a permanent error or the client runs out of retries, an error detailing what went
// wrong will be returned.
// There will be an increasing pause of 2sec between retries by default, but this behaviour can be overridden by
// c.BackoffDuration.
func (c Client) doWithRetries(req *http.Request) (*http.Response, error) {
	c.applyDefaults(req)

	var res *http.Response
	var err error

	for try := 0; try < c.MaxRetries; try++ {
		res, err = c.HttpClient.Do(req)
		if err != nil {
			return nil, err
		}

		if hasBeenProcessed(res.StatusCode) {
			return res, nil
		}

		if isClientError(res.StatusCode) {
			break
		}

		time.Sleep(c.BackoffDuration(try))
	}

	return nil, newAPIError(req, res)
}

func (c Client) applyDefaults(req *http.Request) {
	baseURL := c.BaseURL
	req.URL.Scheme = baseURL.Scheme
	req.URL.Host = baseURL.Host
	req.URL.Path = baseURL.Path + req.URL.Path
	req.Header.Add("User-Agent", defaultUserAgent)
	req.Header.Set("Authorization", "Bearer "+c.AccessToken)
}

// hasBeenProcessed will return true when a request has been processed by typeform as intended.
func hasBeenProcessed(code int) bool {
	return code >= http.StatusOK && code < http.StatusMultipleChoices
}

// hasBeenProcessed will return true when a request has failed due to a permanent error.
func isClientError(code int) bool {
	return code >= http.StatusBadRequest && code < http.StatusInternalServerError
}
