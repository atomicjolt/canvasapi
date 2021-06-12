package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/string_utils"
)

// CopyCourseContent DEPRECATED: Please use the {api:ContentMigrationsController#create Content Migrations API}
//
// Copies content from one course into another. The default is to copy all course
// content. You can control specific types to copy by using either the 'except' option
// or the 'only' option.
//
// The response is the same as the course copy status endpoint
// https://canvas.instructure.com/doc/api/courses.html
//
// Path Parameters:
// # CourseID (Required) ID
//
// Form Parameters:
// # SourceCourse (Optional) ID or SIS-ID of the course to copy the content from
// # Except (Optional) . Must be one of course_settings, assignments, external_tools, files, topics, calendar_events, quizzes, wiki_pages, modules, outcomesA list of the course content types to exclude, all areas not listed will
//    be copied.
// # Only (Optional) . Must be one of course_settings, assignments, external_tools, files, topics, calendar_events, quizzes, wiki_pages, modules, outcomesA list of the course content types to copy, all areas not listed will not
//    be copied.
//
type CopyCourseContent struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
	} `json:"path"`

	Form struct {
		SourceCourse string   `json:"source_course"` //  (Optional)
		Except       []string `json:"except"`        //  (Optional) . Must be one of course_settings, assignments, external_tools, files, topics, calendar_events, quizzes, wiki_pages, modules, outcomes
		Only         []string `json:"only"`          //  (Optional) . Must be one of course_settings, assignments, external_tools, files, topics, calendar_events, quizzes, wiki_pages, modules, outcomes
	} `json:"form"`
}

func (t *CopyCourseContent) GetMethod() string {
	return "POST"
}

func (t *CopyCourseContent) GetURLPath() string {
	path := "courses/{course_id}/course_copy"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *CopyCourseContent) GetQuery() (string, error) {
	return "", nil
}

func (t *CopyCourseContent) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *CopyCourseContent) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	for _, v := range t.Form.Except {
		if !string_utils.Include([]string{"course_settings", "assignments", "external_tools", "files", "topics", "calendar_events", "quizzes", "wiki_pages", "modules", "outcomes"}, v) {
			errs = append(errs, "Except must be one of course_settings, assignments, external_tools, files, topics, calendar_events, quizzes, wiki_pages, modules, outcomes")
		}
	}
	for _, v := range t.Form.Only {
		if !string_utils.Include([]string{"course_settings", "assignments", "external_tools", "files", "topics", "calendar_events", "quizzes", "wiki_pages", "modules", "outcomes"}, v) {
			errs = append(errs, "Only must be one of course_settings, assignments, external_tools, files, topics, calendar_events, quizzes, wiki_pages, modules, outcomes")
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CopyCourseContent) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}