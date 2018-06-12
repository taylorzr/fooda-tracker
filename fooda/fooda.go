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
		User User
		UserId int
	}

	Run struct {
		gorm.Model
	}

	User struct {
		gorm.Model
		Email string `json:"email"`
	}
)

func Forgetters(users []User, orders []Order) []User {
	// orderers := []User{}

	// TODO: Should probably just do this in sql
	// Find users who don't have an order for today
	// Would be an opportunity to practice SQL
	for _, _ = range orders {
		// orders = append(orderers, order.User)
		// Check if they have ordered today
		// If not save to list
	}

	// check all users vs orderers

	// return orderers
	return users
}
