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

// UpdateCreateFrontPageCourses Update the title or contents of the front page
// https://canvas.instructure.com/doc/api/pages.html
//
// Path Parameters:
// # CourseID (Required) ID
//
// Form Parameters:
// # WikiPage (Optional) The title for the new page. NOTE: changing a page's title will change its
//    url. The updated url will be returned in the result.
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
//
type UpdateCreateFrontPageCourses struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
	} `json:"path"`

	Form struct {
		WikiPage struct {
			Title          string `json:"title"`            //  (Optional)
			Body           string `json:"body"`             //  (Optional)
			EditingRoles   string `json:"editing_roles"`    //  (Optional) . Must be one of teachers, students, members, public
			NotifyOfUpdate bool   `json:"notify_of_update"` //  (Optional)
			Published      bool   `json:"published"`        //  (Optional)
		} `json:"wiki_page"`
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

func (t *UpdateCreateFrontPageCourses) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *UpdateCreateFrontPageCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if !string_utils.Include([]string{"teachers", "students", "members", "public"}, t.Form.WikiPage.EditingRoles) {
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
