package typeform

const (
	FieldTypeYesNo          = "yes_no"
	FieldTypeShortText      = "short_text"
	FieldTypeMultipleChoice = "multiple_choice"
	FieldTypePictureChoice  = "picture_choice"
	FieldTypeGroup          = "group"
	FieldTypeEmail          = "email"
	FieldTypeDropdown       = "dropdown"
	FieldTypeLongText       = "long_text"
	FieldTypeFileUpload     = "file_upload"
	FieldTypeNumber         = "number"
	FieldTypeWebsite        = "website"
	FieldTypeLegal          = "legal"
	FieldTypeDate           = "date"
	FieldTypeRating         = "rating"
	FieldTypeStatement      = "statement"
	FieldTypePayment        = "payment"
	FieldTypeOpinionScale   = "opinion_scale"
	FieldTypePhoneNumber    = "phone_number"
)

type Form struct {
	ID              string                 `json:"id,omitempty"`
	Title           string                 `json:"title"`
	Workspace       *Link                  `json:"workspace,omitempty"`
	Theme           *Link                  `json:"theme,omitempty"`
	Settings        *Settings              `json:"settings,omitempty"`
	ThankyouScreens []*ThankYouScreen      `json:"thankyou_screens,omitempty"`
	WelcomeScreens  []*WelcomeScreen       `json:"welcome_screens,omitempty"`
	Fields          []*Field               `json:"fields,omitempty"`
	Hidden          []string               `json:"hidden,omitempty"`
	Variables       map[string]interface{} `json:"variables,omitempty"`
	Logic           []*Logic               `json:"logic,omitempty"`
	Outcome         *Outcome               `json:"outcome,omitempty"`
	Links           *FormLinks             `json:"_links,omitempty"`
}

type Settings struct {
	Language               string         `json:"language"`
	ProgressBar            string         `json:"progress_bar"`
	Meta                   SettingsMeta   `json:"meta"`
	IsPublic               bool           `json:"is_public"`
	IsTrial                bool           `json:"is_trial"`
	ShowProgressBar        bool           `json:"show_progress_bar"`
	ShowTypeformBranding   bool           `json:"show_typeform_branding"`
	RedirectAfterSubmitURL string         `json:"redirect_after_submit_url,omitempty"`
	FacebookPixel          string         `json:"facebook_pixel,omitempty"`
	GoogleAnalytics        string         `json:"google_analytics,omitempty"`
	GoogleTagManager       string         `json:"google_tag_manager,omitempty"`
	Notifications          *Notifications `json:"Notifications,omitempty"`
}

type SettingsMeta struct {
	AllowIndexing bool   `json:"allow_indexing"`
	Description   string `json:"description,omitempty"`
	Image         *Link  `json:"image,omitempty"`
}

type Notifications struct {
	Self       *SelfNotification       `json:"self,omitempty"`
	Respondent *RespondentNotification `json:"respondent,omitempty"`
}

type SelfNotification struct {
	Enabled    bool     `json:"enabled"`
	Message    string   `json:"message"`
	Recipients []string `json:"recipients"`
	ReplyTo    string   `json:"reply_to,omitempty"`
	Subject    string   `json:"subject"`
}

type RespondentNotification struct {
	Enabled   bool      `json:"enabled"`
	Message   string    `json:"message"`
	Recipient string    `json:"recipient"`
	ReplyTo   *[]string `json:"reply_to,omitempty"`
	Subject   string    `json:"subject"`
}

type ThankYouScreen struct {
	ID         string                   `json:"id,omitempty"`
	Ref        string                   `json:"ref,omitempty"`
	Title      string                   `json:"title"`
	Properties ThankYouScreenProperties `json:"properties"`
	Attachment *Attachment              `json:"attachment,omitempty"`
}

type ThankYouScreenProperties struct {
	ShowButton  *bool  `json:"show_button,omitempty"`
	ShareIcons  *bool  `json:"share_icons,omitempty"`
	ButtonMode  string `json:"button_mode,omitempty"`
	ButtonText  string `json:"button_text,omitempty"`
	RedirectURL string `json:"redirect_url,omitempty"`
	Description string `json:"description,omitempty"`
}

type WelcomeScreen struct {
	ID         string                  `json:"id,omitempty"`
	Ref        string                  `json:"ref,omitempty"`
	Title      string                  `json:"title"`
	Properties WelcomeScreenProperties `json:"properties"`
	Attachment *Attachment             `json:"attachment,omitempty"`
}

type WelcomeScreenProperties struct {
	ShowButton  *bool  `json:"show_button,omitempty"`
	ButtonText  string `json:"button_text,omitempty"`
	Description string `json:"description,omitempty"`
}

type Attachment struct {
	Type       string                `json:"type"`
	Href       string                `json:"href"`
	Scale      float64               `json:"scale,omitempty"`
	Properties *attachmentProperties `json:"FieldProperties,omitempty"`
}

type attachmentProperties struct {
	Brightness *float64    `json:"brightness,omitempty"`
	FocalPoint *FocalPoint `json:"focal_point,omitempty"`
}

type FocalPoint struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Layout struct {
	Type       string     `json:"type"`
	Attachment Attachment `json:"Attachment"`
}

type Link struct {
	Href string `json:"href,omitempty"`
}

type FormLinks struct {
	Display string `json:"display"`
}

type Field struct {
	ID          string            `json:"id,omitempty"`
	Ref         string            `json:"ref,omitempty"`
	Title       string            `json:"title"`
	Properties  *FieldProperties  `json:"properties,omitempty"`
	Validations *FieldValidations `json:"validations,omitempty"`
	Type        string            `json:"type"`
	Attachment  *Attachment       `json:"attachment,omitempty"`
	Layout      *Layout           `json:"Layout,omitempty"`
}

type FieldProperties struct {
	Description            string          `json:"description,omitempty"`
	ButtonText             string          `json:"button_text,omitempty"`
	Randomize              *bool           `json:"randomize,omitempty"`
	AllowMultipleSelection *bool           `json:"allow_multiple_selection,omitempty"`
	AllowOtherChoice       *bool           `json:"allow_other_choice,omitempty"`
	VerticalAlignment      *bool           `json:"vertical_alignment,omitempty"`
	Supersized             *bool           `json:"supersized,omitempty"`
	ShowLabels             *bool           `json:"show_labels,omitempty"`
	ShowButton             *bool           `json:"show_button,omitempty"`
	AlphabeticalOrder      *bool           `json:"alphabetical_order,omitempty"`
	HideMarks              *bool           `json:"hide_marks,omitempty"`
	StartAtOne             *bool           `json:"start_at_one,omitempty"`
	Choices                *[]*FieldChoice `json:"choices,omitempty"`
	Fields                 []*Field        `json:"fields,omitempty"`
	Separator              string          `json:"separator,omitempty"`
	Structure              string          `json:"structure,omitempty"`
	Shape                  string          `json:"shape,omitempty"`
	DefaultCountryCode     string          `json:"default_country_code,omitempty"`
	Currency               string          `json:"currency,omitempty"`
	Steps                  int             `json:"steps,omitempty"`
	Price                  *Price          `json:"price,omitempty"`
	Labels                 *PropertyLabels `json:"labels,omitempty"`
}

type Price struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type PropertyLabels struct {
	Left   string `json:"left,omitempty"`
	Center string `json:"center,omitempty"`
	Right  string `json:"right,omitempty"`
}

type FieldChoice struct {
	ID         string      `json:"id,omitempty"`
	Ref        string      `json:"ref,omitempty"`
	Label      string      `json:"label"`
	Attachment *Attachment `json:"attachment,omitempty"`
}

type FieldValidations struct {
	Required     bool `json:"required"`
	MaxLength    *int `json:"max_length,omitempty"`
	MinSelection *int `json:"min_selection,omitempty"`
	MaxSelection *int `json:"max_selection,omitempty"`
	MinValue     *int `json:"min_value,omitempty"`
	MaxValue     *int `json:"max_value,omitempty"`
}

type Logic struct {
	Type    string          `json:"type"`
	Ref     string          `json:"ref,omitempty"`
	Actions []*LogicActions `json:"actions"`
}

type LogicActions struct {
	Action    string             `json:"action"`
	Details   LogicActionDetails `json:"details"`
	Condition LogicCondition     `json:"condition"`
}

type LogicActionDetails struct {
	To     *LogicJumpDestination `json:"to,omitempty"`
	Target *TargetVariable       `json:"target,omitempty"`
	Value  *VariableValue        `json:"value,omitempty"`
}

type LogicJumpDestination struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type TargetVariable struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type VariableValue struct {
	Type       string           `json:"type"`
	Value      interface{}      `json:"value,omitempty"`
	Evaluation *LogicEvaluation `json:"evaluation,omitempty"`
}

type LogicEvaluation struct {
	Op   string            `json:"op"`
	Vars []*TargetVariable `json:"vars"`
}

type LogicCondition struct {
	Op   string      `json:"op"`
	Vars []*LogicVar `json:"vars"`
}

type LogicVar struct {
	*LogicCondition
	Type  string      `json:"type,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

type Outcome struct {
	Variable string           `json:"variable"`
	Choices  []*OutcomeChoice `json:"choices"`
}

type OutcomeChoice struct {
	ID                string `json:"id,omitempty"`
	Ref               string `json:"ref,omitempty"`
	Title             string `json:"title"`
	CounterVariable   string `json:"counter_variable"`
	ThankYouScreenRef string `json:"thankyou_screen_ref"`
}
