package main

import (
	"app/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Rodando APP")

	r := router.Gerar()
	log.Fatal(http.ListenAndServe(":3000", r))
}
