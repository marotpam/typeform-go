package typeform

type Workspace struct {
	ID        string            `json:"id,omitempty"`
	Name      string            `json:"name"`
	AccountID string            `json:"account_id,omitempty"`
	Default   bool              `json:"default,omitempty"`
	Shared    bool              `json:"shared,omitempty"`
	Forms     *Forms            `json:"forms,omitempty"`
	Members   []WorkspaceMember `json:"members,omitempty"`
	Self      *Self             `json:"self,omitempty"`
}

type Forms struct {
	Count int    `json:"count,omitempty"`
	Href  string `json:"href"`
}

type WorkspaceMember struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Role  string `json:"role"`
}

type Self struct {
	Href string `json:"href"`
}
