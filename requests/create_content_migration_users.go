package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
	"time"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
	"github.com/atomicjolt/string_utils"
)

// CreateContentMigrationUsers Create a content migration. If the migration requires a file to be uploaded
// the actual processing of the file will start once the file upload process is completed.
// File uploading works as described in the {file:file_uploads.html File Upload Documentation}
// except that the values are set on a *pre_attachment* sub-hash.
//
// For migrations that don't require a file to be uploaded, like course copy, the
// processing will begin as soon as the migration is created.
//
// You can use the {api:ProgressController#show Progress API} to track the
// progress of the migration. The migration's progress is linked to with the
// _progress_url_ value.
//
// The two general workflows are:
//
// If no file upload is needed:
//
// 1. POST to create
// 2. Use the {api:ProgressController#show Progress} specified in _progress_url_ to monitor progress
//
// For file uploading:
//
// 1. POST to create with file info in *pre_attachment*
// 2. Do {file:file_uploads.html file upload processing} using the data in the *pre_attachment* data
// 3. {api:ContentMigrationsController#show GET} the ContentMigration
// 4. Use the {api:ProgressController#show Progress} specified in _progress_url_ to monitor progress
//
//  (required if doing .zip file upload)
// https://canvas.instructure.com/doc/api/content_migrations.html
//
// Path Parameters:
// # Path.UserID (Required) ID
//
// Form Parameters:
// # Form.MigrationType (Required) The type of the migration. Use the
//    {api:ContentMigrationsController#available_migrators Migrator} endpoint to
//    see all available migrators. Default allowed values:
//    canvas_cartridge_importer, common_cartridge_importer,
//    course_copy_importer, zip_file_importer, qti_converter, moodle_converter
// # Form.PreAttachment.Name (Optional) Required if uploading a file. This is the first step in uploading a file
//    to the content migration. See the {file:file_uploads.html File Upload
//    Documentation} for details on the file upload workflow.
// # Form.PreAttachment.* (Optional) Other file upload properties, See {file:file_uploads.html File Upload
//    Documentation}
// # Form.Settings.FileUrl (Optional) A URL to download the file from. Must not require authentication.
// # Form.Settings.ContentExportID (Optional) The id of a ContentExport to import. This allows you to import content previously exported from Canvas
//    without needing to download and re-upload it.
// # Form.Settings.SourceCourseID (Optional) The course to copy from for a course copy migration. (required if doing
//    course copy)
// # Form.Settings.FolderID (Optional) The folder to unzip the .zip file into for a zip_file_import.
// # Form.Settings.OverwriteQuizzes (Optional) Whether to overwrite quizzes with the same identifiers between content
//    packages.
// # Form.Settings.QuestionBankID (Optional) The existing question bank ID to import questions into if not specified in
//    the content package.
// # Form.Settings.QuestionBankName (Optional) The question bank to import questions into if not specified in the content
//    package, if both bank id and name are set, id will take precedence.
// # Form.Settings.InsertIntoModuleID (Optional) The id of a module in the target course. This will add all imported items
//    (that can be added to a module) to the given module.
// # Form.Settings.InsertIntoModuleType (Optional) . Must be one of assignment, discussion_topic, file, page, quizIf provided (and +insert_into_module_id+ is supplied),
//    only add objects of the specified type to the module.
// # Form.Settings.InsertIntoModulePosition (Optional) The (1-based) position to insert the imported items into the course
//    (if +insert_into_module_id+ is supplied). If this parameter
//    is omitted, items will be added to the end of the module.
// # Form.Settings.MoveToAssignmentGroupID (Optional) The id of an assignment group in the target course. If provided, all
//    imported assignments will be moved to the given assignment group.
// # Form.DateShiftOptions.ShiftDates (Optional) Whether to shift dates in the copied course
// # Form.DateShiftOptions.OldStartDate (Optional) The original start date of the source content/course
// # Form.DateShiftOptions.OldEndDate (Optional) The original end date of the source content/course
// # Form.DateShiftOptions.NewStartDate (Optional) The new start date for the content/course
// # Form.DateShiftOptions.NewEndDate (Optional) The new end date for the source content/course
// # Form.DateShiftOptions (Optional) Move anything scheduled for day 'X' to the specified day. (0-Sunday,
//    1-Monday, 2-Tuesday, 3-Wednesday, 4-Thursday, 5-Friday, 6-Saturday)
// # Form.DateShiftOptions.RemoveDates (Optional) Whether to remove dates in the copied course. Cannot be used
//    in conjunction with *shift_dates*.
// # Form.SelectiveImport (Optional) If set, perform a selective import instead of importing all content.
//    The migration will identify the contents of the package and then stop
//    in the +waiting_for_select+ workflow state. At this point, use the
//    {api:ContentMigrationsController#content_list List items endpoint}
//    to enumerate the contents of the package, identifying the copy
//    parameters for the desired content. Then call the
//    {api:ContentMigrationsController#update Update endpoint} and provide these
//    copy parameters to start the import.
// # Form.Select (Optional) . Must be one of folders, files, attachments, quizzes, assignments, announcements, calendar_events, discussion_topics, modules, module_items, pages, rubricsFor +course_copy_importer+ migrations, this parameter allows you to select
//    the objects to copy without using the +selective_import+ argument and
//    +waiting_for_select+ state as is required for uploaded imports (though that
//    workflow is also supported for course copy migrations).
//    The keys are object types like 'files', 'folders', 'pages', etc. The value
//    for each key is a list of object ids. An id can be an integer or a string.
//    Multiple object types can be selected in the same call.
//
type CreateContentMigrationUsers struct {
	Path struct {
		UserID string `json:"user_id" url:"user_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		MigrationType string `json:"migration_type" url:"migration_type,omitempty"` //  (Required)
		PreAttachment struct {
			Name    string                       `json:"name" url:"name,omitempty"` //  (Optional)
			*string `json:"*" url:"*,omitempty"` //  (Optional)
		} `json:"pre_attachment" url:"pre_attachment,omitempty"`

		Settings struct {
			FileUrl                  string `json:"file_url" url:"file_url,omitempty"`                                       //  (Optional)
			ContentExportID          string `json:"content_export_id" url:"content_export_id,omitempty"`                     //  (Optional)
			SourceCourseID           string `json:"source_course_id" url:"source_course_id,omitempty"`                       //  (Optional)
			FolderID                 string `json:"folder_id" url:"folder_id,omitempty"`                                     //  (Optional)
			OverwriteQuizzes         bool   `json:"overwrite_quizzes" url:"overwrite_quizzes,omitempty"`                     //  (Optional)
			QuestionBankID           int64  `json:"question_bank_id" url:"question_bank_id,omitempty"`                       //  (Optional)
			QuestionBankName         string `json:"question_bank_name" url:"question_bank_name,omitempty"`                   //  (Optional)
			InsertIntoModuleID       int64  `json:"insert_into_module_id" url:"insert_into_module_id,omitempty"`             //  (Optional)
			InsertIntoModuleType     string `json:"insert_into_module_type" url:"insert_into_module_type,omitempty"`         //  (Optional) . Must be one of assignment, discussion_topic, file, page, quiz
			InsertIntoModulePosition int64  `json:"insert_into_module_position" url:"insert_into_module_position,omitempty"` //  (Optional)
			MoveToAssignmentGroupID  int64  `json:"move_to_assignment_group_id" url:"move_to_assignment_group_id,omitempty"` //  (Optional)
		} `json:"settings" url:"settings,omitempty"`

		DateShiftOptions struct {
			ShiftDates       bool      `json:"shift_dates" url:"shift_dates,omitempty"`       //  (Optional)
			OldStartDate     time.Time `json:"old_start_date" url:"old_start_date,omitempty"` //  (Optional)
			OldEndDate       time.Time `json:"old_end_date" url:"old_end_date,omitempty"`     //  (Optional)
			NewStartDate     time.Time `json:"new_start_date" url:"new_start_date,omitempty"` //  (Optional)
			NewEndDate       time.Time `json:"new_end_date" url:"new_end_date,omitempty"`     //  (Optional)
			DaySubstitutions struct {
				X int64 `json:"x" url:"x,omitempty"` //  (Optional)
			} `json:"day_substitutions" url:"day_substitutions,omitempty"`

			RemoveDates bool `json:"remove_dates" url:"remove_dates,omitempty"` //  (Optional)
		} `json:"date_shift_options" url:"date_shift_options,omitempty"`

		SelectiveImport bool                     `json:"selective_import" url:"selective_import,omitempty"` //  (Optional)
		Select          map[string](interface{}) `json:"select" url:"select,omitempty"`                     //  (Optional) . Must be one of folders, files, attachments, quizzes, assignments, announcements, calendar_events, discussion_topics, modules, module_items, pages, rubrics
	} `json:"form"`
}

func (t *CreateContentMigrationUsers) GetMethod() string {
	return "POST"
}

func (t *CreateContentMigrationUsers) GetURLPath() string {
	path := "users/{user_id}/content_migrations"
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	return path
}

func (t *CreateContentMigrationUsers) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateContentMigrationUsers) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *CreateContentMigrationUsers) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *CreateContentMigrationUsers) HasErrors() error {
	errs := []string{}
	if t.Path.UserID == "" {
		errs = append(errs, "'Path.UserID' is required")
	}
	if t.Form.MigrationType == "" {
		errs = append(errs, "'Form.MigrationType' is required")
	}
	if t.Form.Settings.InsertIntoModuleType != "" && !string_utils.Include([]string{"assignment", "discussion_topic", "file", "page", "quiz"}, t.Form.Settings.InsertIntoModuleType) {
		errs = append(errs, "Settings must be one of assignment, discussion_topic, file, page, quiz")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateContentMigrationUsers) Do(c *canvasapi.Canvas) (*models.ContentMigration, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.ContentMigration{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
