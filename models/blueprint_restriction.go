package models

type BlueprintRestriction struct {
	Content           bool `json:"content"`            // Restriction on main content (e.g. title, description)..Example: true
	Points            bool `json:"points"`             // Restriction on points possible for assignments and graded learning objects.Example: true
	DueDates          bool `json:"due_dates"`          // Restriction on due dates for assignments and graded learning objects.
	AvailabilityDates bool `json:"availability_dates"` // Restriction on availability dates for an object.Example: true
}

func (t *BlueprintRestriction) HasError() error {
	return nil
}
