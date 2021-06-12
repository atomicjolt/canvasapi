package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// RemoveUsageRightsGroups Removes copyright and license information associated with one or more files
// https://canvas.instructure.com/doc/api/files.html
//
// Path Parameters:
// # GroupID (Required) ID
//
// Query Parameters:
// # FileIDs (Required) List of ids of files to remove associated usage rights from.
// # FolderIDs (Optional) List of ids of folders. Usage rights will be removed from all files in these folders.
//
type RemoveUsageRightsGroups struct {
	Path struct {
		GroupID string `json:"group_id"` //  (Required)
	} `json:"path"`

	Query struct {
		FileIDs   []string `json:"file_ids"`   //  (Required)
		FolderIDs []string `json:"folder_ids"` //  (Optional)
	} `json:"query"`
}

func (t *RemoveUsageRightsGroups) GetMethod() string {
	return "DELETE"
}

func (t *RemoveUsageRightsGroups) GetURLPath() string {
	path := "groups/{group_id}/usage_rights"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	return path
}

func (t *RemoveUsageRightsGroups) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *RemoveUsageRightsGroups) GetBody() (string, error) {
	return "", nil
}

func (t *RemoveUsageRightsGroups) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'GroupID' is required")
	}
	if t.Query.FileIDs == nil {
		errs = append(errs, "'FileIDs' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *RemoveUsageRightsGroups) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}