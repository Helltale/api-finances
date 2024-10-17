package main

import (
	"fmt"
	"net/http"

	"github.com/helltale/api-finances/internal/handlers"
	"github.com/helltale/api-finances/internal/routers"
)

func main() {

	handlers.Init()
	routers.Init()

	fmt.Println("server start on localhost:8080")
	http.ListenAndServe(":8080", nil)
}
