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

// UpdateLineItem Update new Line Item
// https://canvas.instructure.com/doc/api/line_items.html
//
// Path Parameters:
// # CourseID (Required) ID
// # ID (Required) ID
//
// Form Parameters:
// # ScoreMaximum (Optional) The maximum score for the line item. Scores created for the Line Item may exceed this value.
// # Label (Optional) The label for the Line Item. If no resourceLinkId is specified this value will also be used
//    as the name of the placeholder assignment.
// # ResourceID (Optional) A Tool Provider specified id for the Line Item. Multiple line items may
//    share the same resourceId within a given context.
// # Tag (Optional) A value used to qualify a line Item beyond its ids. Line Items may be queried
//    by this value in the List endpoint. Multiple line items can share the same tag
//    within a given context.
//
type UpdateLineItem struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
		ID       string `json:"id"`        //  (Required)
	} `json:"path"`

	Form struct {
		ScoreMaximum float64 `json:"score_maximum"` //  (Optional)
		Label        string  `json:"label"`         //  (Optional)
		ResourceID   string  `json:"resource_id"`   //  (Optional)
		Tag          string  `json:"tag"`           //  (Optional)
	} `json:"form"`
}

func (t *UpdateLineItem) GetMethod() string {
	return "PUT"
}

func (t *UpdateLineItem) GetURLPath() string {
	path := "/lti/courses/{course_id}/line_items/{id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *UpdateLineItem) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdateLineItem) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *UpdateLineItem) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UpdateLineItem) Do(c *canvasapi.Canvas) (*models.LineItem, error) {
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
