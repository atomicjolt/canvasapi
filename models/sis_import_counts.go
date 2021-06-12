package models

type SISImportCounts struct {
	Accounts                int64 `json:"accounts"`                  // Example: 0
	Terms                   int64 `json:"terms"`                     // Example: 3
	AbstractCourses         int64 `json:"abstract_courses"`          // Example: 0
	Courses                 int64 `json:"courses"`                   // Example: 121
	Sections                int64 `json:"sections"`                  // Example: 278
	Xlists                  int64 `json:"xlists"`                    // Example: 0
	Users                   int64 `json:"users"`                     // Example: 346
	Enrollments             int64 `json:"enrollments"`               // Example: 1542
	Groups                  int64 `json:"groups"`                    // Example: 0
	GroupMemberships        int64 `json:"group_memberships"`         // Example: 0
	GradePublishingResults  int64 `json:"grade_publishing_results"`  // Example: 0
	BatchCoursesDeleted     int64 `json:"batch_courses_deleted"`     // the number of courses that were removed because they were not included in the batch for batch_mode imports. Only included if courses were deleted.Example: 11
	BatchSectionsDeleted    int64 `json:"batch_sections_deleted"`    // the number of sections that were removed because they were not included in the batch for batch_mode imports. Only included if sections were deleted.Example: 0
	BatchEnrollmentsDeleted int64 `json:"batch_enrollments_deleted"` // the number of enrollments that were removed because they were not included in the batch for batch_mode imports. Only included if enrollments were deleted.Example: 150
	ErrorCount              int64 `json:"error_count"`               // Example: 0
	WarningCount            int64 `json:"warning_count"`             // Example: 0
}

func (t *SISImportCounts) HasError() error {
	return nil
}
