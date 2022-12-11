package models

import "time"

type Author struct {
	Id         string     `json:"id"`
	Fullname   string     `json:"fullname" binding:"required" minLength:"2" maxLength:"255" example:"John Doe"`
	Created_at time.Time  `json:"created_at"`
	Updated_at *time.Time `json:"updated_at"`
	Deleted_at *time.Time `json:"deleted_at"`
}

type CreateAuthorModel struct {
	Fullname   string     `json:"fullname" binding:"required" minLength:"2" maxLength:"255" example:"John Doe"`
}

type UpdateAuthorModel struct {
	Id         string `json:"id" binding:"required"`
	Fullname   string `json:"fullname" minLength:"2" maxLength:"255" example:"John Doe"`
}
