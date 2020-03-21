package database

import "time"

type (
	// Admin adminsテーブルの構造体
	Admin struct {
		ID        uint64
		Email     string
		Password  string
		Name      string
		Role      string 
		UnitPrice float64
		Disabled  int 
		Created   time.Time
		Modified  time.Time
	}
)
