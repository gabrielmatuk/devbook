package middlewares

import (
	"app/src/cookies"
	"fmt"
	"log"
	"net/http"
)

// Logger escreve informações da requisição no terminal
func Logger(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		proximaFuncao(w, r)
	}
}

// Autenticar verifica a existência de cookies
func Autenticar(proximaFuncao http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("DESGRAÇAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
		valore, erro := cookies.Ler(r)
		fmt.Println(valore, erro)
		// if _, erro := cookies.Ler(r); erro != nil {
		// 	http.Redirect(w, r, "/login", 302)
		// 	return
		// }
		proximaFuncao(w, r)
	}
}