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

// UpdateCreateFrontPageCourses Update the title or contents of the front page
// https://canvas.instructure.com/doc/api/pages.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
//
// Form Parameters:
// # Form.WikiPage.Title (Optional) The title for the new page. NOTE: changing a page's title will change its
//    url. The updated url will be returned in the result.
// # Form.WikiPage.Body (Optional) The content for the new page.
// # Form.WikiPage.EditingRoles (Optional) . Must be one of teachers, students, members, publicWhich user roles are allowed to edit this page. Any combination
//    of these roles is allowed (separated by commas).
//
//    "teachers":: Allows editing by teachers in the course.
//    "students":: Allows editing by students in the course.
//    "members":: For group wikis, allows editing by members of the group.
//    "public":: Allows editing by any user.
// # Form.WikiPage.NotifyOfUpdate (Optional) Whether participants should be notified when this page changes.
// # Form.WikiPage.Published (Optional) Whether the page is published (true) or draft state (false).
//
type UpdateCreateFrontPageCourses struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		WikiPage struct {
			Title          string `json:"title" url:"title,omitempty"`                       //  (Optional)
			Body           string `json:"body" url:"body,omitempty"`                         //  (Optional)
			EditingRoles   string `json:"editing_roles" url:"editing_roles,omitempty"`       //  (Optional) . Must be one of teachers, students, members, public
			NotifyOfUpdate bool   `json:"notify_of_update" url:"notify_of_update,omitempty"` //  (Optional)
			Published      bool   `json:"published" url:"published,omitempty"`               //  (Optional)
		} `json:"wiki_page" url:"wiki_page,omitempty"`
	} `json:"form"`
}

func (t *UpdateCreateFrontPageCourses) GetMethod() string {
	return "PUT"
}

func (t *UpdateCreateFrontPageCourses) GetURLPath() string {
	path := "courses/{course_id}/front_page"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *UpdateCreateFrontPageCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdateCreateFrontPageCourses) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *UpdateCreateFrontPageCourses) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *UpdateCreateFrontPageCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Form.WikiPage.EditingRoles != "" && !string_utils.Include([]string{"teachers", "students", "members", "public"}, t.Form.WikiPage.EditingRoles) {
		errs = append(errs, "WikiPage must be one of teachers, students, members, public")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UpdateCreateFrontPageCourses) Do(c *canvasapi.Canvas) (*models.Page, error) {
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
