package models

type Deal struct {
	PerformerId int    `json:"performer_id" gorm:"primary_key;autoIncrement:false"`
	NoticeId    int    `json:"notice_id" gorm:"primary_key;autoIncrement:false"`
	Approved    bool   `json:"approved"`
	Performer   User   `json:"performer" gorm:"foreignKey:PerformerId"`
	Notice      Notice `json:"notice" gorm:"foreignKey:NoticeId"`
}
