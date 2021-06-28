package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// GetOutcomeResults Gets the outcome results for users and outcomes in the specified context.
// https://canvas.instructure.com/doc/api/outcome_results.html
//
// Path Parameters:
// # CourseID (Required) ID
//
// Query Parameters:
// # UserIDs (Optional) If specified, only the users whose ids are given will be included in the
//    results. SIS ids can be used, prefixed by "sis_user_id:".
//    It is an error to specify an id for a user who is not a student in
//    the context.
// # OutcomeIDs (Optional) If specified, only the outcomes whose ids are given will be included in the
//    results. it is an error to specify an id for an outcome which is not linked
//    to the context.
// # Include (Optional) [String, "alignments"|"outcomes"|"outcomes.alignments"|"outcome_groups"|"outcome_links"|"outcome_paths"|"users"]
//    Specify additional collections to be side loaded with the result.
//    "alignments" includes only the alignments referenced by the returned
//    results.
//    "outcomes.alignments" includes all alignments referenced by outcomes in the
//    context.
// # IncludeHidden (Optional) If true, results that are hidden from the learning mastery gradebook and student rollup
//    scores will be included
//
type GetOutcomeResults struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		UserIDs       []int64  `json:"user_ids" url:"user_ids,omitempty"`             //  (Optional)
		OutcomeIDs    []int64  `json:"outcome_ids" url:"outcome_ids,omitempty"`       //  (Optional)
		Include       []string `json:"include" url:"include,omitempty"`               //  (Optional)
		IncludeHidden bool     `json:"include_hidden" url:"include_hidden,omitempty"` //  (Optional)
	} `json:"query"`
}

func (t *GetOutcomeResults) GetMethod() string {
	return "GET"
}

func (t *GetOutcomeResults) GetURLPath() string {
	path := "courses/{course_id}/outcome_results"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *GetOutcomeResults) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *GetOutcomeResults) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetOutcomeResults) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetOutcomeResults) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetOutcomeResults) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
