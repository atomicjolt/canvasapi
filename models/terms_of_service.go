package models

type TermsOfService struct {
	ID                   int64  `json:"id" url:"id,omitempty"`                                         // Terms Of Service id.Example: 1
	TermsType            string `json:"terms_type" url:"terms_type,omitempty"`                         // The given type for the Terms of Service.Example: default
	Passive              bool   `json:"passive" url:"passive,omitempty"`                               // Boolean dictating if the user must accept Terms of Service.
	AccountID            int64  `json:"account_id" url:"account_id,omitempty"`                         // The id of the root account that owns the Terms of Service.Example: 1
	Content              string `json:"content" url:"content,omitempty"`                               // Content of the Terms of Service.Example: To be or not to be that is the question
	SelfRegistrationType string `json:"self_registration_type" url:"self_registration_type,omitempty"` // The type of self registration allowed.Example: none, observer, all
}

func (t *TermsOfService) HasErrors() error {
	return nil
}
