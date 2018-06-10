package main

import (
	"time"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Order struct {
	gorm.Model
	OrderedAt time.Time `json:"ordered_at"`
	Email string `json:"email"`
}
