package models

type BlueprintSubscription struct {
	ID              int64  `json:"id"`               // The ID of the blueprint course subscription.Example: 101
	TemplateID      int64  `json:"template_id"`      // The ID of the blueprint template the associated course is subscribed to.Example: 1
	BlueprintCourse string `json:"blueprint_course"` // The blueprint course subscribed to.Example: 2, Biology 100 Blueprint, BIOL 100 BP, Default term
}

func (t *BlueprintSubscription) HasError() error {
	return nil
}
