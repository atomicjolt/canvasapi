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

// CreateGroupGroups Creates a new group. Groups created using the "/api/v1/groups/"
// endpoint will be community groups.
// https://canvas.instructure.com/doc/api/groups.html
//
// Form Parameters:
// # Form.Name (Optional) The name of the group
// # Form.Description (Optional) A description of the group
// # Form.IsPublic (Optional) whether the group is public (applies only to community groups)
// # Form.JoinLevel (Optional) . Must be one of parent_context_auto_join, parent_context_request, invitation_onlyno description
// # Form.StorageQuotaMb (Optional) The allowed file storage for the group, in megabytes. This parameter is
//    ignored if the caller does not have the manage_storage_quotas permission.
// # Form.SISGroupID (Optional) The sis ID of the group. Must have manage_sis permission to set.
//
type CreateGroupGroups struct {
	Form struct {
		Name           string `json:"name" url:"name,omitempty"`                         //  (Optional)
		Description    string `json:"description" url:"description,omitempty"`           //  (Optional)
		IsPublic       bool   `json:"is_public" url:"is_public,omitempty"`               //  (Optional)
		JoinLevel      string `json:"join_level" url:"join_level,omitempty"`             //  (Optional) . Must be one of parent_context_auto_join, parent_context_request, invitation_only
		StorageQuotaMb int64  `json:"storage_quota_mb" url:"storage_quota_mb,omitempty"` //  (Optional)
		SISGroupID     string `json:"sis_group_id" url:"sis_group_id,omitempty"`         //  (Optional)
	} `json:"form"`
}

func (t *CreateGroupGroups) GetMethod() string {
	return "POST"
}

func (t *CreateGroupGroups) GetURLPath() string {
	return ""
}

func (t *CreateGroupGroups) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateGroupGroups) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *CreateGroupGroups) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *CreateGroupGroups) HasErrors() error {
	errs := []string{}
	if t.Form.JoinLevel != "" && !string_utils.Include([]string{"parent_context_auto_join", "parent_context_request", "invitation_only"}, t.Form.JoinLevel) {
		errs = append(errs, "JoinLevel must be one of parent_context_auto_join, parent_context_request, invitation_only")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateGroupGroups) Do(c *canvasapi.Canvas) (*models.Group, error) {
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
