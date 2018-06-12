package db

import(
	"time"

  "github.com/jinzhu/gorm"

	"github.com/taylorzr/fooda-tracker/fooda"
)

var db *gorm.DB

func init() {
	var err error

	db, err = gorm.Open("sqlite3", "fooda-tracker.db")

	if err != nil {
		panic("failed to connect database")
	}

	// defer db.Close()

	// Migrate the schema
	db.AutoMigrate(
		&fooda.Order{},
		&fooda.User{},
	)
}

func GetOrder(email string) fooda.Order {
	var order fooda.Order

	db.First(&order, "email = ?", email)

	return order
}

func CreateOrder(email string, orderedAt time.Time) fooda.Order {
	user := findOrCreateUser(email)

	// TODO: Correct syntax?!?
	order := fooda.Order{ OrderedAt: orderedAt, User: user }
	db.Create(&order)

	return order
}

func LastRun() time.Time {
	var run fooda.Run

	// TODO: Is this a valid method?
	db.Last(&run)

	return run.CreatedAt
}

func TodaysOrders() []fooda.Order {
	var orders []fooda.Order

	// TODO: Find orders for today only!!!
	db.Find(&orders)

	return orders
}

func AllUsers() []fooda.User {
	var users []fooda.User

	db.Find(&users)

	return users
}

func findOrCreateUser(email string) fooda.User {
	var user fooda.User

	db.First(&user, "email = ?", email)

	if user.ID == 0 {
		user = fooda.User{ Email: email }
		db.Create(&user)
	}

	return user
}
