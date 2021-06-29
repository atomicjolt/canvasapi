package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// ShowCollectionOfResults Show existing Results of a line item. Can be used to retrieve a specific student's
// result by adding the user_id (defined as the lti_user_id or the Canvas user_id) as
// a query parameter (i.e. user_id=1000). If user_id is included, it will return only
// one Result in the collection if the result exists, otherwise it will be empty. May
// also limit number of results by adding the limit query param (i.e. limit=100)
// https://canvas.instructure.com/doc/api/result.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.LineItemID (Required) ID
//
type ShowCollectionOfResults struct {
	Path struct {
		CourseID   string `json:"course_id" url:"course_id,omitempty"`       //  (Required)
		LineItemID string `json:"line_item_id" url:"line_item_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ShowCollectionOfResults) GetMethod() string {
	return "GET"
}

func (t *ShowCollectionOfResults) GetURLPath() string {
	path := "/lti/courses/{course_id}/line_items/{line_item_id}/results"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{line_item_id}", fmt.Sprintf("%v", t.Path.LineItemID))
	return path
}

func (t *ShowCollectionOfResults) GetQuery() (string, error) {
	return "", nil
}

func (t *ShowCollectionOfResults) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ShowCollectionOfResults) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ShowCollectionOfResults) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Path.LineItemID == "" {
		errs = append(errs, "'Path.LineItemID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ShowCollectionOfResults) Do(c *canvasapi.Canvas) (*models.Result, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.Result{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
