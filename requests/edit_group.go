package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
// # GroupID (Required) ID
//
// Form Parameters:
// # Name (Optional) The name of the group
// # Description (Optional) A description of the group
// # IsPublic (Optional) Whether the group is public (applies only to community groups). Currently
//    you cannot set a group back to private once it has been made public.
// # JoinLevel (Optional) . Must be one of parent_context_auto_join, parent_context_request, invitation_onlyno description
// # AvatarID (Optional) The id of the attachment previously uploaded to the group that you would
//    like to use as the avatar image for this group.
// # StorageQuotaMb (Optional) The allowed file storage for the group, in megabytes. This parameter is
//    ignored if the caller does not have the manage_storage_quotas permission.
// # Members (Optional) An array of user ids for users you would like in the group.
//    Users not in the group will be sent invitations. Existing group
//    members who aren't in the list will be removed from the group.
// # SISGroupID (Optional) The sis ID of the group. Must have manage_sis permission to set.
//
type EditGroup struct {
	Path struct {
		GroupID string `json:"group_id"` //  (Required)
	} `json:"path"`

	Form struct {
		Name           string   `json:"name"`             //  (Optional)
		Description    string   `json:"description"`      //  (Optional)
		IsPublic       bool     `json:"is_public"`        //  (Optional)
		JoinLevel      string   `json:"join_level"`       //  (Optional) . Must be one of parent_context_auto_join, parent_context_request, invitation_only
		AvatarID       int64    `json:"avatar_id"`        //  (Optional)
		StorageQuotaMb int64    `json:"storage_quota_mb"` //  (Optional)
		Members        []string `json:"members"`          //  (Optional)
		SISGroupID     string   `json:"sis_group_id"`     //  (Optional)
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

func (t *EditGroup) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *EditGroup) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'GroupID' is required")
	}
	if !string_utils.Include([]string{"parent_context_auto_join", "parent_context_request", "invitation_only"}, t.Form.JoinLevel) {
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