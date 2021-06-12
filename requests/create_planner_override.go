package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
	"github.com/atomicjolt/string_utils"
)

// CreatePlannerOverride Create a planner override for the current user
// https://canvas.instructure.com/doc/api/planner.html
//
// Form Parameters:
// # PlannableType (Required) . Must be one of announcement, assignment, discussion_topic, quiz, wiki_page, planner_noteType of the item that you are overriding in the planner
// # PlannableID (Required) ID of the item that you are overriding in the planner
// # MarkedComplete (Optional) If this is true, the item will show in the planner as completed
// # Dismissed (Optional) If this is true, the item will not show in the opportunities list
//
type CreatePlannerOverride struct {
	Form struct {
		PlannableType  string `json:"plannable_type"`  //  (Required) . Must be one of announcement, assignment, discussion_topic, quiz, wiki_page, planner_note
		PlannableID    int64  `json:"plannable_id"`    //  (Required)
		MarkedComplete bool   `json:"marked_complete"` //  (Optional)
		Dismissed      bool   `json:"dismissed"`       //  (Optional)
	} `json:"form"`
}

func (t *CreatePlannerOverride) GetMethod() string {
	return "POST"
}

func (t *CreatePlannerOverride) GetURLPath() string {
	return ""
}

func (t *CreatePlannerOverride) GetQuery() (string, error) {
	return "", nil
}

func (t *CreatePlannerOverride) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *CreatePlannerOverride) HasErrors() error {
	errs := []string{}
	if t.Form.PlannableType == "" {
		errs = append(errs, "'PlannableType' is required")
	}
	if !string_utils.Include([]string{"announcement", "assignment", "discussion_topic", "quiz", "wiki_page", "planner_note"}, t.Form.PlannableType) {
		errs = append(errs, "PlannableType must be one of announcement, assignment, discussion_topic, quiz, wiki_page, planner_note")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreatePlannerOverride) Do(c *canvasapi.Canvas) (*models.PlannerOverride, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.PlannerOverride{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}