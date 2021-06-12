package models

import (
	"fmt"
	"time"

	"github.com/atomicjolt/string_utils"
)

type Module struct {
	ID                        int64         `json:"id"`                          // the unique identifier for the module.Example: 123
	WorkflowState             string        `json:"workflow_state"`              // the state of the module: 'active', 'deleted'.Example: active
	Position                  int64         `json:"position"`                    // the position of this module in the course (1-based).Example: 2
	Name                      string        `json:"name"`                        // the name of this module.Example: Imaginary Numbers and You
	UnlockAt                  time.Time     `json:"unlock_at"`                   // (Optional) the date this module will unlock.Example: 2012-12-31T06:00:00-06:00
	RequireSequentialProgress bool          `json:"require_sequential_progress"` // Whether module items must be unlocked in order.Example: true
	PrerequisiteModuleIDs     []int64       `json:"prerequisite_module_ids"`     // IDs of Modules that must be completed before this one is unlocked.Example: 121, 122
	ItemsCount                int64         `json:"items_count"`                 // The number of items in the module.Example: 10
	ItemsUrl                  string        `json:"items_url"`                   // The API URL to retrive this module's items.Example: https://canvas.example.com/api/v1/modules/123/items
	Items                     []*ModuleItem `json:"items"`                       // The contents of this module, as an array of Module Items. (Present only if requested via include[]=items AND the module is not deemed too large by Canvas.).
	State                     string        `json:"state"`                       // The state of this Module for the calling user one of 'locked', 'unlocked', 'started', 'completed' (Optional; present only if the caller is a student or if the optional parameter 'student_id' is included).Example: started
	CompletedAt               time.Time     `json:"completed_at"`                // the date the calling user completed the module (Optional; present only if the caller is a student or if the optional parameter 'student_id' is included).
	PublishFinalGrade         bool          `json:"publish_final_grade"`         // if the student's final grade for the course should be published to the SIS upon completion of this module.
	Published                 bool          `json:"published"`                   // (Optional) Whether this module is published. This field is present only if the caller has permission to view unpublished modules..Example: true
}

func (t *Module) HasError() error {
	var s []string
	s = []string{"active", "deleted"}
	if !string_utils.Include(s, t.WorkflowState) {
		return fmt.Errorf("expected 'workflow_state' to be one of %v", s)
	}
	s = []string{"locked", "unlocked", "started", "completed"}
	if !string_utils.Include(s, t.State) {
		return fmt.Errorf("expected 'state' to be one of %v", s)
	}
	return nil
}
