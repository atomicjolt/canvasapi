package models

type License struct {
	ID   string `json:"id"`   // a short string identifying the license.Example: cc_by_sa
	Name string `json:"name"` // the name of the license.Example: CC Attribution ShareAlike
	Url  string `json:"url"`  // a link to the license text.Example: http://creativecommons.org/licenses/by-sa/4.0
}

func (t *License) HasError() error {
	return nil
}
