package main

import (
	"net/http"
	"text/template"
)

type Produto struct{
	Nome, Descricao string
	Preco float64 
	Quantidade int 
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)

	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome:"Camisa", Descricao:"Bela", Preco:59, Quantidade:12},
		{"Calça", "Elegante", 99, 20},
		{"Short", "Confortavel", 49, 30},
		{"Brusinha", "Chique", 159, 10},
	}

	temp.ExecuteTemplate(w, "Index", produtos)
}