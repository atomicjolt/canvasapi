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

// SetUsageRightsUsers Sets copyright and license information for one or more files
// https://canvas.instructure.com/doc/api/files.html
//
// Path Parameters:
// # UserID (Required) ID
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
type SetUsageRightsUsers struct {
	Path struct {
		UserID string `json:"user_id"` //  (Required)
	} `json:"path"`

	Form struct {
		FileIDs     []string `json:"file_ids"`   //  (Required)
		FolderIDs   []string `json:"folder_ids"` //  (Optional)
		Publish     bool     `json:"publish"`    //  (Optional)
		UsageRights struct {
			UseJustification string `json:"use_justification"` //  (Required) . Must be one of own_copyright, used_by_permission, fair_use, public_domain, creative_commons
			LegalCopyright   string `json:"legal_copyright"`   //  (Optional)
			License          string `json:"license"`           //  (Optional)
		} `json:"usage_rights"`
	} `json:"form"`
}

func (t *SetUsageRightsUsers) GetMethod() string {
	return "PUT"
}

func (t *SetUsageRightsUsers) GetURLPath() string {
	path := "users/{user_id}/usage_rights"
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	return path
}

func (t *SetUsageRightsUsers) GetQuery() (string, error) {
	return "", nil
}

func (t *SetUsageRightsUsers) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *SetUsageRightsUsers) HasErrors() error {
	errs := []string{}
	if t.Path.UserID == "" {
		errs = append(errs, "'UserID' is required")
	}
	if t.Form.FileIDs == nil {
		errs = append(errs, "'FileIDs' is required")
	}
	if t.Form.UsageRights.UseJustification == "" {
		errs = append(errs, "'UsageRights' is required")
	}
	if !string_utils.Include([]string{"own_copyright", "used_by_permission", "fair_use", "public_domain", "creative_commons"}, t.Form.UsageRights.UseJustification) {
		errs = append(errs, "UsageRights must be one of own_copyright, used_by_permission, fair_use, public_domain, creative_commons")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *SetUsageRightsUsers) Do(c *canvasapi.Canvas) (*models.UsageRights, error) {
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
