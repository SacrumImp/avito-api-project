package models

type StatusesEnum = string

const (
	Created StatusesEnum = "created"
)

type Status struct {
	Id    int
	Title string
}
