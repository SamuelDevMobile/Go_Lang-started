package main

import (
	"fmt"
	"database/sql"

	"github.com/SamuelDevMobile/Go_Lang-started/internal/infra/database"
	"github.com/SamuelDevMobile/Go_Lang-started/internal/usecase"
	_ "github.com/mattn/go-sqlite3"
)

type Car struct {
	Model string
	Color string
}

// metodo
func (c Car) Start() {
	println(c.Model + " has been started")
}

func (c *Car) ChangeColor(color string) { // o ponteiro ele é um apontador para um endereco a onde aquele valor esta guardado na memoria
	c.Color = color
	println("New color: " + c.Color)
}

// func (c Car) ChangeColor(color string) {
// 	c.Color = color // duplicando o valor de c.color na memoria, quando termina esse bloco ele é limpado da memoria e essas infos são perdidas
// 	println("New color: " + c.Color)
// }

// funcao
func soma(x, y int) int {
	return x + y
}

func main() {

	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}

	orderRepository := database.NewOrderRepository(db)
	us := usecase.NewCalculateFinalPrice(orderRepository)

	input := usecase.OrderInput{
		ID:    "1234",
		Price: 10.0,
		Tax:   1.0,
	}
	output, err := us.Execute(input)
	if err != nil {
		panic(err)
	}
	fmt.Println(output)

	// order, err := entitys.NewOrder("1", 10, 1)
	// if err != nil {
	// 	println(err.Error())
	// } else {
	// 	println(order.ID)
	// }

	// car := Car{ // declarando e atribuindo a variavel car
	// 	Model: "Ferrari",
	// 	Color: "Red",
	// }
	// car.Start()
	// car.ChangeColor("Blue")
	// println(car.Color)
}
