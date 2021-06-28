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

// CreateLinkOutcomeCourses Link an outcome into the outcome group. The outcome to link can either be
// specified by a PUT to the link URL for a specific outcome (the outcome_id
// in the PUT URLs) or by supplying the information for a new outcome (title,
// description, ratings, mastery_points) in a POST to the collection.
//
// If linking an existing outcome, the outcome_id must identify an outcome
// available to this context; i.e. an outcome owned by this group's context,
// an outcome owned by an associated account, or a global outcome. With
// outcome_id present, any other parameters (except move_from) are ignored.
//
// If defining a new outcome, the outcome is created in the outcome group's
// context using the provided title, description, ratings, and mastery points;
// the title is required but all other fields are optional. The new outcome
// is then linked into the outcome group.
//
// If ratings are provided when creating a new outcome, an embedded rubric
// criterion is included in the new outcome. This criterion's mastery_points
// default to the maximum points in the highest rating if not specified in the
// mastery_points parameter. Any ratings lacking a description are given a
// default of "No description". Any ratings lacking a point value are given a
// default of 0. If no ratings are provided, the mastery_points parameter is
// ignored.
// https://canvas.instructure.com/doc/api/outcome_groups.html
//
// Path Parameters:
// # CourseID (Required) ID
// # ID (Required) ID
//
// Form Parameters:
// # OutcomeID (Optional) The ID of the existing outcome to link.
// # MoveFrom (Optional) The ID of the old outcome group. Only used if outcome_id is present.
// # Title (Optional) The title of the new outcome. Required if outcome_id is absent.
// # DisplayName (Optional) A friendly name shown in reports for outcomes with cryptic titles,
//    such as common core standards names.
// # Description (Optional) The description of the new outcome.
// # VendorGuid (Optional) A custom GUID for the learning standard.
// # MasteryPoints (Optional) The mastery threshold for the embedded rubric criterion.
// # Ratings (Optional) The description of a rating level for the embedded rubric criterion.
// # Ratings (Optional) The points corresponding to a rating level for the embedded rubric criterion.
// # CalculationMethod (Optional) . Must be one of decaying_average, n_mastery, latest, highestThe new calculation method.  Defaults to "decaying_average"
// # CalculationInt (Optional) The new calculation int.  Only applies if the calculation_method is "decaying_average" or "n_mastery". Defaults to 65
//
type CreateLinkOutcomeCourses struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		ID       string `json:"id" url:"id,omitempty"`               //  (Required)
	} `json:"path"`

	Form struct {
		OutcomeID     int64  `json:"outcome_id" url:"outcome_id,omitempty"`         //  (Optional)
		MoveFrom      int64  `json:"move_from" url:"move_from,omitempty"`           //  (Optional)
		Title         string `json:"title" url:"title,omitempty"`                   //  (Optional)
		DisplayName   string `json:"display_name" url:"display_name,omitempty"`     //  (Optional)
		Description   string `json:"description" url:"description,omitempty"`       //  (Optional)
		VendorGuid    string `json:"vendor_guid" url:"vendor_guid,omitempty"`       //  (Optional)
		MasteryPoints int64  `json:"mastery_points" url:"mastery_points,omitempty"` //  (Optional)
		Ratings       struct {
			Description []string `json:"description" url:"description,omitempty"` //  (Optional)
			Points      []int64  `json:"points" url:"points,omitempty"`           //  (Optional)
		} `json:"ratings" url:"ratings,omitempty"`

		CalculationMethod string `json:"calculation_method" url:"calculation_method,omitempty"` //  (Optional) . Must be one of decaying_average, n_mastery, latest, highest
		CalculationInt    int64  `json:"calculation_int" url:"calculation_int,omitempty"`       //  (Optional)
	} `json:"form"`
}

func (t *CreateLinkOutcomeCourses) GetMethod() string {
	return "POST"
}

func (t *CreateLinkOutcomeCourses) GetURLPath() string {
	path := "courses/{course_id}/outcome_groups/{id}/outcomes"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *CreateLinkOutcomeCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateLinkOutcomeCourses) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *CreateLinkOutcomeCourses) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *CreateLinkOutcomeCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if t.Form.CalculationMethod != "" && !string_utils.Include([]string{"decaying_average", "n_mastery", "latest", "highest"}, t.Form.CalculationMethod) {
		errs = append(errs, "CalculationMethod must be one of decaying_average, n_mastery, latest, highest")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateLinkOutcomeCourses) Do(c *canvasapi.Canvas) (*models.OutcomeLink, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.OutcomeLink{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
