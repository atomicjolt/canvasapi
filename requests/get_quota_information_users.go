package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// GetQuotaInformationUsers Returns the total and used storage quota for the course, group, or user.
// https://canvas.instructure.com/doc/api/files.html
//
// Path Parameters:
// # UserID (Required) ID
//
type GetQuotaInformationUsers struct {
	Path struct {
		UserID string `json:"user_id" url:"user_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *GetQuotaInformationUsers) GetMethod() string {
	return "GET"
}

func (t *GetQuotaInformationUsers) GetURLPath() string {
	path := "users/{user_id}/files/quota"
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	return path
}

func (t *GetQuotaInformationUsers) GetQuery() (string, error) {
	return "", nil
}

func (t *GetQuotaInformationUsers) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetQuotaInformationUsers) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetQuotaInformationUsers) HasErrors() error {
	errs := []string{}
	if t.Path.UserID == "" {
		errs = append(errs, "'UserID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetQuotaInformationUsers) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
