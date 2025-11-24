package user

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Tasks    []Task `json:"tasks"`
}

type Task struct {
	gorm.Model
	WhatIsTheTask string `json:"what_is_the_task" gorm:"column:what_is_the_task"`
	IsDone        bool   `json:"is_done" gorm:"column:is_done"`
	UserID        uint   `json:"user_id" gorm:"column:user_id;not null"`
}
