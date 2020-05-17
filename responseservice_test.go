package typeform_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/marotpam/typeform-go"
)

func TestListResponses(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		svc := typeform.NewResponseService(newFakeServerClient(t))

		list, err := svc.List("formID", typeform.ResponseListParams{})
		assert.Nil(t, err)

		assert.Len(t, list.Items, 1)
	})
}

func TestDeleteResponse(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		svc := typeform.NewResponseService(newFakeServerClient(t))

		assert.Nil(t, svc.Delete("formID", []string{"responseID"}))
	})
}
