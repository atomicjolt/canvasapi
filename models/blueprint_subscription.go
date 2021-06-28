package models

type BlueprintSubscription struct {
	ID              int64  `json:"id" url:"id,omitempty"`                             // The ID of the blueprint course subscription.Example: 101
	TemplateID      int64  `json:"template_id" url:"template_id,omitempty"`           // The ID of the blueprint template the associated course is subscribed to.Example: 1
	BlueprintCourse string `json:"blueprint_course" url:"blueprint_course,omitempty"` // The blueprint course subscribed to.Example: 2, Biology 100 Blueprint, BIOL 100 BP, Default term
}

func (t *BlueprintSubscription) HasError() error {
	return nil
}
