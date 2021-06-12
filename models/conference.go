package models

import (
	"time"
)

type Conference struct {
	ID                  int64                  `json:"id"`                    // The id of the conference.Example: 170
	ConferenceType      string                 `json:"conference_type"`       // The type of conference.Example: AdobeConnect
	ConferenceKey       string                 `json:"conference_key"`        // The 3rd party's ID for the conference.Example: abcdjoelisgreatxyz
	Description         string                 `json:"description"`           // The description for the conference.Example: Conference Description
	Duration            int64                  `json:"duration"`              // The expected duration the conference is supposed to last.Example: 60
	EndedAt             time.Time              `json:"ended_at"`              // The date that the conference ended at, null if it hasn't ended.Example: 2013-12-13T17:23:26Z
	StartedAt           time.Time              `json:"started_at"`            // The date the conference started at, null if it hasn't started.Example: 2013-12-12T23:02:17Z
	Title               string                 `json:"title"`                 // The title of the conference.Example: Test conference
	Users               []int64                `json:"users"`                 // Array of user ids that are participants in the conference.Example: 1, 7, 8, 9, 10
	HasAdvancedSettings bool                   `json:"has_advanced_settings"` // True if the conference type has advanced settings..
	LongRunning         bool                   `json:"long_running"`          // If true the conference is long running and has no expected end time.
	UserSettings        string                 `json:"user_settings"`         // A collection of settings specific to the conference type.Example: true
	Recordings          []*ConferenceRecording `json:"recordings"`            // A List of recordings for the conference.
	Url                 string                 `json:"url"`                   // URL for the conference, may be null if the conference type doesn't set it.
	JoinUrl             string                 `json:"join_url"`              // URL to join the conference, may be null if the conference type doesn't set it.
	ContextType         string                 `json:"context_type"`          // The type of this conference's context, typically 'Course' or 'Group'..
	ContextID           int64                  `json:"context_id"`            // The ID of this conference's context..
}

func (t *Conference) HasError() error {
	return nil
}
