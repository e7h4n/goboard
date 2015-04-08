package models

import (
	"time"
)

type Dashboard struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Config    string
	ProjectId int
	Private   bool
	Owner     User
}
