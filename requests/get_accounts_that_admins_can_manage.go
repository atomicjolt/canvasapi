package requests

import (
	"encoding/json"
	"io/ioutil"
	"net/url"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// GetAccountsThatAdminsCanManage A paginated list of accounts where the current user has permission to create
// or manage courses. List will be empty for students and teachers as only admins
// can view which accounts they are in.
// https://canvas.instructure.com/doc/api/accounts.html
//
type GetAccountsThatAdminsCanManage struct {
}

func (t *GetAccountsThatAdminsCanManage) GetMethod() string {
	return "GET"
}

func (t *GetAccountsThatAdminsCanManage) GetURLPath() string {
	return ""
}

func (t *GetAccountsThatAdminsCanManage) GetQuery() (string, error) {
	return "", nil
}

func (t *GetAccountsThatAdminsCanManage) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetAccountsThatAdminsCanManage) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetAccountsThatAdminsCanManage) HasErrors() error {
	return nil
}

func (t *GetAccountsThatAdminsCanManage) Do(c *canvasapi.Canvas) ([]*models.Account, *canvasapi.PagedResource, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, nil, err
	}
	ret := []*models.Account{}
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
