package model

import (
	"time"
)

type User struct {
	id         int64
	email      string
	password   string
	name       string
	created_at int64
}

type Project struct {
	id           int64
	project_name string
	created_by   int64
	created_at   time.Time
}

type Ticket struct {
	id          int64
	ticket_no   int64
	assignee_id int64
	reporter_id int64
	project_id  int64
	description string
	status      string
	priority    string
	created_at  time.Time
}

type Comment struct {
	id          int64
	description string
	ticket_id   int64
}
