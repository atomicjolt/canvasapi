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

// CreatePageGroups Create a new wiki page
// https://canvas.instructure.com/doc/api/pages.html
//
// Path Parameters:
// # GroupID (Required) ID
//
// Form Parameters:
// # WikiPage (Required) The title for the new page.
// # WikiPage (Optional) The content for the new page.
// # WikiPage (Optional) . Must be one of teachers, students, members, publicWhich user roles are allowed to edit this page. Any combination
//    of these roles is allowed (separated by commas).
//
//    "teachers":: Allows editing by teachers in the course.
//    "students":: Allows editing by students in the course.
//    "members":: For group wikis, allows editing by members of the group.
//    "public":: Allows editing by any user.
// # WikiPage (Optional) Whether participants should be notified when this page changes.
// # WikiPage (Optional) Whether the page is published (true) or draft state (false).
// # WikiPage (Optional) Set an unhidden page as the front page (if true)
//
type CreatePageGroups struct {
	Path struct {
		GroupID string `json:"group_id" url:"group_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		WikiPage struct {
			Title          string `json:"title" url:"title,omitempty"`                       //  (Required)
			Body           string `json:"body" url:"body,omitempty"`                         //  (Optional)
			EditingRoles   string `json:"editing_roles" url:"editing_roles,omitempty"`       //  (Optional) . Must be one of teachers, students, members, public
			NotifyOfUpdate bool   `json:"notify_of_update" url:"notify_of_update,omitempty"` //  (Optional)
			Published      bool   `json:"published" url:"published,omitempty"`               //  (Optional)
			FrontPage      bool   `json:"front_page" url:"front_page,omitempty"`             //  (Optional)
		} `json:"wiki_page" url:"wiki_page,omitempty"`
	} `json:"form"`
}

func (t *CreatePageGroups) GetMethod() string {
	return "POST"
}

func (t *CreatePageGroups) GetURLPath() string {
	path := "groups/{group_id}/pages"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	return path
}

func (t *CreatePageGroups) GetQuery() (string, error) {
	return "", nil
}

func (t *CreatePageGroups) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *CreatePageGroups) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *CreatePageGroups) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'GroupID' is required")
	}
	if t.Form.WikiPage.Title == "" {
		errs = append(errs, "'WikiPage' is required")
	}
	if t.Form.WikiPage.EditingRoles != "" && !string_utils.Include([]string{"teachers", "students", "members", "public"}, t.Form.WikiPage.EditingRoles) {
		errs = append(errs, "WikiPage must be one of teachers, students, members, public")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreatePageGroups) Do(c *canvasapi.Canvas) (*models.Page, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.Page{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
