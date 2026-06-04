package models

import "time"

type Task struct {
	UUID        string    `json:"uuid"`
	Title       string    `json:"title"`
	Type        string    `json:"type"`
	DaysInWeek  []string  `json:"days_in_week"`
	OneTimeTask bool      `json:"one_time_task"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CompletedTask struct {
	UUID        string    `json:"uuid"`
	TaskUUID    string    `json:"task_uuid"`
	CompletedAt time.Time `json:"completed_at"`
}