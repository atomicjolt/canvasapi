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

// EditGroup Modifies an existing group.  Note that to set an avatar image for the
// group, you must first upload the image file to the group, and the use the
// id in the response as the argument to this function.  See the
// {file:file_uploads.html File Upload Documentation} for details on the file
// upload workflow.
// https://canvas.instructure.com/doc/api/groups.html
//
// Path Parameters:
// # Path.GroupID (Required) ID
//
// Form Parameters:
// # Form.Name (Optional) The name of the group
// # Form.Description (Optional) A description of the group
// # Form.IsPublic (Optional) Whether the group is public (applies only to community groups). Currently
//    you cannot set a group back to private once it has been made public.
// # Form.JoinLevel (Optional) . Must be one of parent_context_auto_join, parent_context_request, invitation_onlyno description
// # Form.AvatarID (Optional) The id of the attachment previously uploaded to the group that you would
//    like to use as the avatar image for this group.
// # Form.StorageQuotaMb (Optional) The allowed file storage for the group, in megabytes. This parameter is
//    ignored if the caller does not have the manage_storage_quotas permission.
// # Form.Members (Optional) An array of user ids for users you would like in the group.
//    Users not in the group will be sent invitations. Existing group
//    members who aren't in the list will be removed from the group.
// # Form.SISGroupID (Optional) The sis ID of the group. Must have manage_sis permission to set.
//
type EditGroup struct {
	Path struct {
		GroupID string `json:"group_id" url:"group_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		Name           string   `json:"name" url:"name,omitempty"`                         //  (Optional)
		Description    string   `json:"description" url:"description,omitempty"`           //  (Optional)
		IsPublic       bool     `json:"is_public" url:"is_public,omitempty"`               //  (Optional)
		JoinLevel      string   `json:"join_level" url:"join_level,omitempty"`             //  (Optional) . Must be one of parent_context_auto_join, parent_context_request, invitation_only
		AvatarID       int64    `json:"avatar_id" url:"avatar_id,omitempty"`               //  (Optional)
		StorageQuotaMb int64    `json:"storage_quota_mb" url:"storage_quota_mb,omitempty"` //  (Optional)
		Members        []string `json:"members" url:"members,omitempty"`                   //  (Optional)
		SISGroupID     string   `json:"sis_group_id" url:"sis_group_id,omitempty"`         //  (Optional)
	} `json:"form"`
}

func (t *EditGroup) GetMethod() string {
	return "PUT"
}

func (t *EditGroup) GetURLPath() string {
	path := "groups/{group_id}"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	return path
}

func (t *EditGroup) GetQuery() (string, error) {
	return "", nil
}

func (t *EditGroup) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *EditGroup) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *EditGroup) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'Path.GroupID' is required")
	}
	if t.Form.JoinLevel != "" && !string_utils.Include([]string{"parent_context_auto_join", "parent_context_request", "invitation_only"}, t.Form.JoinLevel) {
		errs = append(errs, "JoinLevel must be one of parent_context_auto_join, parent_context_request, invitation_only")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *EditGroup) Do(c *canvasapi.Canvas) (*models.Group, error) {
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
