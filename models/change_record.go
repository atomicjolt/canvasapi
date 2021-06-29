package models

type ChangeRecord struct {
	AssetID    int64    `json:"asset_id" url:"asset_id,omitempty"`       // The ID of the learning object that was changed in the blueprint course..Example: 2
	AssetType  string   `json:"asset_type" url:"asset_type,omitempty"`   // The type of the learning object that was changed in the blueprint course.  One of 'assignment', 'attachment', 'discussion_topic', 'external_tool', 'quiz', 'wiki_page', 'syllabus', or 'settings'.  For 'syllabus' or 'settings', the asset_id is the course id..Example: assignment
	AssetName  string   `json:"asset_name" url:"asset_name,omitempty"`   // The name of the learning object that was changed in the blueprint course..Example: Some Assignment
	ChangeType string   `json:"change_type" url:"change_type,omitempty"` // The type of change; one of 'created', 'updated', 'deleted'.Example: created
	HtmlUrl    string   `json:"html_url" url:"html_url,omitempty"`       // The URL of the changed object.Example: https://canvas.example.com/courses/101/assignments/2
	Locked     bool     `json:"locked" url:"locked,omitempty"`           // Whether the object is locked in the blueprint.
	Exceptions []string `json:"exceptions" url:"exceptions,omitempty"`   // A list of ExceptionRecords for linked courses that did not receive this update..Example: {'course_id'=>101, 'conflicting_changes'=>['points']}
}

func (t *ChangeRecord) HasErrors() error {
	return nil
}
