package typeform

type ImageList []ImageListItem

type ImageListItem struct {
	ID       string `json:"id"`
	Src      string `json:"src"`
	FileName string `json:"file_name"`
}
