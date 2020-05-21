package main

import (
	"net/http"

	"github.com/pedrohti/go-loja/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
