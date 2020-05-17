package typeform

type WorkspaceService struct {
	resource
}

func NewWorkspaceService(client Client) *WorkspaceService {
	return &WorkspaceService{
		resource: resource{
			client: client,
		},
	}
}

func (s WorkspaceService) Create(w Workspace) (Workspace, error) {
	var created Workspace
	err := s.resource.create(resourceWorkspaces, w, &created)
	if err != nil {
		return Workspace{}, err
	}

	return created, nil
}

func (s WorkspaceService) Retrieve(workspaceID string) (Workspace, error) {
	var created Workspace
	err := s.resource.retrieve(resourceWorkspaces, workspaceID, &created)
	if err != nil {
		return Workspace{}, err
	}

	return created, nil
}

func (s WorkspaceService) Delete(workspaceID string) error {
	return s.resource.delete(resourceWorkspaces, workspaceID)
}

func (s WorkspaceService) List() (WorkspaceList, error) {
	var l WorkspaceList
	err := s.resource.list(resourceWorkspaces, nil, &l)
	if err != nil {
		return WorkspaceList{}, err
	}

	return l, err
}
