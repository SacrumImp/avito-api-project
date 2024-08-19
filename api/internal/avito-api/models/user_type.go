package models

type UserTypesEnum string

const (
	Client    UserTypesEnum = "client"
	Moderator UserTypesEnum = "moderator"
)

type UserType struct {
	Id    int
	Title string
}
