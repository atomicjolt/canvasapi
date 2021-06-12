package models

import (
	"time"
)

type Enrollment struct {
	ID                                int64     `json:"id"`                                    // The ID of the enrollment..Example: 1
	CourseID                          int64     `json:"course_id"`                             // The unique id of the course..Example: 1
	SISCourseID                       string    `json:"sis_course_id"`                         // The SIS Course ID in which the enrollment is associated. Only displayed if present. This field is only included if the user has permission to view SIS information..Example: SHEL93921
	CourseIntegrationID               string    `json:"course_integration_id"`                 // The Course Integration ID in which the enrollment is associated. This field is only included if the user has permission to view SIS information..Example: SHEL93921
	CourseSectionID                   int64     `json:"course_section_id"`                     // The unique id of the user's section..Example: 1
	SectionIntegrationID              string    `json:"section_integration_id"`                // The Section Integration ID in which the enrollment is associated. This field is only included if the user has permission to view SIS information..Example: SHEL93921
	SISAccountID                      string    `json:"sis_account_id"`                        // The SIS Account ID in which the enrollment is associated. Only displayed if present. This field is only included if the user has permission to view SIS information..Example: SHEL93921
	SISSectionID                      string    `json:"sis_section_id"`                        // The SIS Section ID in which the enrollment is associated. Only displayed if present. This field is only included if the user has permission to view SIS information..Example: SHEL93921
	SISUserID                         string    `json:"sis_user_id"`                           // The SIS User ID in which the enrollment is associated. Only displayed if present. This field is only included if the user has permission to view SIS information..Example: SHEL93921
	EnrollmentState                   string    `json:"enrollment_state"`                      // The state of the user's enrollment in the course..Example: active
	LimitPrivilegesToCourseSection    bool      `json:"limit_privileges_to_course_section"`    // User can only access his or her own course section..Example: true
	SISImportID                       int64     `json:"sis_import_id"`                         // The unique identifier for the SIS import. This field is only included if the user has permission to manage SIS information..Example: 83
	RootAccountID                     int64     `json:"root_account_id"`                       // The unique id of the user's account..Example: 1
	Type                              string    `json:"type"`                                  // The enrollment type. One of 'StudentEnrollment', 'TeacherEnrollment', 'TaEnrollment', 'DesignerEnrollment', 'ObserverEnrollment'..Example: StudentEnrollment
	UserID                            int64     `json:"user_id"`                               // The unique id of the user..Example: 1
	AssociatedUserID                  int64     `json:"associated_user_id"`                    // The unique id of the associated user. Will be null unless type is ObserverEnrollment..
	Role                              string    `json:"role"`                                  // The enrollment role, for course-level permissions. This field will match `type` if the enrollment role has not been customized..Example: StudentEnrollment
	RoleID                            int64     `json:"role_id"`                               // The id of the enrollment role..Example: 1
	CreatedAt                         time.Time `json:"created_at"`                            // The created time of the enrollment, in ISO8601 format..Example: 2012-04-18T23:08:51Z
	UpdatedAt                         time.Time `json:"updated_at"`                            // The updated time of the enrollment, in ISO8601 format..Example: 2012-04-18T23:08:51Z
	StartAt                           time.Time `json:"start_at"`                              // The start time of the enrollment, in ISO8601 format..Example: 2012-04-18T23:08:51Z
	EndAt                             time.Time `json:"end_at"`                                // The end time of the enrollment, in ISO8601 format..Example: 2012-04-18T23:08:51Z
	LastActivityAt                    time.Time `json:"last_activity_at"`                      // The last activity time of the user for the enrollment, in ISO8601 format..Example: 2012-04-18T23:08:51Z
	LastAttendedAt                    time.Time `json:"last_attended_at"`                      // The last attended date of the user for the enrollment in a course, in ISO8601 format..Example: 2012-04-18T23:08:51Z
	TotalActivityTime                 int64     `json:"total_activity_time"`                   // The total activity time of the user for the enrollment, in seconds..Example: 260
	HtmlUrl                           string    `json:"html_url"`                              // The URL to the Canvas web UI page for this course enrollment..Example: https://.
	Grades                            *Grade    `json:"grades"`                                // The URL to the Canvas web UI page containing the grades associated with this enrollment..Example: https://., 35, , 6.67,
	User                              *User     `json:"user"`                                  // A description of the user..Example: 3, Student 1, 1, Student, Stud 1
	OverrideGrade                     string    `json:"override_grade"`                        // The user's override grade for the course..Example: A
	OverrideScore                     float64   `json:"override_score"`                        // The user's override score for the course..Example: 99.99
	UnpostedCurrentGrade              string    `json:"unposted_current_grade"`                // The user's current grade in the class including muted/unposted assignments. Only included if user has permissions to view this grade, typically teachers, TAs, and admins..
	UnpostedFinalGrade                string    `json:"unposted_final_grade"`                  // The user's final grade for the class including muted/unposted assignments. Only included if user has permissions to view this grade, typically teachers, TAs, and admins...
	UnpostedCurrentScore              string    `json:"unposted_current_score"`                // The user's current score in the class including muted/unposted assignments. Only included if user has permissions to view this score, typically teachers, TAs, and admins...
	UnpostedFinalScore                string    `json:"unposted_final_score"`                  // The user's final score for the class including muted/unposted assignments. Only included if user has permissions to view this score, typically teachers, TAs, and admins...
	HasGradingPeriods                 bool      `json:"has_grading_periods"`                   // optional: Indicates whether the course the enrollment belongs to has grading periods set up. (applies only to student enrollments, and only available in course endpoints).Example: true
	TotalsForAllGradingPeriodsOption  bool      `json:"totals_for_all_grading_periods_option"` // optional: Indicates whether the course the enrollment belongs to has the Display Totals for 'All Grading Periods' feature enabled. (applies only to student enrollments, and only available in course endpoints).Example: true
	CurrentGradingPeriodTitle         string    `json:"current_grading_period_title"`          // optional: The name of the currently active grading period, if one exists. If the course the enrollment belongs to does not have grading periods, or if no currently active grading period exists, the value will be null. (applies only to student enrollments, and only available in course endpoints).Example: Fall Grading Period
	CurrentGradingPeriodID            int64     `json:"current_grading_period_id"`             // optional: The id of the currently active grading period, if one exists. If the course the enrollment belongs to does not have grading periods, or if no currently active grading period exists, the value will be null. (applies only to student enrollments, and only available in course endpoints).Example: 5
	CurrentPeriodOverrideGrade        string    `json:"current_period_override_grade"`         // The user's override grade for the current grading period..Example: A
	CurrentPeriodOverrideScore        float64   `json:"current_period_override_score"`         // The user's override score for the current grading period..Example: 99.99
	CurrentPeriodUnpostedCurrentScore float64   `json:"current_period_unposted_current_score"` // optional: The student's score in the course for the current grading period, including muted/unposted assignments. Only included if user has permission to view this score, typically teachers, TAs, and admins. If the course the enrollment belongs to does not have grading periods, or if no currently active grading period exists, the value will be null. (applies only to student enrollments, and only available in course endpoints).Example: 95.8
	CurrentPeriodUnpostedFinalScore   float64   `json:"current_period_unposted_final_score"`   // optional: The student's score in the course for the current grading period, including muted/unposted assignments and including ungraded assignments with a score of 0. Only included if user has permission to view this score, typically teachers, TAs, and admins. If the course the enrollment belongs to does not have grading periods, or if no currently active grading period exists, the value will be null. (applies only to student enrollments, and only available in course endpoints).Example: 85.25
	CurrentPeriodUnpostedCurrentGrade string    `json:"current_period_unposted_current_grade"` // optional: The letter grade equivalent of current_period_unposted_current_score, if available. Only included if user has permission to view this grade, typically teachers, TAs, and admins. If the course the enrollment belongs to does not have grading periods, or if no currently active grading period exists, the value will be null. (applies only to student enrollments, and only available in course endpoints).Example: A
	CurrentPeriodUnpostedFinalGrade   string    `json:"current_period_unposted_final_grade"`   // optional: The letter grade equivalent of current_period_unposted_final_score, if available. Only included if user has permission to view this grade, typically teachers, TAs, and admins. If the course the enrollment belongs to does not have grading periods, or if no currently active grading period exists, the value will be null. (applies only to student enrollments, and only available in course endpoints).Example: B
}

func (t *Enrollment) HasError() error {
	return nil
}
