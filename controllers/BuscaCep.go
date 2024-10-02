package controllers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/matheuslimasr/buscaCepGolang/Models"
)

func HandlerBuscaCep(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*30)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://viacep.com.br/ws/"+cep+"/json/", nil)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	var ViaCep Models.ViaCep

	err = json.Unmarshal(body, &ViaCep)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(ViaCep)
}
