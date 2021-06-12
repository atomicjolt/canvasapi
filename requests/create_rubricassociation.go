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

// CreateRubricassociation Returns the rubric with the given id.
// https://canvas.instructure.com/doc/api/rubrics.html
//
// Path Parameters:
// # CourseID (Required) ID
//
// Form Parameters:
// # RubricAssociation (Optional) The id of the Rubric
// # RubricAssociation (Optional) The id of the object with which this rubric is associated
// # RubricAssociation (Optional) . Must be one of Assignment, Course, AccountThe type of object this rubric is associated with
// # RubricAssociation (Optional) The name of the object this rubric is associated with
// # RubricAssociation (Optional) Whether or not the associated rubric is used for grade calculation
// # RubricAssociation (Optional) Whether or not the score total is displayed within the rubric.
//    This option is only available if the rubric is not used for grading.
// # RubricAssociation (Optional) . Must be one of grading, bookmarkWhether or not the association is for grading (and thus linked to an assignment)
//    or if it's to indicate the rubric should appear in its context
// # RubricAssociation (Optional) Whether or not the associated rubric appears in its context
//
type CreateRubricassociation struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
	} `json:"path"`

	Form struct {
		RubricAssociation struct {
			RubricID        int64  `json:"rubric_id"`        //  (Optional)
			AssociationID   int64  `json:"association_id"`   //  (Optional)
			AssociationType string `json:"association_type"` //  (Optional) . Must be one of Assignment, Course, Account
			Title           string `json:"title"`            //  (Optional)
			UseForGrading   bool   `json:"use_for_grading"`  //  (Optional)
			HideScoreTotal  bool   `json:"hide_score_total"` //  (Optional)
			Purpose         string `json:"purpose"`          //  (Optional) . Must be one of grading, bookmark
			Bookmarked      bool   `json:"bookmarked"`       //  (Optional)
		} `json:"rubric_association"`
	} `json:"form"`
}

func (t *CreateRubricassociation) GetMethod() string {
	return "POST"
}

func (t *CreateRubricassociation) GetURLPath() string {
	path := "courses/{course_id}/rubric_associations"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *CreateRubricassociation) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateRubricassociation) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *CreateRubricassociation) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if !string_utils.Include([]string{"Assignment", "Course", "Account"}, t.Form.RubricAssociation.AssociationType) {
		errs = append(errs, "RubricAssociation must be one of Assignment, Course, Account")
	}
	if !string_utils.Include([]string{"grading", "bookmark"}, t.Form.RubricAssociation.Purpose) {
		errs = append(errs, "RubricAssociation must be one of grading, bookmark")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateRubricassociation) Do(c *canvasapi.Canvas) (*models.RubricAssociation, error) {
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
