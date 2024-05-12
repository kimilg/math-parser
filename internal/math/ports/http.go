package ports

import (
	"encoding/json"
	"math-parser/internal/math/app"
	"math-parser/internal/math/app/command"
	"math-parser/internal/math/domain/formula"
	"net/http"
	"strings"
)

type HttpServer struct {
	app app.Application
}

func NewHttpServer(app app.Application) HttpServer {
	return HttpServer{app}
}

func (s *HttpServer) Parse(w http.ResponseWriter, r *http.Request) {
	form := &formula.Equation{}
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	value := strings.TrimSpace(form.Value)
	
	err := s.app.Commands.Parse.Handle(r.Context(), command.Parse{Equation: value})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	
	println("***");
	println("value : " + form.Value)
	println("***");
	
	w.WriteHeader(http.StatusOK)
}

func (s *HttpServer) spreadRandomField(w http.ResponseWriter, r *http.Request) {
	
}

func (s *HttpServer) AddVariable(w http.ResponseWriter, r *http.Request) {
	variable := &formula.Variable{}
	if err := json.NewDecoder(r.Body).Decode(variable); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	println("***");
	println("variable : " + variable.Name)
	println("description : " + variable.Description)
	println("***");


	w.WriteHeader(http.StatusOK)
}

func (s *HttpServer) AddConstant(w http.ResponseWriter, r *http.Request) {
	form := &formula.ConstantForm{}
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	println("***");
	println("value : " + form.Value)
	println("description : " + form.Description)
	println("***");

	w.WriteHeader(http.StatusOK)
}
