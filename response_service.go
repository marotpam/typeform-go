package typeform

import (
	"net/http"
	"strings"

	"github.com/google/go-querystring/query"
)

type ResponseService struct {
	client Client
}

func NewResponseService(client Client) *ResponseService {
	return &ResponseService{client: client}
}

func (s ResponseService) List(formID string, p ResponseListParams) (ResponseList, error) {
	v, err := query.Values(p)
	if err != nil {
		return ResponseList{}, err
	}

	req, _ := http.NewRequest(http.MethodGet, "/forms/"+formID+"/responses", nil)
	req.URL.RawQuery = v.Encode()

	var l ResponseList
	err = s.client.Do(req, &l)
	if err != nil {
		return ResponseList{}, err
	}

	return l, err
}

func (s ResponseService) Delete(formID string, includedTokens []string) error {
	req, err := http.NewRequest(http.MethodDelete, "/forms/"+formID+"/responses", nil)
	if err != nil {
		return err
	}
	req.URL.RawQuery = "included_tokens=" + strings.Join(includedTokens, ",")

	return s.client.Do(req, nil)
}
