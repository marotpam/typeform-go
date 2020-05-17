package typeform_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/marotpam/typeform-go"
)

func TestCreateTheme(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		svc := typeform.NewThemeService(newFakeServerClient(t))

		f, err := svc.Create(typeform.Theme{
			Name: "my new theme",
		})
		assert.Nil(t, err)

		assert.Equal(t, themeIDCreatedTheme, f.ID)
	})
}

func TestRetrieveTheme(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		svc := typeform.NewThemeService(newFakeServerClient(t))

		f, err := svc.Retrieve(themeIDRetrievedTheme)
		assert.Nil(t, err)

		assert.Equal(t, themeIDRetrievedTheme, f.ID)
	})
	t.Run("not found error", func(t *testing.T) {
		svc := typeform.NewThemeService(newFakeServerClient(t))

		_, err := svc.Retrieve("unknownThemeID")
		assert.NotNil(t, err)

		tfErr, ok := err.(typeform.Error)
		assert.True(t, ok)

		assert.Equal(t, typeform.CodeThemeNotFound, tfErr.Code)
	})
}

func TestUpdateTheme(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		svc := typeform.NewThemeService(newFakeServerClient(t))

		f, err := svc.Update(typeform.Theme{
			ID:   themeIDUpdatedTheme,
			Name: "updated theme",
		})
		assert.Nil(t, err)

		assert.Equal(t, themeIDUpdatedTheme, f.ID)
	})
}

func TestListThemes(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		svc := typeform.NewThemeService(newFakeServerClient(t))

		list, err := svc.List(typeform.ThemeListParams{})
		assert.Nil(t, err)

		assert.Len(t, list.Items, 1)
	})
}

func TestDeleteTheme(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		svc := typeform.NewThemeService(newFakeServerClient(t))

		assert.Nil(t, svc.Delete(themeIDDeletedTheme))
	})
}
