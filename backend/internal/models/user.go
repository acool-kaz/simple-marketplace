package models

import "github.com/lib/pq"

type User struct {
	Id       int            `json:"id" db:"id"`
	Name     string         `json:"name" db:"name"`
	Username string         `json:"username" db:"username"`
	City     string         `json:"city" db:"city"`
	Street   string         `json:"street" db:"street"`
	CardNums string         `json:"card_nums" db:"card_nums"`
	CardMY   pq.StringArray `json:"card_m_y" db:"card_m_y"`
	Password string         `json:"password" db:"password"`
}
