package model

type Account struct {
	Name     string `gorm:"notNull;size:60"`
	BankName string `gorm:"notNull;size:40"`
	Number   string `gorm:"notNull;size:20"`
}
