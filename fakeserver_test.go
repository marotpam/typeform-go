package typeform_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	"github.com/marotpam/typeform-go"
)

type typeformServer struct {
	forms      map[string]*typeform.Form
	httpServer *httptest.Server
}

func newFakeTypeformServer() *typeformServer {
	srv := &typeformServer{}

	r := mux.NewRouter()
	r.Use(accessTokenMiddleware)

	r.HandleFunc("/forms", srv.createFormHandler).Methods(http.MethodPost)

	srv.httpServer = httptest.NewServer(r)

	return srv
}

func accessTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("User-Agent") == "" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, `{"code":"AUTHENTICATION_FAILED","description":"Authentication credentials not found on the Request Headers"}`)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (s *typeformServer) createFormHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, `{"id": "fakeFormID"}`)
}

func (s *typeformServer) close() {
	s.httpServer.Close()
}

func newFakeServerClient(t *testing.T) typeform.Client {
	c, err := typeform.NewClient("fakeAccessToken", typeform.Config{
		BaseAddress: fakeServer.httpServer.URL,
	})
	assert.Nil(t, err)

	return c
}
