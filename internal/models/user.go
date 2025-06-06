package models

import "time"

type UserStatus int

const (
	Admin UserStatus = iota
	Default
)

type User struct {
	ID          int64
	FirstName   string
	LastName    string
	Surname     string
	Email       string
	Status      UserStatus
	DateOfBirth time.Time
}
