package models

type StatusesEnum = string

const (
	Created  StatusesEnum = "created"
	Approved StatusesEnum = "approved"
)

type Status struct {
	Id    int
	Title string
}
