package models

import "time"

type Task struct {
	UUID string `json:"uuid"`
	UserUUID string `json:"user_uuid"`
	Title string `json:"title"`
	Type string `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type HabitTask struct {
	Task
	DaysInWeek []time.Weekday `json:"days_in_week"`
}

type WeeklyAndMonthyTask struct {
	Task
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

type OneTimeTask struct {
	Task
	Date time.Time `json:"date"`
}