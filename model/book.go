package model

import (
	"time"
)

type Book struct {
	ID         int        `json:"id"`
	Title      string     `json:"title"`
	Author     string     `json:"author"`
	Category   string     `json:"category"`
	Status     string     `json:"status"`
	Rating     int        `json:"rating"`
	Review     string     `json:"review"`
	StartedAt  *time.Time `json:"started_at,omitempty"`
	FinishedAt *time.Time `json:"finished_at,omitempty"`
	CreatedAt  time.Time  `json:"created_at"`
}
