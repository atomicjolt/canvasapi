package models

type SISImportStatistic struct {
	Created     int64 `json:"created" url:"created,omitempty"`         // This is the number of items that were created..Example: 18
	Concluded   int64 `json:"concluded" url:"concluded,omitempty"`     // This is the number of items that marked as completed. This only applies to courses and enrollments..Example: 3
	Deactivated int64 `json:"deactivated" url:"deactivated,omitempty"` // This is the number of Enrollments that were marked as 'inactive'. This only applies to enrollments..Example: 1
	Restored    int64 `json:"restored" url:"restored,omitempty"`       // This is the number of items that were set to an active state from a completed, inactive, or deleted state..Example: 2
	Deleted     int64 `json:"deleted" url:"deleted,omitempty"`         // This is the number of items that were deleted..Example: 40
}

func (t *SISImportStatistic) HasErrors() error {
	return nil
}
