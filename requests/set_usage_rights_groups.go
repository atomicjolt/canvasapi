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

// SetUsageRightsGroups Sets copyright and license information for one or more files
// https://canvas.instructure.com/doc/api/files.html
//
// Path Parameters:
// # GroupID (Required) ID
//
// Form Parameters:
// # FileIDs (Required) List of ids of files to set usage rights for.
// # FolderIDs (Optional) List of ids of folders to search for files to set usage rights for.
//    Note that new files uploaded to these folders do not automatically inherit these rights.
// # Publish (Optional) Whether the file(s) or folder(s) should be published on save, provided that usage rights have been specified (set to `true` to publish on save).
// # UsageRights (Required) . Must be one of own_copyright, used_by_permission, fair_use, public_domain, creative_commonsThe intellectual property justification for using the files in Canvas
// # UsageRights (Optional) The legal copyright line for the files
// # UsageRights (Optional) The license that applies to the files. See the {api:UsageRightsController#licenses List licenses endpoint} for the supported license types.
//
type SetUsageRightsGroups struct {
	Path struct {
		GroupID string `json:"group_id" url:"group_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		FileIDs     []string `json:"file_ids" url:"file_ids,omitempty"`     //  (Required)
		FolderIDs   []string `json:"folder_ids" url:"folder_ids,omitempty"` //  (Optional)
		Publish     bool     `json:"publish" url:"publish,omitempty"`       //  (Optional)
		UsageRights struct {
			UseJustification string `json:"use_justification" url:"use_justification,omitempty"` //  (Required) . Must be one of own_copyright, used_by_permission, fair_use, public_domain, creative_commons
			LegalCopyright   string `json:"legal_copyright" url:"legal_copyright,omitempty"`     //  (Optional)
			License          string `json:"license" url:"license,omitempty"`                     //  (Optional)
		} `json:"usage_rights" url:"usage_rights,omitempty"`
	} `json:"form"`
}

func (t *SetUsageRightsGroups) GetMethod() string {
	return "PUT"
}

func (t *SetUsageRightsGroups) GetURLPath() string {
	path := "groups/{group_id}/usage_rights"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	return path
}

func (t *SetUsageRightsGroups) GetQuery() (string, error) {
	return "", nil
}

func (t *SetUsageRightsGroups) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *SetUsageRightsGroups) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *SetUsageRightsGroups) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'GroupID' is required")
	}
	if t.Form.FileIDs == nil {
		errs = append(errs, "'FileIDs' is required")
	}
	if t.Form.UsageRights.UseJustification == "" {
		errs = append(errs, "'UsageRights' is required")
	}
	if t.Form.UsageRights.UseJustification != "" && !string_utils.Include([]string{"own_copyright", "used_by_permission", "fair_use", "public_domain", "creative_commons"}, t.Form.UsageRights.UseJustification) {
		errs = append(errs, "UsageRights must be one of own_copyright, used_by_permission, fair_use, public_domain, creative_commons")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *SetUsageRightsGroups) Do(c *canvasapi.Canvas) (*models.UsageRights, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.UsageRights{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
