package template

import "strconv"

type Versions []Version

type Version struct {
	Id          string `json:"id"`
	TemplateId  string `json:"template_id"`
	Active      int    `json:"active"`
	Name        string `json:"name"`
	HtmlContent string `json:"html_content"`
	Subject     string `json:"subject"`
}

type Template struct {
	Id         string   `json:"id"`
	Name       string   `json:"name"`
	Generation string   `json:"generation"`
	UpdatedAt  string   `json:"updated_at"`
	Versions   Versions `json:"versions"`
}

type Templates struct {
	Templates []Template `json:"templates"`
}

// Error reports an error and the operation and URL that caused it.
type Error struct {
	HttpStatus int
	Message    string
}

func (e *Error) Error() string {
	return "HTTP Response Status: " + strconv.Itoa(e.HttpStatus) + ", message: " + e.Message
}
