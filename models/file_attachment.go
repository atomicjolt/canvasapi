package models

type FileAttachment struct {
	Contenttype string `json:"content_type"` // Example: unknown/unknown
	Url         string `json:"url"`          // Example: http://www.example.com/courses/1/files/1/download
	Filename    string `json:"filename"`     // Example: content.txt
	DisplayName string `json:"display_name"` // Example: content.txt
}

func (t *FileAttachment) HasError() error {
	return nil
}
