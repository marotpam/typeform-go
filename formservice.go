package typeform

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type FormService struct {
	client Client
}

func NewFormService(client Client) *FormService {
	return &FormService{client: client}
}

func (s FormService) Create(f Form) (Form, error) {
	b, err := json.Marshal(f)
	if err != nil {
		return Form{}, err
	}

	req, _ := http.NewRequest(http.MethodPost, "/forms", bytes.NewBuffer(b))

	var created Form
	err = s.client.Do(req, &created)
	if err != nil {
		return Form{}, err
	}

	return created, nil
}

func (s FormService) Retrieve(formID string) (Form, error) {
	req, _ := http.NewRequest(http.MethodGet, "/forms/"+formID, nil)

	var created Form
	err := s.client.Do(req, &created)
	if err != nil {
		return Form{}, err
	}

	return created, nil
}

func (s FormService) Delete(formID string) error {
	req, _ := http.NewRequest(http.MethodDelete, "/forms/"+formID, nil)

	return s.client.Do(req, nil)
}
