package model

import "github.com/google/uuid"

type Player struct {
	Username string `json:"username"`
	Id       uuid.UUID
}
