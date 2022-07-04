package model

import "gorm.io/gorm"

type Status string

const (
	TRANSFERRED Status = "TRANSFERRED"
	IN_PROCESS  Status = "IN_PROCESS"
)

type Donation struct {
	gorm.Model
	CampaignID       uint
	Amount           float64 `gorm:"notNull;check:amount >= 1.0"`
	DonorName        string  `gorm:"notNull;size:40"`
	ReceiptRequested bool    `gorm:"notNull"`
	Status           Status  `gorm:"notNull;type:ENUM('TRANSFERRED','IN_PROCESS')"`
	Account          Account `gorm:"embedded;embeddedPrefix:account_"`
}
