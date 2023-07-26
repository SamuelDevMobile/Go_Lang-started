package main

import (
	"encoding/json"
	"fmt"
	"database/sql"

	"github.com/SamuelDevMobile/Go_Lang-started/internal/infra/database"
	"github.com/SamuelDevMobile/Go_Lang-started/internal/usecase"
	_ "github.com/mattn/go-sqlite3"
	amqp "github.com/rabbitmq/amqp091-go"
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
	defer db.Close() // espera tudo rodar e depois executa o close
	orderRepository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPrice(orderRepository)

	

	// input := usecase.OrderInput{
	// 	ID:    "12334",
	// 	Price: 10.0,
	// 	Tax:   1.0,
	// }
	// output, err := uc.Execute(input)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(output)
}

func rabbitmqWorker(msgChan chan amqp.Delivery, uc usecase.CalculateFinalPrice) {
	fmt.Println("Starting rabbitmq")

	for msg := range msgChan {
		var input usecase.OrderInput
		err := json.Unmarshal(msg.Body, &input)
		if err != nil { 
			panic(err)
		}
		output, err := uc.Execute(input)
		if err != nil {
			panic(err)
		}
		msg.Ack(false) 
		fmt.Println("Mensagem processada e salva no banco: ", output)
	}
}
