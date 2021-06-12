package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// NamesAndRoleListGroupMemberships Return active NamesAndRoleMemberships in the given group.
// https://canvas.instructure.com/doc/api/names_and_role.html
//
// Path Parameters:
// # GroupID (Required) ID
//
// Query Parameters:
// # RLID (Optional) If specified only NamesAndRoleMemberships with access to the LTI link references by this `rlid` will be included.
//    Also causes the member array to be included for each returned NamesAndRoleMembership.
//    If the role parameter is also present, it will be 'and-ed' together with this parameter
// # Role (Optional) If specified only NamesAndRoleMemberships having this role in the given Group will be included.
//    Value must be a fully-qualified LTI/LIS role URN. Further, only
//    http://purl.imsglobal.org/vocab/lis/v2/membership#Member and
//    http://purl.imsglobal.org/vocab/lis/v2/membership#Manager are supported.
//    If the `rlid` parameter is also present, it will be 'and-ed' together with this parameter
// # Limit (Optional) May be used to limit the number of NamesAndRoleMemberships returned in a page. Defaults to 50.
//
type NamesAndRoleListGroupMemberships struct {
	Path struct {
		GroupID string `json:"group_id"` //  (Required)
	} `json:"path"`

	Query struct {
		RLID  string `json:"rlid"`  //  (Optional)
		Role  string `json:"role"`  //  (Optional)
		Limit string `json:"limit"` //  (Optional)
	} `json:"query"`
}

func (t *NamesAndRoleListGroupMemberships) GetMethod() string {
	return "GET"
}

func (t *NamesAndRoleListGroupMemberships) GetURLPath() string {
	path := "/lti/groups/{group_id}/names_and_roles"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	return path
}

func (t *NamesAndRoleListGroupMemberships) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *NamesAndRoleListGroupMemberships) GetBody() (string, error) {
	return "", nil
}

func (t *NamesAndRoleListGroupMemberships) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'GroupID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *NamesAndRoleListGroupMemberships) Do(c *canvasapi.Canvas) (*models.NamesAndRoleMemberships, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.NamesAndRoleMemberships{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
