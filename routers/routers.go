package routers

import (
	"net/http"

	"github.com/helltale/api-finances/handlers"
)

func Init() {
	http.HandleFunc("/income/all", handlers.GetAllIncomes) //средний ожидаемый совокупный доход
}
