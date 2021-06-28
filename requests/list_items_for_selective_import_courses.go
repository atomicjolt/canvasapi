package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/string_utils"
)

// ListItemsForSelectiveImportCourses Enumerates the content available for selective import in a tree structure. Each node provides
// a +property+ copy argument that can be supplied to the {api:ContentMigrationsController#update Update endpoint}
// to selectively copy the content associated with that tree node and its children. Each node may also
// provide a +sub_items_url+ or an array of +sub_items+ which you can use to obtain copy parameters
// for a subset of the resources in a given node.
//
// If no +type+ is sent you will get a list of the top-level sections in the content. It will look something like this:
//
//   [{
//     "type": "course_settings",
//     "property": "copy[all_course_settings]",
//     "title": "Course Settings"
//   },
//   {
//     "type": "context_modules",
//     "property": "copy[all_context_modules]",
//     "title": "Modules",
//     "count": 5,
//     "sub_items_url": "http://example.com/api/v1/courses/22/content_migrations/77/selective_data?type=context_modules"
//   },
//   {
//     "type": "assignments",
//     "property": "copy[all_assignments]",
//     "title": "Assignments",
//     "count": 2,
//     "sub_items_url": "http://localhost:3000/api/v1/courses/22/content_migrations/77/selective_data?type=assignments"
//   }]
//
// When a +type+ is provided, nodes may be further divided via +sub_items+. For example, using +type=assignments+
// results in a node for each assignment group and a sub_item for each assignment, like this:
//
//   [{
//     "type": "assignment_groups",
//     "title": "An Assignment Group",
//     "property": "copy[assignment_groups][id_i855cf145e5acc7435e1bf1c6e2126e5f]",
//     "sub_items": [{
//         "type": "assignments",
//         "title": "Assignment 1",
//         "property": "copy[assignments][id_i2102a7fa93b29226774949298626719d]"
//     }, {
//         "type": "assignments",
//         "title": "Assignment 2",
//         "property": "copy[assignments][id_i310cba275dc3f4aa8a3306bbbe380979]"
//     }]
//   }]
//
//
// To import the items corresponding to a particular tree node, use the +property+ as a parameter to the
// {api:ContentMigrationsController#update Update endpoint} and assign a value of 1, for example:
//
//   copy[assignments][id_i310cba275dc3f4aa8a3306bbbe380979]=1
//
// You can include multiple copy parameters to selectively import multiple items or groups of items.
// https://canvas.instructure.com/doc/api/content_migrations.html
//
// Path Parameters:
// # CourseID (Required) ID
// # ID (Required) ID
//
// Query Parameters:
// # Type (Optional) . Must be one of context_modules, assignments, quizzes, assessment_question_banks, discussion_topics, wiki_pages, context_external_tools, tool_profiles, announcements, calendar_events, rubrics, groups, learning_outcomes, attachmentsThe type of content to enumerate.
//
type ListItemsForSelectiveImportCourses struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
		ID       string `json:"id" url:"id,omitempty"`               //  (Required)
	} `json:"path"`

	Query struct {
		Type string `json:"type" url:"type,omitempty"` //  (Optional) . Must be one of context_modules, assignments, quizzes, assessment_question_banks, discussion_topics, wiki_pages, context_external_tools, tool_profiles, announcements, calendar_events, rubrics, groups, learning_outcomes, attachments
	} `json:"query"`
}

func (t *ListItemsForSelectiveImportCourses) GetMethod() string {
	return "GET"
}

func (t *ListItemsForSelectiveImportCourses) GetURLPath() string {
	path := "courses/{course_id}/content_migrations/{id}/selective_data"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *ListItemsForSelectiveImportCourses) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *ListItemsForSelectiveImportCourses) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListItemsForSelectiveImportCourses) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListItemsForSelectiveImportCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if t.Query.Type != "" && !string_utils.Include([]string{"context_modules", "assignments", "quizzes", "assessment_question_banks", "discussion_topics", "wiki_pages", "context_external_tools", "tool_profiles", "announcements", "calendar_events", "rubrics", "groups", "learning_outcomes", "attachments"}, t.Query.Type) {
		errs = append(errs, "Type must be one of context_modules, assignments, quizzes, assessment_question_banks, discussion_topics, wiki_pages, context_external_tools, tool_profiles, announcements, calendar_events, rubrics, groups, learning_outcomes, attachments")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListItemsForSelectiveImportCourses) Do(c *canvasapi.Canvas) (*[]string, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []string{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
