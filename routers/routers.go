package routers

import (
	"net/http"

	"github.com/helltale/api-finances/handlers"
)

func Init() {
	http.HandleFunc("/income/all", handlers.GetAllIncomes)                  //все из реального дохода
	http.HandleFunc("/income_expected/all", handlers.GetAllIncomesExpected) //все из ожидаемого дохода
	http.HandleFunc("/account/all", handlers.GetAllAccounts)                //все аккаунты
}
