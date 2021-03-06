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

// UpdateRubricassociation Returns the rubric with the given id.
// https://canvas.instructure.com/doc/api/rubrics.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
// # Path.ID (Required) The id of the RubricAssociation to update
//
// Form Parameters:
// # Form.RubricAssociation.RubricID (Optional) The id of the Rubric
// # Form.RubricAssociation.AssociationID (Optional) The id of the object with which this rubric is associated
// # Form.RubricAssociation.AssociationType (Optional) . Must be one of Assignment, Course, AccountThe type of object this rubric is associated with
// # Form.RubricAssociation.Title (Optional) The name of the object this rubric is associated with
// # Form.RubricAssociation.UseForGrading (Optional) Whether or not the associated rubric is used for grade calculation
// # Form.RubricAssociation.HideScoreTotal (Optional) Whether or not the score total is displayed within the rubric.
//    This option is only available if the rubric is not used for grading.
// # Form.RubricAssociation.Purpose (Optional) . Must be one of grading, bookmarkWhether or not the association is for grading (and thus linked to an assignment)
//    or if it's to indicate the rubric should appear in its context
// # Form.RubricAssociation.Bookmarked (Optional) Whether or not the associated rubric appears in its context
//
type UpdateRubricassociation struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		ID       int64  `json:"id" url:"id,omitempty"`               //  (Required)
	} `json:"path"`

	Form struct {
		RubricAssociation struct {
			RubricID        int64  `json:"rubric_id" url:"rubric_id,omitempty"`               //  (Optional)
			AssociationID   int64  `json:"association_id" url:"association_id,omitempty"`     //  (Optional)
			AssociationType string `json:"association_type" url:"association_type,omitempty"` //  (Optional) . Must be one of Assignment, Course, Account
			Title           string `json:"title" url:"title,omitempty"`                       //  (Optional)
			UseForGrading   bool   `json:"use_for_grading" url:"use_for_grading,omitempty"`   //  (Optional)
			HideScoreTotal  bool   `json:"hide_score_total" url:"hide_score_total,omitempty"` //  (Optional)
			Purpose         string `json:"purpose" url:"purpose,omitempty"`                   //  (Optional) . Must be one of grading, bookmark
			Bookmarked      bool   `json:"bookmarked" url:"bookmarked,omitempty"`             //  (Optional)
		} `json:"rubric_association" url:"rubric_association,omitempty"`
	} `json:"form"`
}

func (t *UpdateRubricassociation) GetMethod() string {
	return "PUT"
}

func (t *UpdateRubricassociation) GetURLPath() string {
	path := "courses/{course_id}/rubric_associations/{id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *UpdateRubricassociation) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdateRubricassociation) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *UpdateRubricassociation) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *UpdateRubricassociation) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Form.RubricAssociation.AssociationType != "" && !string_utils.Include([]string{"Assignment", "Course", "Account"}, t.Form.RubricAssociation.AssociationType) {
		errs = append(errs, "RubricAssociation must be one of Assignment, Course, Account")
	}
	if t.Form.RubricAssociation.Purpose != "" && !string_utils.Include([]string{"grading", "bookmark"}, t.Form.RubricAssociation.Purpose) {
		errs = append(errs, "RubricAssociation must be one of grading, bookmark")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UpdateRubricassociation) Do(c *canvasapi.Canvas) (*models.RubricAssociation, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.RubricAssociation{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
