package models

type ColumnDatum struct {
	Content string `json:"content"` // Example: Nut allergy
	UserID  int64  `json:"user_id"` // Example: 2
}

func (t *ColumnDatum) HasError() error {
	return nil
}
