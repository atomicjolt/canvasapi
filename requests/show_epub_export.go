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

// ShowEpubExport Get information about a single ePub export.
// https://canvas.instructure.com/doc/api/e_pub_exports.html
//
// Path Parameters:
// # CourseID (Required) ID
// # ID (Required) ID
//
type ShowEpubExport struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		ID       string `json:"id" url:"id,omitempty"`               //  (Required)
	} `json:"path"`
}

func (t *ShowEpubExport) GetMethod() string {
	return "GET"
}

func (t *ShowEpubExport) GetURLPath() string {
	path := "courses/{course_id}/epub_exports/{id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *ShowEpubExport) GetQuery() (string, error) {
	return "", nil
}

func (t *ShowEpubExport) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ShowEpubExport) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ShowEpubExport) HasErrors() error {
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

func (t *ShowEpubExport) Do(c *canvasapi.Canvas) (*models.EpubExport, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.EpubExport{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
