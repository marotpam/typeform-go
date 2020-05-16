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

const (
	formIDCreatedForm   = "createdFormID"
	formIDRetrievedForm = "retrievedFormID"
	formIDDeletedForm   = "formToDelete"

	responsePayloadFormNotFound = `{"code":"FORM_NOT_FOUND","description":"Non existing form with uid %s"}`
	responsePayloadUnauthorized = `{"code":"AUTHENTICATION_FAILED","description":"Authentication credentials not found on the Request Headers"}`
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
	r.HandleFunc("/forms/{id}", srv.retrieveFormHandler).Methods(http.MethodGet)
	r.HandleFunc("/forms/{id}", srv.deleteFormHandler).Methods(http.MethodDelete)

	srv.httpServer = httptest.NewServer(r)

	return srv
}

func accessTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("User-Agent") == "" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, responsePayloadUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (s *typeformServer) createFormHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, `{"id": "`+formIDCreatedForm+`"}`)
}

func (s *typeformServer) retrieveFormHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	if vars["id"] != formIDRetrievedForm {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, responsePayloadFormNotFound, vars["id"])
		return
	}
	fmt.Fprint(w, `{"id": "`+formIDRetrievedForm+`"}`)
}

func (s *typeformServer) deleteFormHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["id"] != formIDDeletedForm {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, responsePayloadFormNotFound, vars["id"])
		return
	}
	w.WriteHeader(http.StatusNoContent)
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
