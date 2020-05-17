package typeform

import (
	"net/http"
	"strings"
)

type ImageService struct {
	resource
}

func NewImageService(client Client) *ImageService {
	return &ImageService{
		resource: resource{
			client: client,
		},
	}
}

func (s ImageService) Create(p CreateImageParams) (Image, error) {
	var created Image
	err := s.resource.create(resourceImages, p, &created)
	if err != nil {
		return Image{}, err
	}

	return created, nil
}

func (s ImageService) Retrieve(imageID string) (Image, error) {
	var created Image
	err := s.resource.retrieve(resourceImages, imageID, &created)
	if err != nil {
		return Image{}, err
	}

	return created, nil
}

func (s ImageService) RetrieveFormat(imageID, format, size string) (Image, error) {
	path := strings.NewReplacer(
		"{imageID}", imageID,
		"{format}", format,
		"{size}", size,
	).Replace("/images/{imageID}/{format}/{size}")
	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return Image{}, err
	}

	var i Image
	err = s.client.Do(req, &i)
	if err != nil {
		return Image{}, err
	}
	return i, nil
}

func (s ImageService) Delete(imageID string) error {
	return s.resource.delete(resourceImages, imageID)
}

func (s ImageService) List() (ImageList, error) {
	var l ImageList
	err := s.resource.list(resourceImages, nil, &l)
	if err != nil {
		return ImageList{}, err
	}

	return l, err
}
