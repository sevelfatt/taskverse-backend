package models

import "time"

type Task struct {
	UUID string `json:"uuid"`
	Title string `json:"title"`
	Type string `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type HabitTask struct {
	Task
	DaysInWeek []string `json:"days_in_week"`
}

type WeeklyTask struct {
	Task
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

type MonthlyTask struct {
	Task
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

type OneTimeTask struct {
	Task
	Date time.Time `json:"date"`
}