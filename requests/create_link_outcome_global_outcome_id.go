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

// CreateLinkOutcomeGlobalOutcomeID Link an outcome into the outcome group. The outcome to link can either be
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
// # ID (Required) ID
// # OutcomeID (Required) The ID of the existing outcome to link.
//
// Form Parameters:
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
type CreateLinkOutcomeGlobalOutcomeID struct {
	Path struct {
		ID        string `json:"id"`         //  (Required)
		OutcomeID int64  `json:"outcome_id"` //  (Required)
	} `json:"path"`

	Form struct {
		MoveFrom      int64  `json:"move_from"`      //  (Optional)
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

func (t *CreateLinkOutcomeGlobalOutcomeID) GetMethod() string {
	return "PUT"
}

func (t *CreateLinkOutcomeGlobalOutcomeID) GetURLPath() string {
	path := "global/outcome_groups/{id}/outcomes/{outcome_id}"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	path = strings.ReplaceAll(path, "{outcome_id}", fmt.Sprintf("%v", t.Path.OutcomeID))
	return path
}

func (t *CreateLinkOutcomeGlobalOutcomeID) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateLinkOutcomeGlobalOutcomeID) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *CreateLinkOutcomeGlobalOutcomeID) HasErrors() error {
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

func (t *CreateLinkOutcomeGlobalOutcomeID) Do(c *canvasapi.Canvas) (*models.OutcomeLink, error) {
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
