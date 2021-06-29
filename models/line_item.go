package models

type LineItem struct {
	ID             string  `json:"id" url:"id,omitempty"`                             // The fully qualified URL for showing, updating, and deleting the Line Item.Example: http://institution.canvas.com/api/lti/courses/5/line_items/2
	ScoreMaximum   float64 `json:"score_maximum" url:"score_maximum,omitempty"`       // The maximum score of the Line Item.Example: 50
	Label          string  `json:"label" url:"label,omitempty"`                       // The label of the Line Item..Example: 50
	Tag            string  `json:"tag" url:"tag,omitempty"`                           // Tag used to qualify a line Item beyond its ids.Example: 50
	ResourceID     string  `json:"resource_id" url:"resource_id,omitempty"`           // A Tool Provider specified id for the Line Item. Multiple line items can share the same resourceId within a given context.Example: 50
	ResourceLinkID string  `json:"resource_link_id" url:"resource_link_id,omitempty"` // The resource link id the Line Item is attached to.Example: 50
	SubmissionType string  `json:"submission_type" url:"submission_type,omitempty"`   // The extension that defines the submission_type of the line_item. Only returns if set through the line_item create endpoint..Example: { 	'type':'external_tool', 	'external_tool_url':'https://my.launch.url', }
}

func (t *LineItem) HasErrors() error {
	return nil
}
