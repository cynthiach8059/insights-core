package email

// Email struct that contains data to send email
type Email struct {
	From      string
	FromName  string
	To        string
	Subject   string
	Body      string
	BodyType  BodyType
	AttachURL interface{}
}

// BodyType enum to select body type in email
type BodyType string

const (
	BodyTextPlain BodyType = "text/plain"
	BodyHtml      BodyType = "text/html"
)
