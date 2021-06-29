package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/string_utils"
)

// GetOutcomeResultRollups Gets the outcome rollups for the users and outcomes in the specified
// context.
// https://canvas.instructure.com/doc/api/outcome_results.html
//
// Path Parameters:
// # CourseID (Required) ID
//
// Query Parameters:
// # Aggregate (Optional) . Must be one of courseIf specified, instead of returning one rollup for each user, all the user
//    rollups will be combined into one rollup for the course that will contain
//    the average (or median, see below) rollup score for each outcome.
// # AggregateStat (Optional) . Must be one of mean, medianIf aggregate rollups requested, then this value determines what
//    statistic is used for the aggregate. Defaults to "mean" if this value
//    is not specified.
// # UserIDs (Optional) If specified, only the users whose ids are given will be included in the
//    results or used in an aggregate result. it is an error to specify an id
//    for a user who is not a student in the context
// # OutcomeIDs (Optional) If specified, only the outcomes whose ids are given will be included in the
//    results. it is an error to specify an id for an outcome which is not linked
//    to the context.
// # Include (Optional) [String, "courses"|"outcomes"|"outcomes.alignments"|"outcome_groups"|"outcome_links"|"outcome_paths"|"users"]
//    Specify additional collections to be side loaded with the result.
// # Exclude (Optional) . Must be one of missing_user_rollupsSpecify additional values to exclude. "missing_user_rollups" excludes
//    rollups for users without results.
// # SortBy (Optional) . Must be one of student, outcomeIf specified, sorts outcome result rollups. "student" sorting will sort
//    by a user's sortable name. "outcome" sorting will sort by the given outcome's
//    rollup score. The latter requires specifying the "sort_outcome_id" parameter.
//    By default, the sort order is ascending.
// # SortOutcomeID (Optional) If outcome sorting requested, then this determines which outcome to use
//    for rollup score sorting.
// # SortOrder (Optional) . Must be one of asc, descIf sorting requested, then this allows changing the default sort order of
//    ascending to descending.
//
type GetOutcomeResultRollups struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		Aggregate     string   `json:"aggregate" url:"aggregate,omitempty"`             //  (Optional) . Must be one of course
		AggregateStat string   `json:"aggregate_stat" url:"aggregate_stat,omitempty"`   //  (Optional) . Must be one of mean, median
		UserIDs       []int64  `json:"user_ids" url:"user_ids,omitempty"`               //  (Optional)
		OutcomeIDs    []int64  `json:"outcome_ids" url:"outcome_ids,omitempty"`         //  (Optional)
		Include       []string `json:"include" url:"include,omitempty"`                 //  (Optional)
		Exclude       []string `json:"exclude" url:"exclude,omitempty"`                 //  (Optional) . Must be one of missing_user_rollups
		SortBy        string   `json:"sort_by" url:"sort_by,omitempty"`                 //  (Optional) . Must be one of student, outcome
		SortOutcomeID int64    `json:"sort_outcome_id" url:"sort_outcome_id,omitempty"` //  (Optional)
		SortOrder     string   `json:"sort_order" url:"sort_order,omitempty"`           //  (Optional) . Must be one of asc, desc
	} `json:"query"`
}

func (t *GetOutcomeResultRollups) GetMethod() string {
	return "GET"
}

func (t *GetOutcomeResultRollups) GetURLPath() string {
	path := "courses/{course_id}/outcome_rollups"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *GetOutcomeResultRollups) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *GetOutcomeResultRollups) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetOutcomeResultRollups) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetOutcomeResultRollups) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Query.Aggregate != "" && !string_utils.Include([]string{"course"}, t.Query.Aggregate) {
		errs = append(errs, "Aggregate must be one of course")
	}
	if t.Query.AggregateStat != "" && !string_utils.Include([]string{"mean", "median"}, t.Query.AggregateStat) {
		errs = append(errs, "AggregateStat must be one of mean, median")
	}
	for _, v := range t.Query.Exclude {
		if v != "" && !string_utils.Include([]string{"missing_user_rollups"}, v) {
			errs = append(errs, "Exclude must be one of missing_user_rollups")
		}
	}
	if t.Query.SortBy != "" && !string_utils.Include([]string{"student", "outcome"}, t.Query.SortBy) {
		errs = append(errs, "SortBy must be one of student, outcome")
	}
	if t.Query.SortOrder != "" && !string_utils.Include([]string{"asc", "desc"}, t.Query.SortOrder) {
		errs = append(errs, "SortOrder must be one of asc, desc")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetOutcomeResultRollups) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
