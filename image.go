package typeform

const (
	ImageFormatImage      = "image"
	ImageFormatChoice     = "choice"
	ImageFormatBackground = "background"

	ImageSizeDefault        = "default"
	ImageSizeMobile         = "mobile"
	ImageSizeSuperMobile    = "supermobile"
	ImageSizeSuperMobileFit = "supermobilefit"
	ImageSizeSuperSize      = "supersize"
	ImageSizeSuperSizeFit   = "supersizefit"
	ImageSizeSuperTablet    = "supertablet"
	ImageSizeTablet         = "tablet"
	ImageSizeThumbnail      = "thumbnail"
)

type Image struct {
	ID        string `json:"id,omitempty"`
	Src       string `json:"src,omitempty"`
	FileName  string `json:"file_name,omitempty"`
	Width     int    `json:"width,omitempty"`
	Height    int    `json:"height,omitempty"`
	MediaType string `json:"media_type,omitempty"`
	HasAlpha  bool   `json:"has_alpha,omitempty"`
	AvgColor  string `json:"avg_color,omitempty"`
}

type CreateImageParams struct {
	Image    string `json:"image,omitempty"`
	FileName string `json:"file_name"`
	Url      string `json:"url,omitempty"`
}
