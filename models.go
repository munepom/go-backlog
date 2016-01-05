package gobacklog

import (
	"time"
)

// Issue represents
// +gen * slice:"Where,Count,SortBy,GroupBy[string]"
type Issue struct {
	ID          *int      `json:"id,omitempty"`
	ProjectID   *int      `json:"projectId,omitempty"`
	IssueKey    *string   `json:"issueKey,omitempty"`
	KeyID       *int      `json:"keyId,omitempty"`
	IssueType   IssueType `json:"issueType,omitempty"`
	Summary     *string   `json:"summary,omitempty"`
	Description *string   `json:"description,omitempty"`
	// Resolution interface{} `json:"resolution,omitempty"`
	// Priority interface{} `json:"priority,omitempty"`
	// Status interface{} `json:"status,omitempty"`
	// Assignee interface{} `json:"assignee,omitempty"`
	// Category interface{} `json:"category,omitempty"`
	// Versions interface{} `json:"versions,omitempty"`
	// Milestone interface{} `json:"milestone,omitempty"`
	StartDate *time.Time `json:"startDate,omitempty"`
	DueDate   *time.Time `json:"dueDate,omitempty"`
	// EstimatedHours interface{} `json:"estimatedHours,omitempty"`
	// ActualHours interface{} `json:"actualHours,omitempty"`
	ParentIssueID *int `json:"parentIssueId,omitempty"`
	// CreatedUser interface{} `json:"createdUser,omitempty"`
	Created *time.Time `json:"created,omitempty"`

	// UpdatedUser interface{} `json:"updatedUser,omitempty"`
	Updated *time.Time `json:"updated,omitempty"`
	// CustomFields []interface{} `json:"customFields,omitempty"`
	// Attachments interface{} `json:"attachments,omitempty"`
	// SharedFiles interface{} `json:"sharedFiles,omitempty"`
	// Stars interface{} `json:"stars,omitempty"`
}

// Issues represents
type Issues struct {
	Issues IssueSlice
}

// IssueType represents
type IssueType struct {
	ID           *int    `json:"id,omitempty"`
	ProjectID    *int    `json:"projectId,omitempty"`
	Name         *string `json:"name,omitempty"`
	Color        *string `json:"color,omitempty"`
	DisplayOrder *int    `json:"displayOrder,omitempty"`
}