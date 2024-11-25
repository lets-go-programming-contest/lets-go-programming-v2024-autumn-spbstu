package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	Name  string `json:"name" validate:"required"`
  Phone string `json:"phone" validate:"required,e164"`
}
