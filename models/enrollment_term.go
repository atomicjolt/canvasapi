package models

import (
	"time"
)

type EnrollmentTerm struct {
	ID            int64                    `json:"id" url:"id,omitempty"`                         // The unique identifier for the enrollment term..Example: 1
	SISTermID     string                   `json:"sis_term_id" url:"sis_term_id,omitempty"`       // The SIS id of the term. Only included if the user has permission to view SIS information..Example: Sp2014
	SISImportID   int64                    `json:"sis_import_id" url:"sis_import_id,omitempty"`   // the unique identifier for the SIS import. This field is only included if the user has permission to manage SIS information..Example: 34
	Name          string                   `json:"name" url:"name,omitempty"`                     // The name of the term..Example: Spring 2014
	StartAt       time.Time                `json:"start_at" url:"start_at,omitempty"`             // The datetime of the start of the term..Example: 2014-01-06T08:00:00-05:00
	EndAt         time.Time                `json:"end_at" url:"end_at,omitempty"`                 // The datetime of the end of the term..Example: 2014-05-16T05:00:00-04:00
	WorkflowState string                   `json:"workflow_state" url:"workflow_state,omitempty"` // The state of the term. Can be 'active' or 'deleted'..Example: active
	Overrides     map[string](interface{}) `json:"overrides" url:"overrides,omitempty"`           // Term date overrides for specific enrollment types.Example: {'start_at'=>'2014-01-07T08:00:00-05:00', 'end_at'=>'2014-05-14T05:00:00-04:0'}
}

func (t *EnrollmentTerm) HasErrors() error {
	return nil
}
