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

// ListFilesFolders Returns the paginated list of files for the folder or course.
// https://canvas.instructure.com/doc/api/files.html
//
// Path Parameters:
// # Path.ID (Required) ID
//
// Query Parameters:
// # Query.ContentTypes (Optional) Filter results by content-type. You can specify type/subtype pairs (e.g.,
//    'image/jpeg'), or simply types (e.g., 'image', which will match
//    'image/gif', 'image/jpeg', etc.).
// # Query.ExcludeContentTypes (Optional) Exclude given content-types from your results. You can specify type/subtype pairs (e.g.,
//    'image/jpeg'), or simply types (e.g., 'image', which will match
//    'image/gif', 'image/jpeg', etc.).
// # Query.SearchTerm (Optional) The partial name of the files to match and return.
// # Query.Include (Optional) . Must be one of userArray of additional information to include.
//
//    "user":: the user who uploaded the file or last edited its content
//    "usage_rights":: copyright and license information for the file (see UsageRights)
// # Query.Only (Optional) Array of information to restrict to. Overrides include[]
//
//    "names":: only returns file name information
// # Query.Sort (Optional) . Must be one of name, size, created_at, updated_at, content_type, userSort results by this field. Defaults to 'name'. Note that `sort=user` implies `include[]=user`.
// # Query.Order (Optional) . Must be one of asc, descThe sorting order. Defaults to 'asc'.
//
type ListFilesFolders struct {
	Path struct {
		ID string `json:"id" url:"id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		ContentTypes        []string `json:"content_types" url:"content_types,omitempty"`                 //  (Optional)
		ExcludeContentTypes []string `json:"exclude_content_types" url:"exclude_content_types,omitempty"` //  (Optional)
		SearchTerm          string   `json:"search_term" url:"search_term,omitempty"`                     //  (Optional)
		Include             []string `json:"include" url:"include,omitempty"`                             //  (Optional) . Must be one of user
		Only                []string `json:"only" url:"only,omitempty"`                                   //  (Optional)
		Sort                string   `json:"sort" url:"sort,omitempty"`                                   //  (Optional) . Must be one of name, size, created_at, updated_at, content_type, user
		Order               string   `json:"order" url:"order,omitempty"`                                 //  (Optional) . Must be one of asc, desc
	} `json:"query"`
}

func (t *ListFilesFolders) GetMethod() string {
	return "GET"
}

func (t *ListFilesFolders) GetURLPath() string {
	path := "folders/{id}/files"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *ListFilesFolders) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *ListFilesFolders) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListFilesFolders) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListFilesFolders) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	for _, v := range t.Query.Include {
		if v != "" && !string_utils.Include([]string{"user"}, v) {
			errs = append(errs, "Include must be one of user")
		}
	}
	if t.Query.Sort != "" && !string_utils.Include([]string{"name", "size", "created_at", "updated_at", "content_type", "user"}, t.Query.Sort) {
		errs = append(errs, "Sort must be one of name, size, created_at, updated_at, content_type, user")
	}
	if t.Query.Order != "" && !string_utils.Include([]string{"asc", "desc"}, t.Query.Order) {
		errs = append(errs, "Order must be one of asc, desc")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListFilesFolders) Do(c *canvasapi.Canvas) ([]*models.File, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.File{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
