package ports

import (
	"encoding/json"
	"math-parser/internal/math/app"
	"math-parser/internal/math/app/command"
	"math-parser/internal/math/domain/formula"
	"net/http"
	"strconv"
	"strings"
)

type HttpServer struct {
	app app.Application
}

func NewHttpServer(app app.Application) HttpServer {
	return HttpServer{app}
}

func (s *HttpServer) Parse(w http.ResponseWriter, r *http.Request) {
	equation := &formula.Equation{}
	if err := json.NewDecoder(r.Body).Decode(equation); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	equation.Value = strings.TrimSpace(equation.Value)

	err := s.app.Commands.Parse.Handle(r.Context(), command.Parse{Equation: equation})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	println("***")
	println("value : " + equation.Value)
	println("***")

	w.WriteHeader(http.StatusOK)
}

func (s *HttpServer) AddConstant(w http.ResponseWriter, r *http.Request) {
	constants := &formula.EquationConstants{}
	if err := json.NewDecoder(r.Body).Decode(constants); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := s.app.Commands.SetConstants.Handle(r.Context(), command.SetConstant{EquationConstants: constants})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	println("***")
	println("equation : " + constants.Equation)
	println("name : " + constants.Variables[0].Name)
	println("category : " + constants.Variables[0].Category)
	println("value: " + strconv.FormatFloat(constants.Variables[0].Constant.GetSize(), 'f', -1, 32))
	println("***")

	w.WriteHeader(http.StatusOK)
}

func (s *HttpServer) spreadRandomField(w http.ResponseWriter, r *http.Request) {

}
