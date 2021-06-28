package models

import (
	"time"
)

type Folder struct {
	ContextType    string    `json:"context_type" url:"context_type,omitempty"`         // Example: Course
	ContextID      int64     `json:"context_id" url:"context_id,omitempty"`             // Example: 1401
	FilesCount     int64     `json:"files_count" url:"files_count,omitempty"`           // Example: 0
	Position       int64     `json:"position" url:"position,omitempty"`                 // Example: 3
	UpdatedAt      time.Time `json:"updated_at" url:"updated_at,omitempty"`             // Example: 2012-07-06T14:58:50Z
	FoldersUrl     string    `json:"folders_url" url:"folders_url,omitempty"`           // Example: https://www.example.com/api/v1/folders/2937/folders
	FilesUrl       string    `json:"files_url" url:"files_url,omitempty"`               // Example: https://www.example.com/api/v1/folders/2937/files
	FullName       string    `json:"full_name" url:"full_name,omitempty"`               // Example: course files/11folder
	LockAt         time.Time `json:"lock_at" url:"lock_at,omitempty"`                   // Example: 2012-07-06T14:58:50Z
	ID             int64     `json:"id" url:"id,omitempty"`                             // Example: 2937
	FoldersCount   int64     `json:"folders_count" url:"folders_count,omitempty"`       // Example: 0
	Name           string    `json:"name" url:"name,omitempty"`                         // Example: 11folder
	ParentFolderID int64     `json:"parent_folder_id" url:"parent_folder_id,omitempty"` // Example: 2934
	CreatedAt      time.Time `json:"created_at" url:"created_at,omitempty"`             // Example: 2012-07-06T14:58:50Z
	UnlockAt       time.Time `json:"unlock_at" url:"unlock_at,omitempty"`               //
	Hidden         bool      `json:"hidden" url:"hidden,omitempty"`                     //
	HiddenForUser  bool      `json:"hidden_for_user" url:"hidden_for_user,omitempty"`   //
	Locked         bool      `json:"locked" url:"locked,omitempty"`                     // Example: true
	LockedForUser  bool      `json:"locked_for_user" url:"locked_for_user,omitempty"`   //
	ForSubmissions bool      `json:"for_submissions" url:"for_submissions,omitempty"`   // If true, indicates this is a read-only folder containing files submitted to assignments.
}

func (t *Folder) HasError() error {
	return nil
}
