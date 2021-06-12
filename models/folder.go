package models

import (
	"time"
)

type Folder struct {
	ContextType    string    `json:"context_type"`     // Example: Course
	ContextID      int64     `json:"context_id"`       // Example: 1401
	FilesCount     int64     `json:"files_count"`      // Example: 0
	Position       int64     `json:"position"`         // Example: 3
	UpdatedAt      time.Time `json:"updated_at"`       // Example: 2012-07-06T14:58:50Z
	FoldersUrl     string    `json:"folders_url"`      // Example: https://www.example.com/api/v1/folders/2937/folders
	FilesUrl       string    `json:"files_url"`        // Example: https://www.example.com/api/v1/folders/2937/files
	FullName       string    `json:"full_name"`        // Example: course files/11folder
	LockAt         time.Time `json:"lock_at"`          // Example: 2012-07-06T14:58:50Z
	ID             int64     `json:"id"`               // Example: 2937
	FoldersCount   int64     `json:"folders_count"`    // Example: 0
	Name           string    `json:"name"`             // Example: 11folder
	ParentFolderID int64     `json:"parent_folder_id"` // Example: 2934
	CreatedAt      time.Time `json:"created_at"`       // Example: 2012-07-06T14:58:50Z
	UnlockAt       time.Time `json:"unlock_at"`        //
	Hidden         bool      `json:"hidden"`           //
	HiddenForUser  bool      `json:"hidden_for_user"`  //
	Locked         bool      `json:"locked"`           // Example: true
	LockedForUser  bool      `json:"locked_for_user"`  //
	ForSubmissions bool      `json:"for_submissions"`  // If true, indicates this is a read-only folder containing files submitted to assignments.
}

func (t *Folder) HasError() error {
	return nil
}
