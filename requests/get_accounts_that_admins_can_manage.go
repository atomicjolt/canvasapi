package requests

import (
	"encoding/json"
	"io/ioutil"

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

func (t *GetAccountsThatAdminsCanManage) GetBody() (string, error) {
	return "", nil
}

func (t *GetAccountsThatAdminsCanManage) HasErrors() error {
	return nil
}

func (t *GetAccountsThatAdminsCanManage) Do(c *canvasapi.Canvas) ([]*models.Account, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.Account{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
