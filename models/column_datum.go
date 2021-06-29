package models

type ColumnDatum struct {
	Content string `json:"content" url:"content,omitempty"` // Example: Nut allergy
	UserID  int64  `json:"user_id" url:"user_id,omitempty"` // Example: 2
}

func (t *ColumnDatum) HasErrors() error {
	return nil
}
