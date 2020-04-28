package database

import "time"

type (
	VirtualWriter struct {
		Id           uint64
		Name         string
		Profile      string
		ImgPath      string
		Expert       int
		IntroduceUrl string
		Job          string
		Created      time.Time
		Modified     time.Time
	}
)
