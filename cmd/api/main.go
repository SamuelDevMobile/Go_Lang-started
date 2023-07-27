package main

import (
	// "encoding/json"
	"net/http"
	"github.com/SamuelDevMobile/Go_Lang-started/internal/entitys"
	// "github.com/go-chi/chi/v5"
	// "github.com/go-chi/chi/v5/middleware"
	"github.com/labstack/echo/v4"
)

func main() {
	//chi
	// r := chi.NewRouter()
	// r.Use(middleware.Logger)
	// r.Get("/order", Order)
	// http.ListenAndServe(":8888", r)

	// echo server http
	e := echo.New()
	e.GET("/order", Order)
	e.Logger.Fatal(e.Start(":8888"))
}

func Order(c echo.Context) error {
	order := entitys.Order{
		ID:    "1",
		Price: 10,
		Tax:   1,
	}
	err := order.CalculateFinalPrice()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, order)
}

// func Order(w http.ResponseWriter, r *http.Request) {
// 	order, err := entitys.NewOrder("123", 1000, 10)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode(err.Error())
// 		return
// 	}
// 	order.CalculateFinalPrice()
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(order)
// }
