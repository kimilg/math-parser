package math

import (
	"encoding/json"
	"github.com/antlr4-go/antlr/v4"
	"math-parser/parser"
	"net/http"
	"strings"
)

type API struct {
	
}

func New() *API {
	return &API{}
}

func (a *API) Parse(w http.ResponseWriter, r *http.Request) {
	form := &EquationForm{}
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//apply grammar to parse after removing space.
	equation := strings.TrimSpace(form.Equation)
	
	is := antlr.NewInputStream(equation)
	
	lexer := parser.NewFormulaLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	
	p := parser.NewFormulaParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	//var formulaListener listener.FormulaTreeListener
	//antlr.ParseTreeWalkerDefault.Walk(&formulaListener, p.Equation())

	var formulaVisitor parser.FormulaVisitorImpl
	formulaVisitor.Visit(p.Equation())
	
	
	println("***");
	println("equation : " + form.Equation)
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
