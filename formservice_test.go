package typeform_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/marotpam/typeform-go"
)

func TestCreateForm(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		svc := typeform.NewFormService(newFakeServerClient(t))

		f, err := svc.Create(typeform.Form{
			Title: "my new form",
		})
		assert.Nil(t, err)

		assert.Equal(t, formIDCreatedForm, f.ID)
	})
}

func TestRetrieveForm(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		svc := typeform.NewFormService(newFakeServerClient(t))

		f, err := svc.Retrieve(formIDRetrievedForm)
		assert.Nil(t, err)

		assert.Equal(t, formIDRetrievedForm, f.ID)
	})
	t.Run("not found error", func(t *testing.T) {
		svc := typeform.NewFormService(newFakeServerClient(t))

		_, err := svc.Retrieve("unknownFormID")
		assert.NotNil(t, err)

		tfErr, ok := err.(typeform.Error)
		assert.True(t, ok)

		assert.Equal(t, typeform.CodeFormNotFound, tfErr.Code)
	})
}

func TestUpdateForm(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		svc := typeform.NewFormService(newFakeServerClient(t))

		f, err := svc.Update(typeform.Form{
			ID:    formIDUpdatedForm,
			Title: "updated form",
		})
		assert.Nil(t, err)

		assert.Equal(t, formIDUpdatedForm, f.ID)
	})
}

func TestDeleteForm(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		svc := typeform.NewFormService(newFakeServerClient(t))

		assert.Nil(t, svc.Delete(formIDDeletedForm))
	})
}
