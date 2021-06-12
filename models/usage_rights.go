package models

type UsageRights struct {
	LegalCopyright   string  `json:"legal_copyright"`   // Copyright line for the file.Example: (C) 2014 Incom Corporation Ltd
	UseJustification string  `json:"use_justification"` // Justification for using the file in a Canvas course. Valid values are 'own_copyright', 'public_domain', 'used_by_permission', 'fair_use', 'creative_commons'.Example: creative_commons
	License          string  `json:"license"`           // License identifier for the file..Example: cc_by_sa
	LicenseName      string  `json:"license_name"`      // Readable license name.Example: CC Attribution Share-Alike
	Message          string  `json:"message"`           // Explanation of the action performed.Example: 4 files updated
	FileIDs          []int64 `json:"file_ids"`          // List of ids of files that were updated.Example: 1, 2, 3
}

func (t *UsageRights) HasError() error {
	return nil
}