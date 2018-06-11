package fooda

import (
	"time"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
)

type (
	Order struct {
		gorm.Model
		OrderedAt time.Time `json:"ordered_at"`
		Email string `json:"email"`
	}

	Run struct {
		gorm.Model
	}

	User struct {
		gorm.Model
		Email string `json:"email"`
	}
)

func Forgetters(orders []Order, time time.Time) {
	for _, order := range orders {
		// Check if they have ordered today
		// If not save to list
	}
}
