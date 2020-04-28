package database

import "time"

type (
	Special struct {
		Id                   uint64
		AssignId             uint64
		VirtualWriterId      uint64
		ExpertWriterId       string
		ManualAdd            int
		GroupId              uint64
		Nodeid               uint64
		MdbsaiId             uint64
		PageName             string
		Pankuzu              string
		Title                string
		MetaDescription      string
		MainVisual           string
		PageMainVisual       string
		H1                   string
		Contents             string
		RankingPosition      string
		RankingTitle         string
		RankingContent       string
		RankingContentBottom string
		SpecialEnabled       int
		SelfPv               uint64
		Last30daysPv         uint64 `gorm:"column:last_30days_pv"`
		Status               uint64
		IsForFeed            int
		Created              time.Time
		Modified             time.Time
		PriceModified        time.Time
		RewriteModified      time.Time
		ModifiedAutoRanking  time.Time
		LastStatusModified   time.Time
		Display              int
	}
)
