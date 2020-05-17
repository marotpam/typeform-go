package typeform

type Theme struct {
	ID                   string      `json:"id,omitempty"`
	Name                 string      `json:"name,omitempty"`
	Visibility           string      `json:"visibility,omitempty"`
	Background           *Background `json:"background,omitempty"`
	Colors               *Colors     `json:"colors,omitempty"`
	Font                 string      `json:"font,omitempty"`
	HasTransparentButton bool        `json:"has_transparent_button,omitempty"`
}
type Background struct {
	Brightness float64 `json:"brightness,omitempty"`
	ImageID    int     `json:"image_id,omitempty"`
	Layout     string  `json:"layout,omitempty"`
}
type Colors struct {
	Answer     string `json:"answer,omitempty"`
	Background string `json:"background,omitempty"`
	Button     string `json:"button,omitempty"`
	Question   string `json:"question,omitempty"`
}
