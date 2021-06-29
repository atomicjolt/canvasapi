package models

type AssignmentGroupAttributes struct {
	ID              int64                    `json:"id" url:"id,omitempty"`                             // the id of the Assignment Group.Example: 1
	Name            string                   `json:"name" url:"name,omitempty"`                         // the name of the Assignment Group.Example: group2
	GroupWeight     int64                    `json:"group_weight" url:"group_weight,omitempty"`         // the weight of the Assignment Group.Example: 20
	SISSourceID     string                   `json:"sis_source_id" url:"sis_source_id,omitempty"`       // the sis source id of the Assignment Group.Example: 1234
	IntegrationData map[string](interface{}) `json:"integration_data" url:"integration_data,omitempty"` // the integration data of the Assignment Group.Example: 0954
}

func (t *AssignmentGroupAttributes) HasErrors() error {
	return nil
}
