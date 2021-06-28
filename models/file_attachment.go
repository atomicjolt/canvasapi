package models

type FileAttachment struct {
	Contenttype string `json:"content_type" url:"content_type,omitempty"` // Example: unknown/unknown
	Url         string `json:"url" url:"url,omitempty"`                   // Example: http://www.example.com/courses/1/files/1/download
	Filename    string `json:"filename" url:"filename,omitempty"`         // Example: content.txt
	DisplayName string `json:"display_name" url:"display_name,omitempty"` // Example: content.txt
}

func (t *FileAttachment) HasError() error {
	return nil
}
