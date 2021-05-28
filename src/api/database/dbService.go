package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Order struct {
	Order_id    int    `json:"order_id"`
	Username    string `json:"username"`
	Order_date  string `json:"order_date"`
	Size        string `json:"size"`
	Coffee_type string `json:"type"`
	Num_milk    int    `json:"num_milk"`
	Num_cream   int    `json:"num_cream"`
	Num_sugar   int    `json:"num_sugar"`
}

func Connect() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading the .env file")
	}
	connect := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	db, err := sql.Open("mysql", connect)
	if err != nil {
		log.Fatal("Error connecting to DB")
	}
	// defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Error Unable to ping DB ...")
	}
	fmt.Println("CONNECTED SUCCESFULLY")
	return db

}

func GetOrderHistory(db *sql.DB) []Order {
	results, err := db.Query("SELECT * FROM orders;")
	if err != nil {
		log.Fatal("Error getting result", err)
	}

	var orders []Order

	for results.Next() {
		var order Order
		err = results.Scan(&order.Order_id, &order.Username, &order.Order_date, &order.Size, &order.Coffee_type, &order.Num_milk, &order.Num_cream, &order.Num_sugar)
		if err != nil {
			log.Fatal("Error parsing row", err)
		}
		orders = append(orders, order)
	}
	return orders
}

func CreateOrder(orders []Order, db *sql.DB) {
	var datetime = time.Now().Format(time.RFC3339)

	var query string
	username := "login"
	for _, order := range orders {
		query = fmt.Sprintf("INSERT INTO orders  (username, size, type, num_milk, num_cream, num_sugar, order_date) VALUES ('%s', '%s', '%s', %d, %d, %d, '%s');", username, order.Size, order.Coffee_type, order.Num_milk, order.Num_cream, order.Num_sugar, datetime)
		fmt.Println(query)
		_, err := db.Query(query)
		if err != nil {
			panic(err)
		}
	}
}
