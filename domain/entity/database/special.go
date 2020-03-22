package database

import "time"

type (
	Special struct {
		ID       uint64
		Title    string
		Created  time.Time
		Modified time.Time
	}
)
