package models

import "time"

type Notice struct {
	Id            int       `json:"id" gorm:"primary_key"`
	TypeNotice    int       `json:"type"`
	ClientId      int       `json:"client_id"`
	Containing    string    `json:"containing"`
	Category      int       `json:"category"`
	Cost          int       `json:"cost"`
	Client        User      `json:"client"`
	TimeAvaliable time.Time `json:"time_avaliable"`
	CreatedAt     time.Time
}
