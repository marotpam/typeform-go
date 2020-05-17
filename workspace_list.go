package typeform

type WorkspaceList struct {
	Items      []*Workspace `json:"items,omitempty"`
	PageCount  int          `json:"page_count"`
	TotalItems int          `json:"total_items"`
}
