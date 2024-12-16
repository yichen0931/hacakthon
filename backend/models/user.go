package models

type User struct {
	UserID       string
	Username     string
	UserPassword string
	Firstname    string
	Lastname     string
	UserRole     string
	Salt         string
}
