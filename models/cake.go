package models

import "time"

type Cake struct {
	ID          uint64    `gorm:"primary_key:auto_increment" json:"id"`
	Title       string    `gorm:"type:varchar(255)" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	Rating      string    `gorm:"type:varchar(255)" json:"rating"`
	Image       string    `gorm:"type:varchar(255)" json:"image"`
	Created_At  time.Time `gorm:"autoCreateTime:true" json="created_at"`
	Update_dAt  time.Time `gorm:"autoUpdateTime:true" json="updated_at"`
}
