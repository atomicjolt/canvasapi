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
	"github.com/atomicjolt/string_utils"
)

// CreateGroupGroupCategories Creates a new group. Groups created using the "/api/v1/groups/"
// endpoint will be community groups.
// https://canvas.instructure.com/doc/api/groups.html
//
// Path Parameters:
// # GroupCategoryID (Required) ID
//
// Form Parameters:
// # Name (Optional) The name of the group
// # Description (Optional) A description of the group
// # IsPublic (Optional) whether the group is public (applies only to community groups)
// # JoinLevel (Optional) . Must be one of parent_context_auto_join, parent_context_request, invitation_onlyno description
// # StorageQuotaMb (Optional) The allowed file storage for the group, in megabytes. This parameter is
//    ignored if the caller does not have the manage_storage_quotas permission.
// # SISGroupID (Optional) The sis ID of the group. Must have manage_sis permission to set.
//
type CreateGroupGroupCategories struct {
	Path struct {
		GroupCategoryID string `json:"group_category_id" url:"group_category_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		Name           string `json:"name" url:"name,omitempty"`                         //  (Optional)
		Description    string `json:"description" url:"description,omitempty"`           //  (Optional)
		IsPublic       bool   `json:"is_public" url:"is_public,omitempty"`               //  (Optional)
		JoinLevel      string `json:"join_level" url:"join_level,omitempty"`             //  (Optional) . Must be one of parent_context_auto_join, parent_context_request, invitation_only
		StorageQuotaMb int64  `json:"storage_quota_mb" url:"storage_quota_mb,omitempty"` //  (Optional)
		SISGroupID     string `json:"sis_group_id" url:"sis_group_id,omitempty"`         //  (Optional)
	} `json:"form"`
}

func (t *CreateGroupGroupCategories) GetMethod() string {
	return "POST"
}

func (t *CreateGroupGroupCategories) GetURLPath() string {
	path := "group_categories/{group_category_id}/groups"
	path = strings.ReplaceAll(path, "{group_category_id}", fmt.Sprintf("%v", t.Path.GroupCategoryID))
	return path
}

func (t *CreateGroupGroupCategories) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateGroupGroupCategories) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *CreateGroupGroupCategories) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *CreateGroupGroupCategories) HasErrors() error {
	errs := []string{}
	if t.Path.GroupCategoryID == "" {
		errs = append(errs, "'GroupCategoryID' is required")
	}
	if t.Form.JoinLevel != "" && !string_utils.Include([]string{"parent_context_auto_join", "parent_context_request", "invitation_only"}, t.Form.JoinLevel) {
		errs = append(errs, "JoinLevel must be one of parent_context_auto_join, parent_context_request, invitation_only")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateGroupGroupCategories) Do(c *canvasapi.Canvas) (*models.Group, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.Group{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
