package controllers

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/pedrohti/go-loja/errors"
	"github.com/pedrohti/go-loja/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos()

	temp.ExecuteTemplate(w, "Index", todosOsProdutos)
}

func Cadastro(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Cadastro", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		errors.CheckErrorMsg(err, "Erro ao converter o preço")

		quantidadeConvertida, err := strconv.Atoi(quantidade)
		errors.CheckErrorMsg(err, "Erro ao converter a quantidade")

		models.CriarNovoProduto(nome, descricao, precoConvertido, quantidadeConvertida)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	models.DeletaProduto(idDoProduto)
	http.Redirect(w, r, "/", 301)
}

func Editar(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	produto := models.EditaProduto(idDoProduto)
	temp.ExecuteTemplate(w, "Editar", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco, err := strconv.ParseFloat((r.FormValue("preco")), 64)
		errors.CheckErrorMsg(err, "Erro ao converter o preço")

		quantidade, err := strconv.Atoi(r.FormValue("quantidade"))
		errors.CheckErrorMsg(err, "Erro ao converter a quantidade")

		models.AtualizaProduto(id, nome, descricao, preco, quantidade)

	}

	http.Redirect(w, r, "/", 301)
}

//301 = tudo deu certo
