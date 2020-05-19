package routes

import (
	"net/http"

	"github.com/pedrohti/loja/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/cadastro", controllers.Cadastro)

	// Ação
	http.HandleFunc("/insert", controllers.Insert)
}
