package routes

import (
	"net/http"
	"LOJA.WEB/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
}
