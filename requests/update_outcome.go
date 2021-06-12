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
// # ID (Required) ID
//
// Form Parameters:
// # Title (Optional) The new outcome title.
// # DisplayName (Optional) A friendly name shown in reports for outcomes with cryptic titles,
//    such as common core standards names.
// # Description (Optional) The new outcome description.
// # VendorGuid (Optional) A custom GUID for the learning standard.
// # MasteryPoints (Optional) The new mastery threshold for the embedded rubric criterion.
// # Ratings (Optional) The description of a new rating level for the embedded rubric criterion.
// # Ratings (Optional) The points corresponding to a new rating level for the embedded rubric
//    criterion.
// # CalculationMethod (Optional) . Must be one of decaying_average, n_mastery, latest, highestThe new calculation method.
// # CalculationInt (Optional) The new calculation int.  Only applies if the calculation_method is "decaying_average" or "n_mastery"
//
type UpdateOutcome struct {
	Path struct {
		ID string `json:"id"` //  (Required)
	} `json:"path"`

	Form struct {
		Title         string `json:"title"`          //  (Optional)
		DisplayName   string `json:"display_name"`   //  (Optional)
		Description   string `json:"description"`    //  (Optional)
		VendorGuid    string `json:"vendor_guid"`    //  (Optional)
		MasteryPoints int64  `json:"mastery_points"` //  (Optional)
		Ratings       struct {
			Description []string `json:"description"` //  (Optional)
			Points      []int64  `json:"points"`      //  (Optional)
		} `json:"ratings"`

		CalculationMethod string `json:"calculation_method"` //  (Optional) . Must be one of decaying_average, n_mastery, latest, highest
		CalculationInt    int64  `json:"calculation_int"`    //  (Optional)
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

func (t *UpdateOutcome) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *UpdateOutcome) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if !string_utils.Include([]string{"decaying_average", "n_mastery", "latest", "highest"}, t.Form.CalculationMethod) {
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
