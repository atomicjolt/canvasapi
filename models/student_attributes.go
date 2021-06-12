package models

type StudentAttributes struct {
	UserID    int64  `json:"user_id"`     // The unique Canvas identifier for the user.Example: 511
	SISUserID string `json:"sis_user_id"` // The SIS ID associated with the user.  This field is only included if the user came from a SIS import and has permissions to view SIS information..Example: SHEL93921
}

func (t *StudentAttributes) HasError() error {
	return nil
}