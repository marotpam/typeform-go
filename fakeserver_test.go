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

	themeIDCreatedTheme   = "createdThemeID"
	themeIDRetrievedTheme = "retrievedThemeID"
	themeIDDeletedTheme   = "themeToDelete"
	themeIDUpdatedTheme   = "themeToUpdate"

	responsePayloadFormNotFound  = `{"code":"FORM_NOT_FOUND","description":"Non existing form with uid %s"}`
	responsePayloadThemeNotFound = `{"code":"THEME_NOT_FOUND","description":"Non existing theme with uid %s"}`
	responsePayloadUnauthorized  = `{"code":"AUTHENTICATION_FAILED","description":"Authentication credentials not found on the Request Headers"}`
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

	r.HandleFunc("/themes", srv.createThemeHandler).Methods(http.MethodPost)
	r.HandleFunc("/themes", srv.listThemesHandler).Methods(http.MethodGet)
	r.HandleFunc("/themes/{id}", srv.retrieveThemeHandler).Methods(http.MethodGet)
	r.HandleFunc("/themes/{id}", srv.updateThemeHandler).Methods(http.MethodPut)
	r.HandleFunc("/themes/{id}", srv.deleteThemeHandler).Methods(http.MethodDelete)

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
	http.ServeFile(w, r, "_fixtures/form_create.json")
}

func (s *typeformServer) listFormsHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "_fixtures/form_list.json")
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

func (s *typeformServer) createThemeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "_fixtures/theme_create.json")
}

func (s *typeformServer) listThemesHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "_fixtures/theme_list.json")
}

func (s *typeformServer) retrieveThemeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["id"] != themeIDRetrievedTheme {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, responsePayloadThemeNotFound, vars["id"])
		return
	}
	fmt.Fprint(w, `{"id": "`+themeIDRetrievedTheme+`"}`)
}

func (s *typeformServer) updateThemeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["id"] != themeIDUpdatedTheme {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, responsePayloadThemeNotFound, vars["id"])
		return
	}
	fmt.Fprint(w, `{"id": "`+themeIDUpdatedTheme+`"}`)
}

func (s *typeformServer) deleteThemeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["id"] != themeIDDeletedTheme {
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
