package models

type Migrator struct {
	Type               string   `json:"type"`                 // The value to pass to the create endpoint.Example: common_cartridge_importer
	RequiresFileUpload bool     `json:"requires_file_upload"` // Whether this endpoint requires a file upload.Example: true
	Name               string   `json:"name"`                 // Description of the package type expected.Example: Common Cartridge 1.0/1.1/1.2 Package
	RequiredSettings   []string `json:"required_settings"`    // A list of fields this system requires.Example: source_course_id
}

func (t *Migrator) HasError() error {
	return nil
}
