package models

type Users struct {
	Id       int64
	Name     string
	Surname  string
	Age      int64
	Gender   string
	Login    string
	Password string
	Remove   bool
}