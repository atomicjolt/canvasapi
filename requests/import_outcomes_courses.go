package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// ImportOutcomesCourses Import outcomes into Canvas.
//
// For more information on the format that's expected here, please see the
// "Outcomes CSV" section in the API docs.
// https://canvas.instructure.com/doc/api/outcome_imports.html
//
// Path Parameters:
// # CourseID (Required) ID
//
// Form Parameters:
// # ImportType (Optional) Choose the data format for reading outcome data. With a standard Canvas
//    install, this option can only be 'instructure_csv', and if unprovided,
//    will be assumed to be so. Can be part of the query string.
// # Attachment (Optional) There are two ways to post outcome import data - either via a
//    multipart/form-data form-field-style attachment, or via a non-multipart
//    raw post request.
//
//    'attachment' is required for multipart/form-data style posts. Assumed to
//    be outcome data from a file upload form field named 'attachment'.
//
//    Examples:
//      curl -F attachment=@<filename> -H "Authorization: Bearer <token>" \
//          'https://<canvas>/api/v1/accounts/<account_id>/outcome_imports?import_type=instructure_csv'
//      curl -F attachment=@<filename> -H "Authorization: Bearer <token>" \
//          'https://<canvas>/api/v1/courses/<course_id>/outcome_imports?import_type=instructure_csv'
//
//    If you decide to do a raw post, you can skip the 'attachment' argument,
//    but you will then be required to provide a suitable Content-Type header.
//    You are encouraged to also provide the 'extension' argument.
//
//    Examples:
//      curl -H 'Content-Type: text/csv' --data-binary @<filename>.csv \
//          -H "Authorization: Bearer <token>" \
//          'https://<canvas>/api/v1/accounts/<account_id>/outcome_imports?import_type=instructure_csv'
//
//      curl -H 'Content-Type: text/csv' --data-binary @<filename>.csv \
//          -H "Authorization: Bearer <token>" \
//          'https://<canvas>/api/v1/courses/<course_id>/outcome_imports?import_type=instructure_csv'
// # Extension (Optional) Recommended for raw post request style imports. This field will be used to
//    distinguish between csv and other file format extensions that
//    would usually be provided with the filename in the multipart post request
//    scenario. If not provided, this value will be inferred from the
//    Content-Type, falling back to csv-file format if all else fails.
//
type ImportOutcomesCourses struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		ImportType string `json:"import_type" url:"import_type,omitempty"` //  (Optional)
		Attachment string `json:"attachment" url:"attachment,omitempty"`   //  (Optional)
		Extension  string `json:"extension" url:"extension,omitempty"`     //  (Optional)
	} `json:"form"`
}

func (t *ImportOutcomesCourses) GetMethod() string {
	return "POST"
}

func (t *ImportOutcomesCourses) GetURLPath() string {
	path := "courses/{course_id}/outcome_imports"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *ImportOutcomesCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *ImportOutcomesCourses) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *ImportOutcomesCourses) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *ImportOutcomesCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ImportOutcomesCourses) Do(c *canvasapi.Canvas) (*models.OutcomeImport, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.OutcomeImport{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
