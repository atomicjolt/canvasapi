package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// ListFeaturesUsers A paginated list of all features that apply to a given Account, Course, or User.
// https://canvas.instructure.com/doc/api/feature_flags.html
//
// Path Parameters:
// # Path.UserID (Required) ID
//
type ListFeaturesUsers struct {
	Path struct {
		UserID string `json:"user_id" url:"user_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ListFeaturesUsers) GetMethod() string {
	return "GET"
}

func (t *ListFeaturesUsers) GetURLPath() string {
	path := "users/{user_id}/features"
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	return path
}

func (t *ListFeaturesUsers) GetQuery() (string, error) {
	return "", nil
}

func (t *ListFeaturesUsers) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListFeaturesUsers) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListFeaturesUsers) HasErrors() error {
	errs := []string{}
	if t.Path.UserID == "" {
		errs = append(errs, "'Path.UserID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListFeaturesUsers) Do(c *canvasapi.Canvas, next *url.URL) ([]*models.Feature, *canvasapi.PagedResource, error) {
	var err error
	var response *http.Response
	if next != nil {
		response, err = c.Send(next, t.GetMethod(), nil)
	} else {
		response, err = c.SendRequest(t)
	}

	if err != nil {
		return nil, nil, err
	}
	if err != nil {
		return nil, nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, nil, err
	}
	ret := []*models.Feature{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, nil, err
	}

	pagedResource, err := canvasapi.ExtractPagedResource(response.Header)
	if err != nil {
		return nil, nil, err
	}

	return ret, pagedResource, nil
}
