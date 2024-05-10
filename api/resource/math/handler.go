package math

import (
	"encoding/json"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"math-parser/parser"
	"net/http"
	"strings"
)

type API struct {
	equationService *EquationService
}

func New(repository EquationRepository) *API {
	return &API{
		equationService: NewEquationService(repository),
	}
}

func (a *API) Parse(w http.ResponseWriter, r *http.Request) {
	form := &Equation{}
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	value := strings.TrimSpace(form.Value)
	
	eq, err := a.equationService.GetFromValue(r.Context(), value)
	if err != nil {
		eq, err = a.equationService.Create(r.Context(), value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	fmt.Printf("equation Id: %d, Value: %s\n", eq.Id, eq.Value)
	equation := a.parseEquation(value)
	
	print(equation.Description)	
	println("***");
	println("value : " + form.Value)
	println("***");
	
	w.WriteHeader(http.StatusOK)
}

func (a *API) AddVariable(w http.ResponseWriter, r *http.Request) {
	variable := &Variable{}
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

func (a *API) AddConstant(w http.ResponseWriter, r *http.Request) {
	form := &ConstantForm{}
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

func (a *API) parseEquation(value string) Expression {
	is := antlr.NewInputStream(value)

	lexer := parser.NewFormulaLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewFormulaParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	var formulaVisitor FormulaVisitorImpl
	eqn := formulaVisitor.Visit(p.Equation())
	if eqn == nil {
		fmt.Errorf("nil value")
	}
	return eqn.(Expression)
}