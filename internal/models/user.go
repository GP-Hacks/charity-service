package models

type UserStatus int

const (
	Admin   UserStatus = iota
	Default UserStatus = iota
)

type User struct {
	ID       string
	Username string
	Email    string
	Status   UserStatus
}
