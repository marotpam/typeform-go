package typeform_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/marotpam/typeform-go"
)

func TestDefaultClient(t *testing.T) {
	c := typeform.NewDefaultClient("")

	assert.Equal(t, "https://api.typeform.com", c.BaseURL.String(), "incorrect base URL")
	assert.Equal(t, 30*time.Second, c.HttpClient.Timeout, "incorrect http client timeout")
}

func TestClientWithCustomConfig(t *testing.T) {
	t.Run("with valid config", func(t *testing.T) {
		accessToken := "943af478d3ff3d4d760020c11af102b79c440513"
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, r.Header.Get("Authorization"), "Bearer "+accessToken, "accessToken should be sent in all requests")
		}))
		defer testServer.Close()

		c, err := typeform.NewClient(accessToken,
			typeform.Config{
				BaseAddress:        testServer.URL,
				HttpTimeoutSeconds: 3,
			},
		)

		assert.Nil(t, err)
		assert.Equal(t, testServer.URL, c.BaseURL.String(), "incorrect base URL")
		assert.Equal(t, 3*time.Second, c.HttpClient.Timeout, "incorrect timeout")

		req, _ := http.NewRequest(http.MethodGet, "", nil)
		assert.Nil(t, c.Do(req, nil))
	})
	t.Run("with custom httpclient", func(t *testing.T) {
		var requestInBackend *http.Request
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestInBackend = r
		}))
		defer srv.Close()

		requestSpy := &requestSpy{}
		c, err := typeform.NewClient("", typeform.Config{
			BaseAddress: srv.URL,
			HttpClient: &http.Client{
				Transport: requestSpy,
			},
		})

		assert.Nil(t, err)
		request, _ := http.NewRequest(http.MethodGet, "/path", nil)

		err = c.Do(request, nil)
		assert.Nil(t, err)

		assert.Equal(t, "spy", requestInBackend.Header.Get("X-User-Agent"))
	})
	t.Run("with invalid config", func(t *testing.T) {
		_, err := typeform.NewClient("", typeform.Config{
			BaseAddress:        "invalid/without.scheme/scheme",
			HttpTimeoutSeconds: -100,
		})

		assert.NotNil(t, err)
	})
}

func TestRetryLogic(t *testing.T) {
	t.Run("success on first try", func(t *testing.T) {
		calls := 0
		failedCalls := 0
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			calls++
			if calls > failedCalls {
				_, _ = fmt.Fprint(w, `{"code": "success after retries"}`)
				return
			}
			w.WriteHeader(http.StatusBadGateway)
		}))
		defer srv.Close()

		client, _ := typeform.NewClient("", typeform.Config{
			BaseAddress:     srv.URL,
			BackoffDuration: typeform.NoBackoff,
		})

		request, _ := http.NewRequest(http.MethodPost, "", nil)
		err := client.Do(request, nil)
		assert.Nil(t, err, "should not get an error if it succeeds on first try")
		assert.Equal(t, 1, calls)
	})
	t.Run("success after first retry", func(t *testing.T) {
		calls := 0
		failedCalls := 1
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			calls++
			if calls > failedCalls {
				_, _ = fmt.Fprint(w, `{"code": "success after retries"}`)
				return
			}
			w.WriteHeader(http.StatusBadGateway)
		}))
		defer srv.Close()

		client, _ := typeform.NewClient("", typeform.Config{
			BaseAddress:     srv.URL,
			BackoffDuration: typeform.NoBackoff,
		})

		request, _ := http.NewRequest(http.MethodPost, "", nil)
		err := client.Do(request, nil)
		assert.Nil(t, err, "should not get an error if it succeeds after first retry")
		assert.Equal(t, 2, calls)
	})
	t.Run("error after permanent error on first try", func(t *testing.T) {
		calls := 0
		failedCalls := 5
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			calls++
			if calls > failedCalls {
				_, _ = fmt.Fprint(w, `{"code": "success after retries"}`)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
		}))
		defer srv.Close()

		client, _ := typeform.NewClient("", typeform.Config{
			BaseAddress:     srv.URL,
			BackoffDuration: typeform.NoBackoff,
		})

		request, _ := http.NewRequest(http.MethodPost, "", nil)
		err := client.Do(request, nil)
		assert.NotNil(t, err, "should get a permanent error after first try")
		assert.Equal(t, 1, calls)
	})
	t.Run("error after running out of retries", func(t *testing.T) {
		calls := 0
		failedCalls := 2
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			calls++
			if calls > failedCalls {
				_, _ = fmt.Fprint(w, `{"code": "success after retries"}`)
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer srv.Close()

		client, _ := typeform.NewClient("", typeform.Config{
			BaseAddress:     srv.URL,
			MaxRetries:      2,
			BackoffDuration: typeform.NoBackoff,
		})

		request, _ := http.NewRequest(http.MethodPost, "", nil)
		err := client.Do(request, nil)
		assert.NotNil(t, err, "should get an error after runnning out of retries")
		assert.Equal(t, 2, calls)
	})
}

// requestSpy is a struct that will implement the http.Roundtripper interface, and will be used to assert that custom
// http clients can be used in the config.
type requestSpy struct {
	lastRequest *http.Request
}

// RoundTrip will keep track of requests done and flag them with a `X-User-Agent` header.
func (c *requestSpy) RoundTrip(r *http.Request) (*http.Response, error) {
	c.lastRequest = r
	r.Header.Add("X-User-Agent", "spy")
	return http.DefaultTransport.RoundTrip(r)
}
