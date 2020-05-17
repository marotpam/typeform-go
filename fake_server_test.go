package typeform_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	"github.com/marotpam/typeform-go"
)

const (
	formIDCreatedForm   = "createdFormID"
	formIDRetrievedForm = "retrievedFormID"
	formIDDeletedForm   = "deletedFormID"
	formIDUpdatedForm   = "updatedFormID"

	imageIDCreatedImage   = "createdImageID"
	imageIDRetrievedImage = "retrievedImageID"
	imageIDDeletedImage   = "deletedImageID"

	themeIDCreatedTheme   = "createdThemeID"
	themeIDRetrievedTheme = "retrievedThemeID"
	themeIDDeletedTheme   = "deletedThemeID"
	themeIDUpdatedTheme   = "updatedThemeID"

	workspaceIDCreatedWorkspace   = "createdWorkspaceID"
	workspaceIDRetrievedWorkspace = "retrievedWorkspaceID"
	workspaceIDDeletedWorkspace   = "deletedWorkspaceID"
	workspaceIDUpdatedWorkspace   = "updatedWorkspaceID"

	responsePayloadNotFound     = `{"code":"%s_NOT_FOUND","description":"Non existing %s with uid %s"}`
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

	// 'create API' resources
	r.HandleFunc("/{collection:forms|images|themes|workspaces}", srv.createHandler).Methods(http.MethodPost)
	r.HandleFunc("/{collection:forms|images|themes|workspaces}", srv.listHandler).Methods(http.MethodGet)
	r.HandleFunc("/{collection:forms|images|themes|workspaces}/{id}", srv.retrieveHandler).Methods(http.MethodGet)
	r.HandleFunc("/{collection:forms|themes}/{id}", srv.updateHandler).Methods(http.MethodPut)
	r.HandleFunc("/{collection:forms|images|themes|workspaces}/{id}", srv.deleteHandler).Methods(http.MethodDelete)

	// image format endpoint
	r.HandleFunc(
		"/{collection:images}/{id}/{format:image|background|choice}/{size:default|mobile|thumbnail}",
		srv.retrieveFormattedImageHandler,
	).Methods(http.MethodGet)

	// 'results API' endpoints
	r.HandleFunc("/forms/{id}/responses", srv.retrieveResponsesHandler).Methods(http.MethodGet)
	r.HandleFunc("/forms/{id}/responses", srv.deleteHandler).Methods(http.MethodDelete)

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

func (s *typeformServer) createHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, getResourceFixture(r, "create"))
}

func (s *typeformServer) listHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, fmt.Sprintf("_fixtures/%s_list.json", getResource(r)))
}

func (s *typeformServer) retrieveHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if strings.HasPrefix(vars["id"], "unknown") {
		w.WriteHeader(http.StatusNotFound)
		resource := getResource(r)
		fmt.Fprintf(w, responsePayloadNotFound, strings.ToUpper(resource), resource, vars["id"])
		return
	}
	fmt.Fprint(w, getResourceFixture(r, "retrieve"))
}

func (s *typeformServer) retrieveFormattedImageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, getResourceFixture(r, "retrieve"))
}

func (s *typeformServer) updateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if strings.HasPrefix(vars["id"], "unknown") {
		w.WriteHeader(http.StatusNotFound)
		resource := getResource(r)
		fmt.Fprintf(w, responsePayloadNotFound, strings.ToUpper(resource), resource, vars["id"])
		return
	}
	fmt.Fprint(w, getResourceFixture(r, "update"))
}

func (s *typeformServer) deleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if strings.HasPrefix(vars["id"], "unknown") {
		w.WriteHeader(http.StatusNotFound)
		resource := getResource(r)
		fmt.Fprintf(w, responsePayloadNotFound, strings.ToUpper(resource), resource, vars["id"])
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (s *typeformServer) retrieveResponsesHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "_fixtures/response_list.json")
}

func (s *typeformServer) close() {
	s.httpServer.Close()
}

// getResource will return the resource of a given request path, for example, `form` is the resource
// for paths `/forms` and `/forms/{id}`
func getResource(r *http.Request) string {
	vars := mux.Vars(r)
	return strings.TrimSuffix(vars["collection"], "s")
}

func getResourceFixture(r *http.Request, action string) string {
	resource := getResource(r)
	b, _ := ioutil.ReadFile(fmt.Sprintf("_fixtures/%s.json", resource))

	id := fmt.Sprintf("%sd%sID", action, strings.Title(resource))

	return strings.Replace(string(b), "{{id}}", id, 1)
}

func newFakeServerClient(t *testing.T) typeform.Client {
	c, err := typeform.NewClient("fakeAccessToken", typeform.Config{
		BaseAddress: fakeServer.httpServer.URL,
	})
	assert.Nil(t, err)

	return c
}
