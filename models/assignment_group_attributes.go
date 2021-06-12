package models

type AssignmentGroupAttributes struct {
	ID              int64  `json:"id"`               // the id of the Assignment Group.Example: 1
	Name            string `json:"name"`             // the name of the Assignment Group.Example: group2
	GroupWeight     int64  `json:"group_weight"`     // the weight of the Assignment Group.Example: 20
	SISSourceID     string `json:"sis_source_id"`    // the sis source id of the Assignment Group.Example: 1234
	IntegrationData string `json:"integration_data"` // the integration data of the Assignment Group.Example: 0954
}

func (t *AssignmentGroupAttributes) HasError() error {
	return nil
}
