package typeform

import "time"

type FormListParams struct {
	Search      string `url:"search,omitempty"`
	Page        int    `url:"page,omitempty"`
	PageSize    int    `url:"page_size,omitempty"`
	WorkspaceID string `url:"workspace_id,omitempty"`
}

type FormList struct {
	TotalItems int            `json:"total_items"`
	PageCount  int            `json:"page_count"`
	Items      []FormListItem `json:"items"`
}

type FormListItem struct {
	ID            string    `json:"id"`
	Title         string    `json:"title"`
	LastUpdatedAt time.Time `json:"last_updated_at"`
	Self          struct {
		Href string `json:"href"`
	} `json:"self"`
	Theme struct {
		Href string `json:"href"`
	} `json:"theme"`
	Links struct {
		Display string `json:"display"`
	} `json:"_links"`
}
