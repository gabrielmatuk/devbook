package controllers

import (
	"app/src/respostas"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// FazerLogin utiliza o e-mail e senha do usuário para autenticar o usuário
func FazerLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	usuario, erro := json.Marshal(map[string]string{
		"email": r.FormValue("email"),
		"senha": r.FormValue("senha"),
	})
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	response, erro := http.Post("http://localhost:5000/login", "application/json", bytes.NewBuffer(usuario))
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	token, _ := io.ReadAll(response.Body)

	fmt.Println(response.StatusCode, string(token))
}