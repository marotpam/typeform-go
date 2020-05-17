package typeform

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	CodeAuthenticationFailed = "AUTHENTICATION_FAILED"
	CodeFormNotFound         = "FORM_NOT_FOUND"
	CodeThemeNotFound        = "THEME_NOT_FOUND"
)

type Error struct {
	URL         string
	Method      string
	StatusCode  int
	Code        string `json:"code"`
	Description string `json:"description"`
}

func newAPIError(req *http.Request, res *http.Response) error {
	apiError := Error{
		URL:        req.URL.String(),
		Method:     req.Method,
		StatusCode: res.StatusCode,
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &apiError)
	if err != nil {
		return err
	}

	return apiError
}

func (e Error) Error() string {
	return fmt.Sprintf(
		"url: %s, method: %s, statusCode: %d, code: %s, description: %s",
		e.URL, e.Method, e.StatusCode, e.Code, e.Description,
	)
}
