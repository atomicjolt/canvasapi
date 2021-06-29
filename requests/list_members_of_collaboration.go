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

// ListMembersOfCollaboration A paginated list of the collaborators of a given collaboration
// https://canvas.instructure.com/doc/api/collaborations.html
//
// Path Parameters:
// # Path.ID (Required) ID
//
// Query Parameters:
// # Query.Include (Optional) . Must be one of collaborator_lti_id, avatar_image_url- "collaborator_lti_id": Optional information to include with each member.
//      Represents an identifier to be used for the member in an LTI context.
//    - "avatar_image_url": Optional information to include with each member.
//      The url for the avatar of a collaborator with type 'user'.
//
type ListMembersOfCollaboration struct {
	Path struct {
		ID string `json:"id" url:"id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		Include []string `json:"include" url:"include,omitempty"` //  (Optional) . Must be one of collaborator_lti_id, avatar_image_url
	} `json:"query"`
}

func (t *ListMembersOfCollaboration) GetMethod() string {
	return "GET"
}

func (t *ListMembersOfCollaboration) GetURLPath() string {
	path := "collaborations/{id}/members"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *ListMembersOfCollaboration) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *ListMembersOfCollaboration) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListMembersOfCollaboration) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListMembersOfCollaboration) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	for _, v := range t.Query.Include {
		if v != "" && !string_utils.Include([]string{"collaborator_lti_id", "avatar_image_url"}, v) {
			errs = append(errs, "Include must be one of collaborator_lti_id, avatar_image_url")
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListMembersOfCollaboration) Do(c *canvasapi.Canvas) ([]*models.Collaborator, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.Collaborator{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
