package typeform

import "time"

type Response struct {
	LandingID   string    `json:"landing_id"`
	Token       string    `json:"token"`
	ResponseID  string    `json:"response_id,omitempty"`
	LandedAt    time.Time `json:"landed_at"`
	SubmittedAt time.Time `json:"submitted_at"`
	Metadata    struct {
		UserAgent string `json:"user_agent"`
		Platform  string `json:"platform"`
		Referer   string `json:"referer"`
		NetworkID string `json:"network_id"`
		Browser   string `json:"browser"`
	} `json:"metadata"`
	Answers []struct {
		Field struct {
			ID   string `json:"id"`
			Type string `json:"type"`
			Ref  string `json:"ref"`
		} `json:"field"`
		Type    string `json:"type"`
		Text    string `json:"text,omitempty"`
		Boolean bool   `json:"boolean,omitempty"`
		Email   string `json:"email,omitempty"`
		Number  int    `json:"number,omitempty"`
		Choices struct {
			Labels []string `json:"labels"`
		} `json:"choices,omitempty"`
		Date   time.Time `json:"date,omitempty"`
		Choice struct {
			Label string `json:"label"`
		} `json:"choice,omitempty"`
	} `json:"answers"`
	Hidden     map[string]string      `json:"hidden,omitempty"`
	Calculated map[string]interface{} `json:"calculated"`
}
