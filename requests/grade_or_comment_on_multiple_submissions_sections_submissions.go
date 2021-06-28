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
)

// GradeOrCommentOnMultipleSubmissionsSectionsSubmissions Update the grading and comments on multiple student's assignment
// submissions in an asynchronous job.
//
// The user must have permission to manage grades in the appropriate context
// (course or section).
// https://canvas.instructure.com/doc/api/submissions.html
//
// Path Parameters:
// # SectionID (Required) ID
//
// Form Parameters:
// # GradeData (Optional) See documentation for the posted_grade argument in the
//    {api:SubmissionsApiController#update Submissions Update} documentation
// # GradeData (Optional) See documentation for the excuse argument in the
//    {api:SubmissionsApiController#update Submissions Update} documentation
// # GradeData (Optional) See documentation for the rubric_assessment argument in the
//    {api:SubmissionsApiController#update Submissions Update} documentation
// # GradeData (Optional) no description
// # GradeData (Optional) no description
// # GradeData (Optional) no description
// # GradeData (Optional) . Must be one of audio, videono description
// # GradeData (Optional) See documentation for the comment[] arguments in the
//    {api:SubmissionsApiController#update Submissions Update} documentation
// # GradeData (Optional) Specifies which assignment to grade.  This argument is not necessary when
//    using the assignment-specific endpoints.
//
type GradeOrCommentOnMultipleSubmissionsSectionsSubmissions struct {
	Path struct {
		SectionID string `json:"section_id" url:"section_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		GradeData map[string]GradeOrCommentOnMultipleSubmissionsSectionsSubmissionsGradeData
	} `json:"form"`
}

func (t *GradeOrCommentOnMultipleSubmissionsSectionsSubmissions) GetMethod() string {
	return "POST"
}

func (t *GradeOrCommentOnMultipleSubmissionsSectionsSubmissions) GetURLPath() string {
	path := "sections/{section_id}/submissions/update_grades"
	path = strings.ReplaceAll(path, "{section_id}", fmt.Sprintf("%v", t.Path.SectionID))
	return path
}

func (t *GradeOrCommentOnMultipleSubmissionsSectionsSubmissions) GetQuery() (string, error) {
	return "", nil
}

func (t *GradeOrCommentOnMultipleSubmissionsSectionsSubmissions) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *GradeOrCommentOnMultipleSubmissionsSectionsSubmissions) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *GradeOrCommentOnMultipleSubmissionsSectionsSubmissions) HasErrors() error {
	errs := []string{}
	if t.Path.SectionID == "" {
		errs = append(errs, "'SectionID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GradeOrCommentOnMultipleSubmissionsSectionsSubmissions) Do(c *canvasapi.Canvas) (*models.Progress, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.Progress{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}

type GradeOrCommentOnMultipleSubmissionsSectionsSubmissionsGradeData struct {
	PostedGrade      string  `json:"posted_grade" url:"posted_grade,omitempty"`             //  (Optional)
	Excuse           bool    `json:"excuse" url:"excuse,omitempty"`                         //  (Optional)
	RubricAssessment string  `json:"rubric_assessment" url:"rubric_assessment,omitempty"`   //  (Optional)
	TextComment      string  `json:"text_comment" url:"text_comment,omitempty"`             //  (Optional)
	GroupComment     bool    `json:"group_comment" url:"group_comment,omitempty"`           //  (Optional)
	MediaCommentID   string  `json:"media_comment_id" url:"media_comment_id,omitempty"`     //  (Optional)
	MediaCommentType string  `json:"media_comment_type" url:"media_comment_type,omitempty"` //  (Optional) . Must be one of audio, video
	FileIDs          []int64 `json:"file_ids" url:"file_ids,omitempty"`                     //  (Optional)
	AssignmentID     int64   `json:"assignment_id" url:"assignment_id,omitempty"`           //  (Optional)
}
