package models

type License struct {
	ID   string `json:"id" url:"id,omitempty"`     // a short string identifying the license.Example: cc_by_sa
	Name string `json:"name" url:"name,omitempty"` // the name of the license.Example: CC Attribution ShareAlike
	Url  string `json:"url" url:"url,omitempty"`   // a link to the license text.Example: http://creativecommons.org/licenses/by-sa/4.0
}

func (t *License) HasError() error {
	return nil
}
