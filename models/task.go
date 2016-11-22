package models

import "time"

/*
	A Task
*/

type Task struct { /* using a custom name */
	ID        int       `schema:"-"`
	Content   string    `schema:"content"`
	Date      time.Time `schema:"date"`
	Status    string    `schema:"status"`
	Type      string    `schema:"type"`
	Signifier string    `schema:"signifier"`
	CreatedAt time.Time `schema:"-"`
	UpdatedAt time.Time `schema:"-"`
}
