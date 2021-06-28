package models

type SISImportStatistics struct {
	TotalStateChanges    int64               `json:"total_state_changes" url:"total_state_changes,omitempty"`     // This is the total number of items that were changed in the sis import. There are a few caveats that can cause this number to not add up to the individual counts. There are some state changes that happen that have no impact to the object. An example would be changing a course from 'created' to 'claimed'. Both of these would be considered an active course, but would increment this counter. In this example the course would not increment the created or restored counters for course statistic..Example: 382
	Account              *SISImportStatistic `json:"account" url:"account,omitempty"`                             // This contains that statistics for accounts..
	EnrollmentTerm       *SISImportStatistic `json:"enrollment_term" url:"enrollment_term,omitempty"`             // This contains that statistics for terms..
	CommunicationChannel *SISImportStatistic `json:"communication_channel" url:"communication_channel,omitempty"` // This contains that statistics for communication channels. This is an indirect effect from creating or deleting a user..
	AbstractCourse       *SISImportStatistic `json:"abstract_course" url:"abstract_course,omitempty"`             // This contains that statistics for abstract courses..
	Course               *SISImportStatistic `json:"course" url:"course,omitempty"`                               // This contains that statistics for courses..
	CourseSection        *SISImportStatistic `json:"course_section" url:"course_section,omitempty"`               // This contains that statistics for course sections..
	Enrollment           *SISImportStatistic `json:"enrollment" url:"enrollment,omitempty"`                       // This contains that statistics for enrollments..
	GroupCategory        *SISImportStatistic `json:"group_category" url:"group_category,omitempty"`               // This contains that statistics for group categories..
	Group                *SISImportStatistic `json:"group" url:"group,omitempty"`                                 // This contains that statistics for groups..
	GroupMembership      *SISImportStatistic `json:"group_membership" url:"group_membership,omitempty"`           // This contains that statistics for group memberships. This can be a direct impact from the import or indirect from an enrollment being deleted..
	Pseudonym            *SISImportStatistic `json:"pseudonym" url:"pseudonym,omitempty"`                         // This contains that statistics for pseudonyms. Pseudonyms are logins for users, and are the object that ties an enrollment to a user. This would be impacted from the user importer. .
	UserObserver         *SISImportStatistic `json:"user_observer" url:"user_observer,omitempty"`                 // This contains that statistics for user observers..
	AccountUser          *SISImportStatistic `json:"account_user" url:"account_user,omitempty"`                   // This contains that statistics for account users..
}

func (t *SISImportStatistics) HasError() error {
	return nil
}
