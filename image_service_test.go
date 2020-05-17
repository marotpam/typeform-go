package typeform_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/marotpam/typeform-go"
)

func TestCreateImage(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		svc := typeform.NewImageService(newFakeServerClient(t))

		f, err := svc.Create(typeform.CreateImageParams{})
		assert.Nil(t, err)

		assert.Equal(t, imageIDCreatedImage, f.ID)
	})
}

func TestRetrieveImage(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		svc := typeform.NewImageService(newFakeServerClient(t))

		f, err := svc.Retrieve(imageIDRetrievedImage)
		assert.Nil(t, err)

		assert.Equal(t, imageIDRetrievedImage, f.ID)
	})
	t.Run("not found error", func(t *testing.T) {
		svc := typeform.NewImageService(newFakeServerClient(t))

		_, err := svc.Retrieve("unknownImageID")
		assert.NotNil(t, err)

		tfErr, ok := err.(typeform.Error)
		assert.True(t, ok)

		assert.Equal(t, typeform.CodeImageNotFound, tfErr.Code)
	})
}

func TestRetrieveImageFormat(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		svc := typeform.NewImageService(newFakeServerClient(t))

		f, err := svc.RetrieveFormat(imageIDRetrievedImage, typeform.ImageFormatBackground, typeform.ImageSizeDefault)
		assert.Nil(t, err)

		assert.Equal(t, imageIDRetrievedImage, f.ID)
	})
	t.Run("not found error", func(t *testing.T) {
		svc := typeform.NewImageService(newFakeServerClient(t))

		_, err := svc.Retrieve("unknownImageID")
		assert.NotNil(t, err)

		tfErr, ok := err.(typeform.Error)
		assert.True(t, ok)

		assert.Equal(t, typeform.CodeImageNotFound, tfErr.Code)
	})
}

func TestListImages(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		svc := typeform.NewImageService(newFakeServerClient(t))

		list, err := svc.List()
		assert.Nil(t, err)

		assert.Len(t, list, 1)
	})
}

func TestDeleteImage(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		svc := typeform.NewImageService(newFakeServerClient(t))

		assert.Nil(t, svc.Delete(imageIDDeletedImage))
	})
}
