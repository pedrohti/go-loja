package models

import (
	"github.com/pedrohti/go-loja/db"
	"github.com/pedrohti/go-loja/errors"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {
	db := db.ConectaComBancoDeDados()

	selectProdutos, err := db.Query("SELECT * FROM Produtos ORDER BY id asc")
	errors.CheckError(err)

	p := Produto{}
	produtos := []Produto{}

	for selectProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		errors.CheckError(err)

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)

	}
	defer db.Close()
	return produtos
}

func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()

	insereDadosNoBanco, err := db.Prepare("INSERT INTO Produtos(nome, descricao, preco, quantidade) VALUES ($1, $2, $3, $4)")
	errors.CheckError(err)

	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)

	defer db.Close()
}

func DeletaProduto(id string) {
	db := db.ConectaComBancoDeDados()

	deletarOProduto, err := db.Prepare("DELETE FROM Produtos WHERE id=$1")
	errors.CheckError(err)

	deletarOProduto.Exec(id)

	defer db.Close()
}

func EditaProduto(id string) Produto {
	db := db.ConectaComBancoDeDados()

	produtoDoBanco, err := db.Query("SELECT * FROM Produtos WHERE id=$1", id)
	errors.CheckError(err)

	produtoParaAtualizar := Produto{}

	for produtoDoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoDoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
		errors.CheckError(err)

		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Descricao = descricao
		produtoParaAtualizar.Preco = preco
		produtoParaAtualizar.Quantidade = quantidade
	}
	defer db.Close()
	return produtoParaAtualizar
}

func AtualizaProduto(id string, nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()

	atualizaProduto, err := db.Prepare("UPDATE Produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 WHERE id=$5")
	errors.CheckError(err)

	atualizaProduto.Exec(nome, descricao, preco, quantidade, id)

	defer db.Close()
}
