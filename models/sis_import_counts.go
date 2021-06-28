package models

type SISImportCounts struct {
	Accounts                int64 `json:"accounts" url:"accounts,omitempty"`                                   // Example: 0
	Terms                   int64 `json:"terms" url:"terms,omitempty"`                                         // Example: 3
	AbstractCourses         int64 `json:"abstract_courses" url:"abstract_courses,omitempty"`                   // Example: 0
	Courses                 int64 `json:"courses" url:"courses,omitempty"`                                     // Example: 121
	Sections                int64 `json:"sections" url:"sections,omitempty"`                                   // Example: 278
	Xlists                  int64 `json:"xlists" url:"xlists,omitempty"`                                       // Example: 0
	Users                   int64 `json:"users" url:"users,omitempty"`                                         // Example: 346
	Enrollments             int64 `json:"enrollments" url:"enrollments,omitempty"`                             // Example: 1542
	Groups                  int64 `json:"groups" url:"groups,omitempty"`                                       // Example: 0
	GroupMemberships        int64 `json:"group_memberships" url:"group_memberships,omitempty"`                 // Example: 0
	GradePublishingResults  int64 `json:"grade_publishing_results" url:"grade_publishing_results,omitempty"`   // Example: 0
	BatchCoursesDeleted     int64 `json:"batch_courses_deleted" url:"batch_courses_deleted,omitempty"`         // the number of courses that were removed because they were not included in the batch for batch_mode imports. Only included if courses were deleted.Example: 11
	BatchSectionsDeleted    int64 `json:"batch_sections_deleted" url:"batch_sections_deleted,omitempty"`       // the number of sections that were removed because they were not included in the batch for batch_mode imports. Only included if sections were deleted.Example: 0
	BatchEnrollmentsDeleted int64 `json:"batch_enrollments_deleted" url:"batch_enrollments_deleted,omitempty"` // the number of enrollments that were removed because they were not included in the batch for batch_mode imports. Only included if enrollments were deleted.Example: 150
	ErrorCount              int64 `json:"error_count" url:"error_count,omitempty"`                             // Example: 0
	WarningCount            int64 `json:"warning_count" url:"warning_count,omitempty"`                         // Example: 0
}

func (t *SISImportCounts) HasError() error {
	return nil
}
