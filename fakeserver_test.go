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
	formIDUpdatedForm   = "formToUpdate"

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
	r.HandleFunc("/forms", srv.listFormsHandler).Methods(http.MethodGet)
	r.HandleFunc("/forms/{id}", srv.retrieveFormHandler).Methods(http.MethodGet)
	r.HandleFunc("/forms/{id}", srv.updateFormHandler).Methods(http.MethodPut)
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
	fmt.Fprint(w, `{"id": "`+formIDCreatedForm+`"}`)
}

func (s *typeformServer) listFormsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `{
  "total_items": 2,
  "page_count": 1,
  "items": [
    {
      "id": "abc123",
      "title": "My first typeform!",
      "last_updated_at": "2017-07-24T13:10:54.000Z",
      "self": {
        "href": "https://api.typeform.com/forms/abc123"
      },
      "theme": {
        "href": "https://api.typeform.com/themes/ghi789"
      },
      "_links": {
        "display": "https://subdomain.typeform.com/to/abc123"
      }
    },
    {
      "id": "def456",
      "title": "My second typeform",
      "last_updated_at": "2017-07-25T09:56:31.000Z",
      "self": {
        "href": "https://api.typeform.com/forms/def456"
      },
      "theme": {
        "href": "https://api.typeform.com/themes/ghi789"
      },
      "_links": {
        "display": "https://subdomain.typeform.com/to/def456"
      }
    }
  ]
}`)
}

func (s *typeformServer) retrieveFormHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["id"] != formIDRetrievedForm {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, responsePayloadFormNotFound, vars["id"])
		return
	}
	fmt.Fprint(w, `{"id": "`+formIDRetrievedForm+`"}`)
}

func (s *typeformServer) updateFormHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["id"] != formIDUpdatedForm {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, responsePayloadFormNotFound, vars["id"])
		return
	}
	fmt.Fprint(w, `{"id": "`+formIDUpdatedForm+`"}`)
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
