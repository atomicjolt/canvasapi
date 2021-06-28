package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// RemoveUsageRightsUsers Removes copyright and license information associated with one or more files
// https://canvas.instructure.com/doc/api/files.html
//
// Path Parameters:
// # UserID (Required) ID
//
// Query Parameters:
// # FileIDs (Required) List of ids of files to remove associated usage rights from.
// # FolderIDs (Optional) List of ids of folders. Usage rights will be removed from all files in these folders.
//
type RemoveUsageRightsUsers struct {
	Path struct {
		UserID string `json:"user_id" url:"user_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		FileIDs   []string `json:"file_ids" url:"file_ids,omitempty"`     //  (Required)
		FolderIDs []string `json:"folder_ids" url:"folder_ids,omitempty"` //  (Optional)
	} `json:"query"`
}

func (t *RemoveUsageRightsUsers) GetMethod() string {
	return "DELETE"
}

func (t *RemoveUsageRightsUsers) GetURLPath() string {
	path := "users/{user_id}/usage_rights"
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	return path
}

func (t *RemoveUsageRightsUsers) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *RemoveUsageRightsUsers) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *RemoveUsageRightsUsers) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *RemoveUsageRightsUsers) HasErrors() error {
	errs := []string{}
	if t.Path.UserID == "" {
		errs = append(errs, "'UserID' is required")
	}
	if t.Query.FileIDs == nil {
		errs = append(errs, "'FileIDs' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *RemoveUsageRightsUsers) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
