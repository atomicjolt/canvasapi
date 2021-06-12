package models

import (
	"fmt"

	"github.com/atomicjolt/string_utils"
)

type ModuleItem struct {
	ID                    int64                  `json:"id"`                     // the unique identifier for the module item.Example: 768
	ModuleID              int64                  `json:"module_id"`              // the id of the Module this item appears in.Example: 123
	Position              int64                  `json:"position"`               // the position of this item in the module (1-based).Example: 1
	Title                 string                 `json:"title"`                  // the title of this item.Example: Square Roots: Irrational numbers or boxy vegetables?
	Indent                int64                  `json:"indent"`                 // 0-based indent level; module items may be indented to show a hierarchy.Example: 0
	Type                  string                 `json:"type"`                   // the type of object referred to one of 'File', 'Page', 'Discussion', 'Assignment', 'Quiz', 'SubHeader', 'ExternalUrl', 'ExternalTool'.Example: Assignment
	ContentID             int64                  `json:"content_id"`             // the id of the object referred to applies to 'File', 'Discussion', 'Assignment', 'Quiz', 'ExternalTool' types.Example: 1337
	HtmlUrl               string                 `json:"html_url"`               // link to the item in Canvas.Example: https://canvas.example.edu/courses/222/modules/items/768
	Url                   string                 `json:"url"`                    // (Optional) link to the Canvas API object, if applicable.Example: https://canvas.example.edu/api/v1/courses/222/assignments/987
	PageUrl               string                 `json:"page_url"`               // (only for 'Page' type) unique locator for the linked wiki page.Example: my-page-title
	ExternalUrl           string                 `json:"external_url"`           // (only for 'ExternalUrl' and 'ExternalTool' types) external url that the item points to.Example: https://www.example.com/externalurl
	NewTab                bool                   `json:"new_tab"`                // (only for 'ExternalTool' type) whether the external tool opens in a new tab.
	CompletionRequirement *CompletionRequirement `json:"completion_requirement"` // Completion requirement for this module item.Example: min_score, 10, true
	ContentDetails        *ContentDetails        `json:"content_details"`        // (Present only if requested through include[]=content_details) If applicable, returns additional details specific to the associated object.Example: 20, 2012-12-31T06:00:00-06:00, 2012-12-31T06:00:00-06:00, 2012-12-31T06:00:00-06:00
	Published             bool                   `json:"published"`              // (Optional) Whether this module item is published. This field is present only if the caller has permission to view unpublished items..Example: true
}

func (t *ModuleItem) HasError() error {
	var s []string
	s = []string{"File", "Page", "Discussion", "Assignment", "Quiz", "SubHeader", "ExternalUrl", "ExternalTool"}
	if !string_utils.Include(s, t.Type) {
		return fmt.Errorf("expected 'type' to be one of %v", s)
	}
	return nil
}