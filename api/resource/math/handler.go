package math

import (
	"encoding/json"
	"net/http"
)

type API struct {
	
}

func New() *API {
	return &API{}
}

func (a *API) Parse(w http.ResponseWriter, r *http.Request) {
	
	form := &Form{}
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return;
	}
	
	println("***");
	println("equation : " + form.Equation)
	println("***");
	
	
	w.WriteHeader(http.StatusOK)
}