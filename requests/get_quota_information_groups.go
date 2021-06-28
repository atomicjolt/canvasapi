package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// GetQuotaInformationGroups Returns the total and used storage quota for the course, group, or user.
// https://canvas.instructure.com/doc/api/files.html
//
// Path Parameters:
// # GroupID (Required) ID
//
type GetQuotaInformationGroups struct {
	Path struct {
		GroupID string `json:"group_id" url:"group_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *GetQuotaInformationGroups) GetMethod() string {
	return "GET"
}

func (t *GetQuotaInformationGroups) GetURLPath() string {
	path := "groups/{group_id}/files/quota"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	return path
}

func (t *GetQuotaInformationGroups) GetQuery() (string, error) {
	return "", nil
}

func (t *GetQuotaInformationGroups) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetQuotaInformationGroups) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetQuotaInformationGroups) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'GroupID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetQuotaInformationGroups) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
