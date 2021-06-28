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

// ImportCategoryGroups Create Groups in a Group Category through a CSV import
//
// For more information on the format that's expected here, please see the
// "Group Category CSV" section in the API docs.
// https://canvas.instructure.com/doc/api/group_categories.html
//
// Path Parameters:
// # GroupCategoryID (Required) ID
//
// Form Parameters:
// # Attachment (Optional) There are two ways to post group category import data - either via a
//    multipart/form-data form-field-style attachment, or via a non-multipart
//    raw post request.
//
//    'attachment' is required for multipart/form-data style posts. Assumed to
//    be outcome data from a file upload form field named 'attachment'.
//
//    Examples:
//      curl -F attachment=@<filename> -H "Authorization: Bearer <token>" \
//          'https://<canvas>/api/v1/group_categories/<category_id>/import'
//
//    If you decide to do a raw post, you can skip the 'attachment' argument,
//    but you will then be required to provide a suitable Content-Type header.
//    You are encouraged to also provide the 'extension' argument.
//
//    Examples:
//      curl -H 'Content-Type: text/csv' --data-binary @<filename>.csv \
//          -H "Authorization: Bearer <token>" \
//          'https://<canvas>/api/v1/group_categories/<category_id>/import'
//
type ImportCategoryGroups struct {
	Path struct {
		GroupCategoryID string `json:"group_category_id" url:"group_category_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		Attachment string `json:"attachment" url:"attachment,omitempty"` //  (Optional)
	} `json:"form"`
}

func (t *ImportCategoryGroups) GetMethod() string {
	return "POST"
}

func (t *ImportCategoryGroups) GetURLPath() string {
	path := "group_categories/{group_category_id}/import"
	path = strings.ReplaceAll(path, "{group_category_id}", fmt.Sprintf("%v", t.Path.GroupCategoryID))
	return path
}

func (t *ImportCategoryGroups) GetQuery() (string, error) {
	return "", nil
}

func (t *ImportCategoryGroups) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *ImportCategoryGroups) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *ImportCategoryGroups) HasErrors() error {
	errs := []string{}
	if t.Path.GroupCategoryID == "" {
		errs = append(errs, "'GroupCategoryID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ImportCategoryGroups) Do(c *canvasapi.Canvas) (*models.Progress, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.Progress{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
