package db

import(
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
	db.AutoMigrate(&fooda.Order{})
}

func GetOrder(email string) fooda.Order {
	var order fooda.Order

	// db.First(&order, 1)
	db.First(&order, "email = ?", "zach.taylor@avant.com")

	return order
}

func CreateOrder(order *fooda.Order) {
	db.Create(order)
}
