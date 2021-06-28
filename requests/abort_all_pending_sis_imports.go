package requests

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
)

// AbortAllPendingSISImports Abort already created but not processed or processing SIS imports.
// https://canvas.instructure.com/doc/api/sis_imports.html
//
// Path Parameters:
// # AccountID (Required) ID
//
type AbortAllPendingSISImports struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *AbortAllPendingSISImports) GetMethod() string {
	return "PUT"
}

func (t *AbortAllPendingSISImports) GetURLPath() string {
	path := "accounts/{account_id}/sis_imports/abort_all_pending"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *AbortAllPendingSISImports) GetQuery() (string, error) {
	return "", nil
}

func (t *AbortAllPendingSISImports) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *AbortAllPendingSISImports) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *AbortAllPendingSISImports) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *AbortAllPendingSISImports) Do(c *canvasapi.Canvas) (bool, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return false, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return false, err
	}
	// TODO. I doubt these conversions to string and int below really work. Figure what Canvas returns and test against that return value
	ret := string(body) == "true"

	return ret, nil
}
