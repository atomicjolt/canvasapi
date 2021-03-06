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

// AssignUnassignedMembers Assign all unassigned members as evenly as possible among the existing
// student groups.
// https://canvas.instructure.com/doc/api/group_categories.html
//
// Path Parameters:
// # Path.GroupCategoryID (Required) ID
//
// Form Parameters:
// # Form.Sync (Optional) The assigning is done asynchronously by default. If you would like to
//    override this and have the assigning done synchronously, set this value
//    to true.
//
type AssignUnassignedMembers struct {
	Path struct {
		GroupCategoryID string `json:"group_category_id" url:"group_category_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		Sync bool `json:"sync" url:"sync,omitempty"` //  (Optional)
	} `json:"form"`
}

func (t *AssignUnassignedMembers) GetMethod() string {
	return "POST"
}

func (t *AssignUnassignedMembers) GetURLPath() string {
	path := "group_categories/{group_category_id}/assign_unassigned_members"
	path = strings.ReplaceAll(path, "{group_category_id}", fmt.Sprintf("%v", t.Path.GroupCategoryID))
	return path
}

func (t *AssignUnassignedMembers) GetQuery() (string, error) {
	return "", nil
}

func (t *AssignUnassignedMembers) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *AssignUnassignedMembers) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *AssignUnassignedMembers) HasErrors() error {
	errs := []string{}
	if t.Path.GroupCategoryID == "" {
		errs = append(errs, "'Path.GroupCategoryID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *AssignUnassignedMembers) Do(c *canvasapi.Canvas) (*models.GroupMembership, *models.Progress, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, nil, err
	}
	groupMembership := models.GroupMembership{}
	progress := models.Progress{}
	if t.Form.Sync {
		err = json.Unmarshal(body, &groupMembership)
		if err != nil {
			return nil, nil, err
		}
	} else {
		err = json.Unmarshal(body, &progress)
		if err != nil {
			return nil, nil, err
		}
	}

	return &groupMembership, &progress, nil
}
