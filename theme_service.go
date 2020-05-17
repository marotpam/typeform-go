package typeform

type ThemeService struct {
	resource
}

func NewThemeService(client Client) *ThemeService {
	return &ThemeService{
		resource: resource{
			client: client,
		},
	}
}

func (s ThemeService) Create(t Theme) (Theme, error) {
	var created Theme
	err := s.resource.create(resourceThemes, t, &created)
	if err != nil {
		return Theme{}, err
	}

	return created, nil
}

func (s ThemeService) Retrieve(themeID string) (Theme, error) {
	var created Theme
	err := s.resource.retrieve(resourceThemes, themeID, &created)
	if err != nil {
		return Theme{}, err
	}

	return created, nil
}

func (s ThemeService) Update(t Theme) (Theme, error) {
	var updated Theme
	err := s.resource.update(resourceThemes, t.ID, t, &updated)
	if err != nil {
		return Theme{}, err
	}

	return updated, nil
}

func (s ThemeService) Delete(themeID string) error {
	return s.resource.delete(resourceThemes, themeID)
}

func (s ThemeService) List(p ThemeListParams) (ThemeList, error) {
	var l ThemeList
	err := s.resource.list(resourceThemes, p, &l)
	if err != nil {
		return ThemeList{}, err
	}

	return l, err
}
