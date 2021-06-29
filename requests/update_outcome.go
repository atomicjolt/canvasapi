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

// UpdateOutcome Modify an existing outcome. Fields not provided are left as is;
// unrecognized fields are ignored.
//
// If any new ratings are provided, the combination of all new ratings
// provided completely replace any existing embedded rubric criterion; it is
// not possible to tweak the ratings of the embedded rubric criterion.
//
// A new embedded rubric criterion's mastery_points default to the maximum
// points in the highest rating if not specified in the mastery_points
// parameter. Any new ratings lacking a description are given a default of "No
// description". Any new ratings lacking a point value are given a default of
// 0.
// https://canvas.instructure.com/doc/api/outcomes.html
//
// Path Parameters:
// # Path.ID (Required) ID
//
// Form Parameters:
// # Form.Title (Optional) The new outcome title.
// # Form.DisplayName (Optional) A friendly name shown in reports for outcomes with cryptic titles,
//    such as common core standards names.
// # Form.Description (Optional) The new outcome description.
// # Form.VendorGuid (Optional) A custom GUID for the learning standard.
// # Form.MasteryPoints (Optional) The new mastery threshold for the embedded rubric criterion.
// # Form.Ratings.Description (Optional) The description of a new rating level for the embedded rubric criterion.
// # Form.Ratings.Points (Optional) The points corresponding to a new rating level for the embedded rubric
//    criterion.
// # Form.CalculationMethod (Optional) . Must be one of decaying_average, n_mastery, latest, highestThe new calculation method.
// # Form.CalculationInt (Optional) The new calculation int.  Only applies if the calculation_method is "decaying_average" or "n_mastery"
//
type UpdateOutcome struct {
	Path struct {
		ID string `json:"id" url:"id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		Title         string `json:"title" url:"title,omitempty"`                   //  (Optional)
		DisplayName   string `json:"display_name" url:"display_name,omitempty"`     //  (Optional)
		Description   string `json:"description" url:"description,omitempty"`       //  (Optional)
		VendorGuid    string `json:"vendor_guid" url:"vendor_guid,omitempty"`       //  (Optional)
		MasteryPoints int64  `json:"mastery_points" url:"mastery_points,omitempty"` //  (Optional)
		Ratings       struct {
			Description []string `json:"description" url:"description,omitempty"` //  (Optional)
			Points      []string `json:"points" url:"points,omitempty"`           //  (Optional)
		} `json:"ratings" url:"ratings,omitempty"`

		CalculationMethod string `json:"calculation_method" url:"calculation_method,omitempty"` //  (Optional) . Must be one of decaying_average, n_mastery, latest, highest
		CalculationInt    int64  `json:"calculation_int" url:"calculation_int,omitempty"`       //  (Optional)
	} `json:"form"`
}

func (t *UpdateOutcome) GetMethod() string {
	return "PUT"
}

func (t *UpdateOutcome) GetURLPath() string {
	path := "outcomes/{id}"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *UpdateOutcome) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdateOutcome) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *UpdateOutcome) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *UpdateOutcome) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if t.Form.CalculationMethod != "" && !string_utils.Include([]string{"decaying_average", "n_mastery", "latest", "highest"}, t.Form.CalculationMethod) {
		errs = append(errs, "CalculationMethod must be one of decaying_average, n_mastery, latest, highest")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UpdateOutcome) Do(c *canvasapi.Canvas) (*models.Outcome, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.Outcome{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
