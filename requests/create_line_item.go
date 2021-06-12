package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// CreateLineItem Create a new Line Item
// https://canvas.instructure.com/doc/api/line_items.html
//
// Path Parameters:
// # CourseID (Required) ID
//
// Form Parameters:
// # ScoreMaximum (Required) The maximum score for the line item. Scores created for the Line Item may exceed this value.
// # Label (Required) The label for the Line Item. If no resourceLinkId is specified this value will also be used
//    as the name of the placeholder assignment.
// # ResourceID (Optional) A Tool Provider specified id for the Line Item. Multiple line items may
//    share the same resourceId within a given context.
// # Tag (Optional) A value used to qualify a line Item beyond its ids. Line Items may be queried
//    by this value in the List endpoint. Multiple line items can share the same tag
//    within a given context.
// # ResourceLinkID (Optional) The resource link id the Line Item should be attached to. This value should
//    match the LTI id of the Canvas assignment associated with the tool.
// # CanvasLTISubmissionType (Optional) (EXTENSION) - Optional block to set Assignment Submission Type when creating a new assignment is created.
//    type - 'none' or 'external_tool'::
//    external_tool_url - Submission URL only used when type: 'external_tool'::
//
type CreateLineItem struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
	} `json:"path"`

	Form struct {
		ScoreMaximum            float64 `json:"score_maximum"`                                      //  (Required)
		Label                   string  `json:"label"`                                              //  (Required)
		ResourceID              string  `json:"resource_id"`                                        //  (Optional)
		Tag                     string  `json:"tag"`                                                //  (Optional)
		ResourceLinkID          string  `json:"resource_link_id"`                                   //  (Optional)
		CanvasLTISubmissionType string  `json:"https://canvas.instructure.com/lti/submission_type"` //  (Optional)
	} `json:"form"`
}

func (t *CreateLineItem) GetMethod() string {
	return "POST"
}

func (t *CreateLineItem) GetURLPath() string {
	path := "/lti/courses/{course_id}/line_items"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *CreateLineItem) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateLineItem) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *CreateLineItem) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Form.Label == "" {
		errs = append(errs, "'Label' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateLineItem) Do(c *canvasapi.Canvas) (*models.LineItem, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.LineItem{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}