package typeform

type FormService struct {
	resource
}

func NewFormService(client Client) *FormService {
	return &FormService{
		resource: resource{
			client: client,
		},
	}
}

func (s FormService) Create(f Form) (Form, error) {
	var created Form
	err := s.resource.create(resourceForms, f, &created)
	if err != nil {
		return Form{}, err
	}

	return created, nil
}

func (s FormService) Retrieve(formID string) (Form, error) {
	var created Form
	err := s.resource.retrieve(resourceForms, formID, &created)
	if err != nil {
		return Form{}, err
	}

	return created, nil
}

func (s FormService) Update(f Form) (Form, error) {
	var updated Form
	err := s.resource.update(resourceForms, f.ID, f, &updated)
	if err != nil {
		return Form{}, err
	}

	return updated, nil
}

func (s FormService) Delete(formID string) error {
	return s.resource.delete(resourceForms, formID)
}

func (s FormService) List(p FormListParams) (FormList, error) {
	var l FormList
	err := s.resource.list(resourceForms, p, &l)
	if err != nil {
		return FormList{}, err
	}

	return l, err
}
