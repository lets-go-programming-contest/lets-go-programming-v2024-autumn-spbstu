package models

import "gorm.io/gorm"

type Contact struct {
  gorm.Model
  Name string
  Phone string
}
