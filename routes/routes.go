package routes

import (
	"net/http"

	"github.com/pedrohti/loja/controllers"
)

func CarregaRotas() {
	// Paginas
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/cadastro", controllers.Cadastro)
	http.HandleFunc("/editar", controllers.Editar)

	// Ação
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.Delete)
	http.HandleFunc("/update", controllers.Update)
}
