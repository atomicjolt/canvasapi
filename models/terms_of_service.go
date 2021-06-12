package models

type TermsOfService struct {
	ID                   int64  `json:"id"`                     // Terms Of Service id.Example: 1
	TermsType            string `json:"terms_type"`             // The given type for the Terms of Service.Example: default
	Passive              bool   `json:"passive"`                // Boolean dictating if the user must accept Terms of Service.
	AccountID            int64  `json:"account_id"`             // The id of the root account that owns the Terms of Service.Example: 1
	Content              string `json:"content"`                // Content of the Terms of Service.Example: To be or not to be that is the question
	SelfRegistrationType string `json:"self_registration_type"` // The type of self registration allowed.Example: none, observer, all
}

func (t *TermsOfService) HasError() error {
	return nil
}
