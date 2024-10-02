package main

import (
	"fmt"
	"net/http"

	"github.com/matheuslimasr/buscaCepGolang/controllers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", controllers.HandlerBuscaCep)

	fmt.Println("Servidor Iniciado!")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
