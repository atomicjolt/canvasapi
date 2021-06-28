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

// CreateContentShare Share content directly between two or more users
// https://canvas.instructure.com/doc/api/content_shares.html
//
// Path Parameters:
// # UserID (Required) ID
//
// Form Parameters:
// # ReceiverIDs (Required) IDs of users to share the content with.
// # ContentType (Required) . Must be one of assignment, discussion_topic, page, quiz, module, module_itemType of content you are sharing.
// # ContentID (Required) The id of the content that you are sharing
//
type CreateContentShare struct {
	Path struct {
		UserID string `json:"user_id" url:"user_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		ReceiverIDs string `json:"receiver_ids" url:"receiver_ids,omitempty"` //  (Required)
		ContentType string `json:"content_type" url:"content_type,omitempty"` //  (Required) . Must be one of assignment, discussion_topic, page, quiz, module, module_item
		ContentID   int64  `json:"content_id" url:"content_id,omitempty"`     //  (Required)
	} `json:"form"`
}

func (t *CreateContentShare) GetMethod() string {
	return "POST"
}

func (t *CreateContentShare) GetURLPath() string {
	path := "users/{user_id}/content_shares"
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	return path
}

func (t *CreateContentShare) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateContentShare) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *CreateContentShare) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *CreateContentShare) HasErrors() error {
	errs := []string{}
	if t.Path.UserID == "" {
		errs = append(errs, "'UserID' is required")
	}
	if t.Form.ReceiverIDs == "" {
		errs = append(errs, "'ReceiverIDs' is required")
	}
	if t.Form.ContentType == "" {
		errs = append(errs, "'ContentType' is required")
	}
	if t.Form.ContentType != "" && !string_utils.Include([]string{"assignment", "discussion_topic", "page", "quiz", "module", "module_item"}, t.Form.ContentType) {
		errs = append(errs, "ContentType must be one of assignment, discussion_topic, page, quiz, module, module_item")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateContentShare) Do(c *canvasapi.Canvas) (*models.ContentShare, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.ContentShare{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
