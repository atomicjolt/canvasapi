package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// RevertToRevisionGroups Revert a page to a prior revision.
// https://canvas.instructure.com/doc/api/pages.html
//
// Path Parameters:
// # GroupID (Required) ID
// # Url (Required) ID
// # RevisionID (Required) The revision to revert to (use the
//    {api:WikiPagesApiController#revisions List Revisions API} to see
//    available revisions)
//
type RevertToRevisionGroups struct {
	Path struct {
		GroupID    string `json:"group_id"`    //  (Required)
		Url        string `json:"url"`         //  (Required)
		RevisionID int64  `json:"revision_id"` //  (Required)
	} `json:"path"`
}

func (t *RevertToRevisionGroups) GetMethod() string {
	return "POST"
}

func (t *RevertToRevisionGroups) GetURLPath() string {
	path := "groups/{group_id}/pages/{url}/revisions/{revision_id}"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	path = strings.ReplaceAll(path, "{url}", fmt.Sprintf("%v", t.Path.Url))
	path = strings.ReplaceAll(path, "{revision_id}", fmt.Sprintf("%v", t.Path.RevisionID))
	return path
}

func (t *RevertToRevisionGroups) GetQuery() (string, error) {
	return "", nil
}

func (t *RevertToRevisionGroups) GetBody() (string, error) {
	return "", nil
}

func (t *RevertToRevisionGroups) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'GroupID' is required")
	}
	if t.Path.Url == "" {
		errs = append(errs, "'Url' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *RevertToRevisionGroups) Do(c *canvasapi.Canvas) (*models.PageRevision, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.PageRevision{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
