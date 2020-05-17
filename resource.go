package typeform

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

const (
	resourceForms = "forms"
)

type resource struct {
	client Client
}

func getPath(resourceName, resourceID string) string {
	return fmt.Sprintf("/%s/%s", resourceName, resourceID)
}

func (s resource) create(resourceName string, resource, v interface{}) error {
	b, err := json.Marshal(resource)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, "/"+resourceName, bytes.NewBuffer(b))
	if err != nil {
		return err
	}

	return s.client.Do(req, v)
}

func (s resource) retrieve(resourceName, resourceID string, v interface{}) error {
	req, err := http.NewRequest(http.MethodGet, getPath(resourceName, resourceID), nil)
	if err != nil {
		return err
	}

	return s.client.Do(req, v)
}

func (s resource) update(resourceName, resourceID string, resource, v interface{}) error {
	b, err := json.Marshal(resource)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, getPath(resourceName, resourceID), bytes.NewBuffer(b))
	if err != nil {
		return err
	}

	return s.client.Do(req, v)
}

func (s resource) delete(resourceName, resourceID string) error {
	req, err := http.NewRequest(http.MethodDelete, getPath(resourceName, resourceID), nil)
	if err != nil {
		return err
	}

	return s.client.Do(req, nil)
}

func (s resource) list(resourceName string, params, v interface{}) error {
	qv, err := query.Values(params)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodGet, "/"+resourceName, nil)
	if err != nil {
		return err
	}
	req.URL.RawQuery = qv.Encode()

	return s.client.Do(req, v)
}
