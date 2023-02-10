package controllers

import "net/http"

// CriarUsuario insere um usuário no Banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando Usuário!"))
}

// BuscarUsuarios busca todos os usuários no Banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os usuários!"))
}

// BuscarUsuario busca um usuários no Banco de dados
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando Usuário!"))
}

// AtualizarUsuario atualiza um usuário no Banco de dados
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando Usuário!"))
}

// DeletarUsuario remove um usuário no Banco de dados
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando Usuário!"))
}
