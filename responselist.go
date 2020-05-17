package typeform

import "time"

type ResponseListParams struct {
	PageSize            int        `url:"page_size,omitempty"`
	Since               *time.Time `url:"since,omitempty"`
	Until               *time.Time `url:"until,omitempty"`
	After               *time.Time `url:"after,omitempty"`
	Before              *time.Time `url:"before,omitempty"`
	IncludedResponseIDs string     `url:"included_response_ids,omitempty"`
	Completed           *bool      `url:"completed,omitempty"`
	Sort                string     `url:"sort,omitempty"`
	Query               string     `url:"query,omitempty"`
	Fields              string     `url:"fields,omitempty"`
}

type ResponseList struct {
	TotalItems int        `json:"total_items"`
	PageCount  int        `json:"page_count"`
	Items      []Response `json:"items"`
}
