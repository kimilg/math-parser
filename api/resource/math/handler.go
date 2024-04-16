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
		return
	}
	
	println("***");
	println("equation : " + form.Equation)
	println("classification : " + form.Classification)
	println("***");
	
	
	w.WriteHeader(http.StatusOK)
}

func (a *API) AddVariable(w http.ResponseWriter, r *http.Request) {
	form := &VariableForm{}
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	println("***");
	println("classification : " + form.Classification)
	println("variable : " + form.Variable)
	println("description : " + form.Description)
	println("***");


	w.WriteHeader(http.StatusOK)
}

func (a *API) AddConstant(w http.ResponseWriter, r *http.Request) {
	form := &ConstantForm{}
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	println("***");
	println("classification : " + form.Classification)
	println("constant : " + form.Constant)
	println("value : " + form.Value)
	println("description : " + form.Description)
	println("***");
	
	w.WriteHeader(http.StatusOK)
}
