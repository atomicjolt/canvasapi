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

// ListUncollatedSubmissionVersions Gives a paginated, uncollated list of submission versions for all matching
// submissions in the context. This SubmissionVersion objects will not include
// the +new_grade+ or +previous_grade+ keys, only the +grade+; same for
// +graded_at+ and +grader+.
// https://canvas.instructure.com/doc/api/gradebook_history.html
//
// Path Parameters:
// # CourseID (Required) The id of the contextual course for this API call
//
// Query Parameters:
// # AssignmentID (Optional) The ID of the assignment for which you want to see submissions. If
//    absent, versions of submissions from any assignment in the course are
//    included.
// # UserID (Optional) The ID of the user for which you want to see submissions. If absent,
//    versions of submissions from any user in the course are included.
// # Ascending (Optional) Returns submission versions in ascending date order (oldest first). If
//    absent, returns submission versions in descending date order (newest
//    first).
//
type ListUncollatedSubmissionVersions struct {
	Path struct {
		CourseID int64 `json:"course_id"` //  (Required)
	} `json:"path"`

	Query struct {
		AssignmentID int64 `json:"assignment_id"` //  (Optional)
		UserID       int64 `json:"user_id"`       //  (Optional)
		Ascending    bool  `json:"ascending"`     //  (Optional)
	} `json:"query"`
}

func (t *ListUncollatedSubmissionVersions) GetMethod() string {
	return "GET"
}

func (t *ListUncollatedSubmissionVersions) GetURLPath() string {
	path := "courses/{course_id}/gradebook_history/feed"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *ListUncollatedSubmissionVersions) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *ListUncollatedSubmissionVersions) GetBody() (string, error) {
	return "", nil
}

func (t *ListUncollatedSubmissionVersions) HasErrors() error {
	return nil
}

func (t *ListUncollatedSubmissionVersions) Do(c *canvasapi.Canvas) ([]*models.SubmissionVersion, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.SubmissionVersion{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}