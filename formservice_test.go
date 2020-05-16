package typeform_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/marotpam/typeform-go"
)

func TestCreateForm(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		c := newFakeServerClient(t)

		svc := typeform.NewFormService(c)

		f, err := svc.Create(typeform.Form{
			Title: "my new form",
		})
		assert.Nil(t, err)

		assert.Equal(t, formIDCreatedForm, f.ID)
	})
}

func TestDeleteForm(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		c := newFakeServerClient(t)

		svc := typeform.NewFormService(c)

		assert.Nil(t, svc.Delete(formIDDeletedForm))
	})
}
