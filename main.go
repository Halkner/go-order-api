package main

import (
	"database/sql"
	"fmt"

	"github.com/Halkner/go-order-api.git/internal/infra/database"
	use_case "github.com/Halkner/go-order-api.git/internal/use-case"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}

	orderRepository := database.NewOrderRepository(db)

	usecase := use_case.NewCalculateFinalPrice(orderRepository)

	input := use_case.OrderInput{
		ID:    "1",
		Price: 10.0,
		Tax:   1.0,
	}
	output, err := usecase.Execute(input)
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}