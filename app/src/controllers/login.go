package controllers

import (
	"app/src/utils"
	"net/http"
)

// CarregarTelaDeLogin vai renderizar tela de login
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "login.html", nil)
}
