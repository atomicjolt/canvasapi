package models

type PageViewLinks struct {
	User     int64 `json:"user"`      // The ID of the user for this page view.Example: 1234
	Context  int64 `json:"context"`   // The ID of the context for the request (course id if context_type is Course, etc).Example: 1234
	Asset    int64 `json:"asset"`     // The ID of the asset for the request, if any.Example: 1234
	RealUser int64 `json:"real_user"` // The ID of the actual user who made this request, if the request was made by a user who was masquerading.Example: 1234
	Account  int64 `json:"account"`   // The ID of the account context for this page view.Example: 1234
}

func (t *PageViewLinks) HasError() error {
	return nil
}
