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

// ImportSISData Import SIS data into Canvas. Must be on a root account with SIS imports
// enabled.
//
// For more information on the format that's expected here, please see the
// "SIS CSV" section in the API docs.
// https://canvas.instructure.com/doc/api/sis_imports.html
//
// Path Parameters:
// # Path.AccountID (Required) ID
//
// Form Parameters:
// # Form.ImportType (Optional) Choose the data format for reading SIS data. With a standard Canvas
//    install, this option can only be 'instructure_csv', and if unprovided,
//    will be assumed to be so. Can be part of the query string.
// # Form.Attachment (Optional) There are two ways to post SIS import data - either via a
//    multipart/form-data form-field-style attachment, or via a non-multipart
//    raw post request.
//
//    'attachment' is required for multipart/form-data style posts. Assumed to
//    be SIS data from a file upload form field named 'attachment'.
//
//    Examples:
//      curl -F attachment=@<filename> -H "Authorization: Bearer <token>" \
//          https://<canvas>/api/v1/accounts/<account_id>/sis_imports.json?import_type=instructure_csv
//
//    If you decide to do a raw post, you can skip the 'attachment' argument,
//    but you will then be required to provide a suitable Content-Type header.
//    You are encouraged to also provide the 'extension' argument.
//
//    Examples:
//      curl -H 'Content-Type: application/octet-stream' --data-binary @<filename>.zip \
//          -H "Authorization: Bearer <token>" \
//          https://<canvas>/api/v1/accounts/<account_id>/sis_imports.json?import_type=instructure_csv&extension=zip
//
//      curl -H 'Content-Type: application/zip' --data-binary @<filename>.zip \
//          -H "Authorization: Bearer <token>" \
//          https://<canvas>/api/v1/accounts/<account_id>/sis_imports.json?import_type=instructure_csv
//
//      curl -H 'Content-Type: text/csv' --data-binary @<filename>.csv \
//          -H "Authorization: Bearer <token>" \
//          https://<canvas>/api/v1/accounts/<account_id>/sis_imports.json?import_type=instructure_csv
//
//      curl -H 'Content-Type: text/csv' --data-binary @<filename>.csv \
//          -H "Authorization: Bearer <token>" \
//          https://<canvas>/api/v1/accounts/<account_id>/sis_imports.json?import_type=instructure_csv&batch_mode=1&batch_mode_term_id=15
// # Form.Extension (Optional) Recommended for raw post request style imports. This field will be used to
//    distinguish between zip, xml, csv, and other file format extensions that
//    would usually be provided with the filename in the multipart post request
//    scenario. If not provided, this value will be inferred from the
//    Content-Type, falling back to zip-file format if all else fails.
// # Form.BatchMode (Optional) If set, this SIS import will be run in batch mode, deleting any data
//    previously imported via SIS that is not present in this latest import.
//    See the SIS CSV Format page for details.
//    Batch mode cannot be used with diffing.
// # Form.BatchModeTermID (Optional) Limit deletions to only this term. Required if batch mode is enabled.
// # Form.MultiTermBatchMode (Optional) Runs batch mode against all terms in terms file. Requires change_threshold.
// # Form.SkipDeletes (Optional) When set the import will skip any deletes. This does not account for
//    objects that are deleted during the batch mode cleanup process.
// # Form.OverrideSISStickiness (Optional) Many fields on records in Canvas can be marked "sticky," which means that
//    when something changes in the UI apart from the SIS, that field gets
//    "stuck." In this way, by default, SIS imports do not override UI changes.
//    If this field is present, however, it will tell the SIS import to ignore
//    "stickiness" and override all fields.
// # Form.AddSISStickiness (Optional) This option, if present, will process all changes as if they were UI
//    changes. This means that "stickiness" will be added to changed fields.
//    This option is only processed if 'override_sis_stickiness' is also provided.
// # Form.ClearSISStickiness (Optional) This option, if present, will clear "stickiness" from all fields touched
//    by this import. Requires that 'override_sis_stickiness' is also provided.
//    If 'add_sis_stickiness' is also provided, 'clear_sis_stickiness' will
//    overrule the behavior of 'add_sis_stickiness'
// # Form.DiffingDataSetIDentifier (Optional) If set on a CSV import, Canvas will attempt to optimize the SIS import by
//    comparing this set of CSVs to the previous set that has the same data set
//    identifier, and only applying the difference between the two. See the
//    SIS CSV Format documentation for more details.
//    Diffing cannot be used with batch_mode
// # Form.DiffingRemasterDataSet (Optional) If true, and diffing_data_set_identifier is sent, this SIS import will be
//    part of the data set, but diffing will not be performed. See the SIS CSV
//    Format documentation for details.
// # Form.DiffingDropStatus (Optional) . Must be one of deleted, completed, inactiveIf diffing_drop_status is passed, this SIS import will use this status for
//    enrollments that are not included in the sis_batch. Defaults to 'deleted'
// # Form.BatchModeEnrollmentDropStatus (Optional) . Must be one of deleted, completed, inactiveIf batch_mode_enrollment_drop_status is passed, this SIS import will use
//    this status for enrollments that are not included in the sis_batch. This
//    will have an effect if multi_term_batch_mode is set. Defaults to 'deleted'
//    This will still mark courses and sections that are not included in the
//    sis_batch as deleted, and subsequently enrollments in the deleted courses
//    and sections as deleted.
// # Form.ChangeThreshold (Optional) If set with batch_mode, the batch cleanup process will not run if the
//    number of items deleted is higher than the percentage set. If set to 10
//    and a term has 200 enrollments, and batch would delete more than 20 of
//    the enrollments the batch will abort before the enrollments are deleted.
//    The change_threshold will be evaluated for course, sections, and
//    enrollments independently.
//    If set with diffing, diffing will not be performed if the files are
//    greater than the threshold as a percent. If set to 5 and the file is more
//    than 5% smaller or more than 5% larger than the file that is being
//    compared to, diffing will not be performed. If the files are less than 5%,
//    diffing will be performed. The way the percent is calculated is by taking
//    the size of the current import and dividing it by the size of the previous
//    import. The formula used is:
//    |(1 - current_file_size / previous_file_size)| * 100
//    See the SIS CSV Format documentation for more details.
//    Required for multi_term_batch_mode.
// # Form.DiffRowCountThreshold (Optional) If set with diffing, diffing will not be performed if the number of rows
//    to be run in the fully calculated diff import exceeds the threshold.
//
type ImportSISData struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		ImportType                    string `json:"import_type" url:"import_type,omitempty"`                                             //  (Optional)
		Attachment                    string `json:"attachment" url:"attachment,omitempty"`                                               //  (Optional)
		Extension                     string `json:"extension" url:"extension,omitempty"`                                                 //  (Optional)
		BatchMode                     bool   `json:"batch_mode" url:"batch_mode,omitempty"`                                               //  (Optional)
		BatchModeTermID               string `json:"batch_mode_term_id" url:"batch_mode_term_id,omitempty"`                               //  (Optional)
		MultiTermBatchMode            bool   `json:"multi_term_batch_mode" url:"multi_term_batch_mode,omitempty"`                         //  (Optional)
		SkipDeletes                   bool   `json:"skip_deletes" url:"skip_deletes,omitempty"`                                           //  (Optional)
		OverrideSISStickiness         bool   `json:"override_sis_stickiness" url:"override_sis_stickiness,omitempty"`                     //  (Optional)
		AddSISStickiness              bool   `json:"add_sis_stickiness" url:"add_sis_stickiness,omitempty"`                               //  (Optional)
		ClearSISStickiness            bool   `json:"clear_sis_stickiness" url:"clear_sis_stickiness,omitempty"`                           //  (Optional)
		DiffingDataSetIDentifier      string `json:"diffing_data_set_identifier" url:"diffing_data_set_identifier,omitempty"`             //  (Optional)
		DiffingRemasterDataSet        bool   `json:"diffing_remaster_data_set" url:"diffing_remaster_data_set,omitempty"`                 //  (Optional)
		DiffingDropStatus             string `json:"diffing_drop_status" url:"diffing_drop_status,omitempty"`                             //  (Optional) . Must be one of deleted, completed, inactive
		BatchModeEnrollmentDropStatus string `json:"batch_mode_enrollment_drop_status" url:"batch_mode_enrollment_drop_status,omitempty"` //  (Optional) . Must be one of deleted, completed, inactive
		ChangeThreshold               int64  `json:"change_threshold" url:"change_threshold,omitempty"`                                   //  (Optional)
		DiffRowCountThreshold         int64  `json:"diff_row_count_threshold" url:"diff_row_count_threshold,omitempty"`                   //  (Optional)
	} `json:"form"`
}

func (t *ImportSISData) GetMethod() string {
	return "POST"
}

func (t *ImportSISData) GetURLPath() string {
	path := "accounts/{account_id}/sis_imports"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *ImportSISData) GetQuery() (string, error) {
	return "", nil
}

func (t *ImportSISData) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *ImportSISData) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *ImportSISData) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'Path.AccountID' is required")
	}
	if t.Form.DiffingDropStatus != "" && !string_utils.Include([]string{"deleted", "completed", "inactive"}, t.Form.DiffingDropStatus) {
		errs = append(errs, "DiffingDropStatus must be one of deleted, completed, inactive")
	}
	if t.Form.BatchModeEnrollmentDropStatus != "" && !string_utils.Include([]string{"deleted", "completed", "inactive"}, t.Form.BatchModeEnrollmentDropStatus) {
		errs = append(errs, "BatchModeEnrollmentDropStatus must be one of deleted, completed, inactive")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ImportSISData) Do(c *canvasapi.Canvas) (*models.SISImport, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.SISImport{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
