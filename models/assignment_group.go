package models

type AssignmentGroup struct {
	ID              int64         `json:"id"`               // the id of the Assignment Group.Example: 1
	Name            string        `json:"name"`             // the name of the Assignment Group.Example: group2
	Position        int64         `json:"position"`         // the position of the Assignment Group.Example: 7
	GroupWeight     int64         `json:"group_weight"`     // the weight of the Assignment Group.Example: 20
	SISSourceID     string        `json:"sis_source_id"`    // the sis source id of the Assignment Group.Example: 1234
	IntegrationData string        `json:"integration_data"` // the integration data of the Assignment Group.Example: 0954
	Assignments     []int64       `json:"assignments"`      // the assignments in this Assignment Group (see the Assignment API for a detailed list of fields).
	Rules           *GradingRules `json:"rules"`            // the grading rules that this Assignment Group has.
}

func (t *AssignmentGroup) HasError() error {
	return nil
}
