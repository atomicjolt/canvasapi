package models

type EnrollmentTermsList struct {
	EnrollmentTerms []*EnrollmentTerm `json:"enrollment_terms"` // a paginated list of all terms in the account.
}

func (t *EnrollmentTermsList) HasError() error {
	return nil
}
