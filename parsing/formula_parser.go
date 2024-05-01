// Code generated from Formula.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Formula
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type FormulaParser struct {
	*antlr.BaseParser
}

var FormulaParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func formulaParserInit() {
	staticData := &FormulaParserStaticData
	staticData.LiteralNames = []string{
		"", "'frac'", "','", "'.'", "';'", "'('", "')'", "'{'", "'}'", "'_'",
		"'^'", "'\\'", "'+'", "'-'", "'*'", "'/'", "'**'", "'='",
	}
	staticData.SymbolicNames = []string{
		"", "", "COMMA", "DOT", "SEMICOLON", "OPENPAREN", "CLOSEPAREN", "OPENCURLY",
		"CLOSECURLY", "SUBSCRIPT", "SUPERSCRIPT", "BACKSLASH", "PLUS", "MINUS",
		"MULT", "DIV", "POW", "EQUAL", "ID", "LINE_COMMENT", "COMMENT", "INTLIT",
		"FLOATLIT", "WS",
	}
	staticData.RuleNames = []string{
		"equation", "expr", "binaryOperator", "fraction", "constant", "variable",
		"subscriptTail", "argumentTail", "argumentList", "argumentListTail",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 23, 111, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 33,
		8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 44, 8,
		1, 10, 1, 12, 1, 47, 9, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 3, 5, 71, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 86, 8, 6, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 99, 8, 7, 1, 8, 1, 8,
		1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 109, 8, 9, 1, 9, 0, 1, 2, 10,
		0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 0, 4, 1, 0, 14, 15, 1, 0, 12, 13, 1,
		0, 12, 17, 1, 0, 21, 22, 114, 0, 20, 1, 0, 0, 0, 2, 32, 1, 0, 0, 0, 4,
		48, 1, 0, 0, 0, 6, 50, 1, 0, 0, 0, 8, 59, 1, 0, 0, 0, 10, 70, 1, 0, 0,
		0, 12, 85, 1, 0, 0, 0, 14, 98, 1, 0, 0, 0, 16, 100, 1, 0, 0, 0, 18, 108,
		1, 0, 0, 0, 20, 21, 3, 2, 1, 0, 21, 22, 5, 17, 0, 0, 22, 23, 3, 2, 1, 0,
		23, 1, 1, 0, 0, 0, 24, 25, 6, 1, -1, 0, 25, 26, 5, 5, 0, 0, 26, 27, 3,
		2, 1, 0, 27, 28, 5, 6, 0, 0, 28, 33, 1, 0, 0, 0, 29, 33, 3, 8, 4, 0, 30,
		33, 3, 10, 5, 0, 31, 33, 3, 6, 3, 0, 32, 24, 1, 0, 0, 0, 32, 29, 1, 0,
		0, 0, 32, 30, 1, 0, 0, 0, 32, 31, 1, 0, 0, 0, 33, 45, 1, 0, 0, 0, 34, 35,
		10, 6, 0, 0, 35, 36, 5, 16, 0, 0, 36, 44, 3, 2, 1, 6, 37, 38, 10, 5, 0,
		0, 38, 39, 7, 0, 0, 0, 39, 44, 3, 2, 1, 6, 40, 41, 10, 4, 0, 0, 41, 42,
		7, 1, 0, 0, 42, 44, 3, 2, 1, 5, 43, 34, 1, 0, 0, 0, 43, 37, 1, 0, 0, 0,
		43, 40, 1, 0, 0, 0, 44, 47, 1, 0, 0, 0, 45, 43, 1, 0, 0, 0, 45, 46, 1,
		0, 0, 0, 46, 3, 1, 0, 0, 0, 47, 45, 1, 0, 0, 0, 48, 49, 7, 2, 0, 0, 49,
		5, 1, 0, 0, 0, 50, 51, 5, 11, 0, 0, 51, 52, 5, 1, 0, 0, 52, 53, 5, 7, 0,
		0, 53, 54, 3, 2, 1, 0, 54, 55, 5, 8, 0, 0, 55, 56, 5, 7, 0, 0, 56, 57,
		3, 2, 1, 0, 57, 58, 5, 8, 0, 0, 58, 7, 1, 0, 0, 0, 59, 60, 7, 3, 0, 0,
		60, 9, 1, 0, 0, 0, 61, 62, 5, 18, 0, 0, 62, 63, 3, 12, 6, 0, 63, 64, 3,
		14, 7, 0, 64, 71, 1, 0, 0, 0, 65, 66, 5, 11, 0, 0, 66, 67, 5, 18, 0, 0,
		67, 68, 3, 12, 6, 0, 68, 69, 3, 14, 7, 0, 69, 71, 1, 0, 0, 0, 70, 61, 1,
		0, 0, 0, 70, 65, 1, 0, 0, 0, 71, 11, 1, 0, 0, 0, 72, 73, 5, 9, 0, 0, 73,
		74, 5, 7, 0, 0, 74, 75, 5, 18, 0, 0, 75, 86, 5, 8, 0, 0, 76, 77, 5, 9,
		0, 0, 77, 78, 5, 7, 0, 0, 78, 79, 5, 21, 0, 0, 79, 86, 5, 8, 0, 0, 80,
		81, 5, 9, 0, 0, 81, 86, 5, 18, 0, 0, 82, 83, 5, 9, 0, 0, 83, 86, 5, 21,
		0, 0, 84, 86, 1, 0, 0, 0, 85, 72, 1, 0, 0, 0, 85, 76, 1, 0, 0, 0, 85, 80,
		1, 0, 0, 0, 85, 82, 1, 0, 0, 0, 85, 84, 1, 0, 0, 0, 86, 13, 1, 0, 0, 0,
		87, 88, 5, 5, 0, 0, 88, 89, 3, 16, 8, 0, 89, 90, 5, 6, 0, 0, 90, 99, 1,
		0, 0, 0, 91, 92, 5, 5, 0, 0, 92, 93, 3, 16, 8, 0, 93, 94, 5, 4, 0, 0, 94,
		95, 3, 16, 8, 0, 95, 96, 5, 6, 0, 0, 96, 99, 1, 0, 0, 0, 97, 99, 1, 0,
		0, 0, 98, 87, 1, 0, 0, 0, 98, 91, 1, 0, 0, 0, 98, 97, 1, 0, 0, 0, 99, 15,
		1, 0, 0, 0, 100, 101, 3, 2, 1, 0, 101, 102, 3, 18, 9, 0, 102, 17, 1, 0,
		0, 0, 103, 104, 5, 2, 0, 0, 104, 105, 3, 2, 1, 0, 105, 106, 3, 18, 9, 0,
		106, 109, 1, 0, 0, 0, 107, 109, 1, 0, 0, 0, 108, 103, 1, 0, 0, 0, 108,
		107, 1, 0, 0, 0, 109, 19, 1, 0, 0, 0, 7, 32, 43, 45, 70, 85, 98, 108,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// FormulaParserInit initializes any static state used to implement FormulaParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewFormulaParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func FormulaParserInit() {
	staticData := &FormulaParserStaticData
	staticData.once.Do(formulaParserInit)
}

// NewFormulaParser produces a new parser instance for the optional input antlr.TokenStream.
func NewFormulaParser(input antlr.TokenStream) *FormulaParser {
	FormulaParserInit()
	this := new(FormulaParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &FormulaParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Formula.g4"

	return this
}

// FormulaParser tokens.
const (
	FormulaParserEOF          = antlr.TokenEOF
	FormulaParserT__0         = 1
	FormulaParserCOMMA        = 2
	FormulaParserDOT          = 3
	FormulaParserSEMICOLON    = 4
	FormulaParserOPENPAREN    = 5
	FormulaParserCLOSEPAREN   = 6
	FormulaParserOPENCURLY    = 7
	FormulaParserCLOSECURLY   = 8
	FormulaParserSUBSCRIPT    = 9
	FormulaParserSUPERSCRIPT  = 10
	FormulaParserBACKSLASH    = 11
	FormulaParserPLUS         = 12
	FormulaParserMINUS        = 13
	FormulaParserMULT         = 14
	FormulaParserDIV          = 15
	FormulaParserPOW          = 16
	FormulaParserEQUAL        = 17
	FormulaParserID           = 18
	FormulaParserLINE_COMMENT = 19
	FormulaParserCOMMENT      = 20
	FormulaParserINTLIT       = 21
	FormulaParserFLOATLIT     = 22
	FormulaParserWS           = 23
)

// FormulaParser rules.
const (
	FormulaParserRULE_equation         = 0
	FormulaParserRULE_expr             = 1
	FormulaParserRULE_binaryOperator   = 2
	FormulaParserRULE_fraction         = 3
	FormulaParserRULE_constant         = 4
	FormulaParserRULE_variable         = 5
	FormulaParserRULE_subscriptTail    = 6
	FormulaParserRULE_argumentTail     = 7
	FormulaParserRULE_argumentList     = 8
	FormulaParserRULE_argumentListTail = 9
)

// IEquationContext is an interface to support dynamic dispatch.
type IEquationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	EQUAL() antlr.TerminalNode

	// IsEquationContext differentiates from other interfaces.
	IsEquationContext()
}

type EquationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEquationContext() *EquationContext {
	var p = new(EquationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_equation
	return p
}

func InitEmptyEquationContext(p *EquationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_equation
}

func (*EquationContext) IsEquationContext() {}

func NewEquationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EquationContext {
	var p = new(EquationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_equation

	return p
}

func (s *EquationContext) GetParser() antlr.Parser { return s.parser }

func (s *EquationContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *EquationContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *EquationContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(FormulaParserEQUAL, 0)
}

func (s *EquationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EquationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EquationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.EnterEquation(s)
	}
}

func (s *EquationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.ExitEquation(s)
	}
}

func (p *FormulaParser) Equation() (localctx IEquationContext) {
	localctx = NewEquationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FormulaParserRULE_equation)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(20)
		p.expr(0)
	}
	{
		p.SetState(21)
		p.Match(FormulaParserEQUAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(22)
		p.expr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPENPAREN() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	CLOSEPAREN() antlr.TerminalNode
	Constant() IConstantContext
	Variable() IVariableContext
	Fraction() IFractionContext
	POW() antlr.TerminalNode
	MULT() antlr.TerminalNode
	DIV() antlr.TerminalNode
	PLUS() antlr.TerminalNode
	MINUS() antlr.TerminalNode

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) OPENPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserOPENPAREN, 0)
}

func (s *ExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) CLOSEPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserCLOSEPAREN, 0)
}

func (s *ExprContext) Constant() IConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *ExprContext) Variable() IVariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *ExprContext) Fraction() IFractionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFractionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFractionContext)
}

func (s *ExprContext) POW() antlr.TerminalNode {
	return s.GetToken(FormulaParserPOW, 0)
}

func (s *ExprContext) MULT() antlr.TerminalNode {
	return s.GetToken(FormulaParserMULT, 0)
}

func (s *ExprContext) DIV() antlr.TerminalNode {
	return s.GetToken(FormulaParserDIV, 0)
}

func (s *ExprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(FormulaParserPLUS, 0)
}

func (s *ExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(FormulaParserMINUS, 0)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *FormulaParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *FormulaParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, FormulaParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(32)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(25)
			p.Match(FormulaParserOPENPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(26)
			p.expr(0)
		}
		{
			p.SetState(27)
			p.Match(FormulaParserCLOSEPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		{
			p.SetState(29)
			p.Constant()
		}

	case 3:
		{
			p.SetState(30)
			p.Variable()
		}

	case 4:
		{
			p.SetState(31)
			p.Fraction()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(43)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, FormulaParserRULE_expr)
				p.SetState(34)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(35)
					p.Match(FormulaParserPOW)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(36)
					p.expr(6)
				}

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, FormulaParserRULE_expr)
				p.SetState(37)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(38)
					_la = p.GetTokenStream().LA(1)

					if !(_la == FormulaParserMULT || _la == FormulaParserDIV) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(39)
					p.expr(6)
				}

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, FormulaParserRULE_expr)
				p.SetState(40)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(41)
					_la = p.GetTokenStream().LA(1)

					if !(_la == FormulaParserPLUS || _la == FormulaParserMINUS) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(42)
					p.expr(5)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(47)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBinaryOperatorContext is an interface to support dynamic dispatch.
type IBinaryOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PLUS() antlr.TerminalNode
	MINUS() antlr.TerminalNode
	MULT() antlr.TerminalNode
	DIV() antlr.TerminalNode
	POW() antlr.TerminalNode
	EQUAL() antlr.TerminalNode

	// IsBinaryOperatorContext differentiates from other interfaces.
	IsBinaryOperatorContext()
}

type BinaryOperatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinaryOperatorContext() *BinaryOperatorContext {
	var p = new(BinaryOperatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_binaryOperator
	return p
}

func InitEmptyBinaryOperatorContext(p *BinaryOperatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_binaryOperator
}

func (*BinaryOperatorContext) IsBinaryOperatorContext() {}

func NewBinaryOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinaryOperatorContext {
	var p = new(BinaryOperatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_binaryOperator

	return p
}

func (s *BinaryOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *BinaryOperatorContext) PLUS() antlr.TerminalNode {
	return s.GetToken(FormulaParserPLUS, 0)
}

func (s *BinaryOperatorContext) MINUS() antlr.TerminalNode {
	return s.GetToken(FormulaParserMINUS, 0)
}

func (s *BinaryOperatorContext) MULT() antlr.TerminalNode {
	return s.GetToken(FormulaParserMULT, 0)
}

func (s *BinaryOperatorContext) DIV() antlr.TerminalNode {
	return s.GetToken(FormulaParserDIV, 0)
}

func (s *BinaryOperatorContext) POW() antlr.TerminalNode {
	return s.GetToken(FormulaParserPOW, 0)
}

func (s *BinaryOperatorContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(FormulaParserEQUAL, 0)
}

func (s *BinaryOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinaryOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.EnterBinaryOperator(s)
	}
}

func (s *BinaryOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.ExitBinaryOperator(s)
	}
}

func (p *FormulaParser) BinaryOperator() (localctx IBinaryOperatorContext) {
	localctx = NewBinaryOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, FormulaParserRULE_binaryOperator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&258048) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFractionContext is an interface to support dynamic dispatch.
type IFractionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BACKSLASH() antlr.TerminalNode
	AllOPENCURLY() []antlr.TerminalNode
	OPENCURLY(i int) antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllCLOSECURLY() []antlr.TerminalNode
	CLOSECURLY(i int) antlr.TerminalNode

	// IsFractionContext differentiates from other interfaces.
	IsFractionContext()
}

type FractionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFractionContext() *FractionContext {
	var p = new(FractionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_fraction
	return p
}

func InitEmptyFractionContext(p *FractionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_fraction
}

func (*FractionContext) IsFractionContext() {}

func NewFractionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FractionContext {
	var p = new(FractionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_fraction

	return p
}

func (s *FractionContext) GetParser() antlr.Parser { return s.parser }

func (s *FractionContext) BACKSLASH() antlr.TerminalNode {
	return s.GetToken(FormulaParserBACKSLASH, 0)
}

func (s *FractionContext) AllOPENCURLY() []antlr.TerminalNode {
	return s.GetTokens(FormulaParserOPENCURLY)
}

func (s *FractionContext) OPENCURLY(i int) antlr.TerminalNode {
	return s.GetToken(FormulaParserOPENCURLY, i)
}

func (s *FractionContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *FractionContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FractionContext) AllCLOSECURLY() []antlr.TerminalNode {
	return s.GetTokens(FormulaParserCLOSECURLY)
}

func (s *FractionContext) CLOSECURLY(i int) antlr.TerminalNode {
	return s.GetToken(FormulaParserCLOSECURLY, i)
}

func (s *FractionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FractionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FractionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.EnterFraction(s)
	}
}

func (s *FractionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.ExitFraction(s)
	}
}

func (p *FormulaParser) Fraction() (localctx IFractionContext) {
	localctx = NewFractionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, FormulaParserRULE_fraction)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(50)
		p.Match(FormulaParserBACKSLASH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(51)
		p.Match(FormulaParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(52)
		p.Match(FormulaParserOPENCURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(53)
		p.expr(0)
	}
	{
		p.SetState(54)
		p.Match(FormulaParserCLOSECURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(55)
		p.Match(FormulaParserOPENCURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(56)
		p.expr(0)
	}
	{
		p.SetState(57)
		p.Match(FormulaParserCLOSECURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INTLIT() antlr.TerminalNode
	FLOATLIT() antlr.TerminalNode

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_constant
	return p
}

func InitEmptyConstantContext(p *ConstantContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_constant
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) INTLIT() antlr.TerminalNode {
	return s.GetToken(FormulaParserINTLIT, 0)
}

func (s *ConstantContext) FLOATLIT() antlr.TerminalNode {
	return s.GetToken(FormulaParserFLOATLIT, 0)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.ExitConstant(s)
	}
}

func (p *FormulaParser) Constant() (localctx IConstantContext) {
	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, FormulaParserRULE_constant)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(59)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FormulaParserINTLIT || _la == FormulaParserFLOATLIT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	SubscriptTail() ISubscriptTailContext
	ArgumentTail() IArgumentTailContext
	BACKSLASH() antlr.TerminalNode

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_variable
	return p
}

func InitEmptyVariableContext(p *VariableContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_variable
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) ID() antlr.TerminalNode {
	return s.GetToken(FormulaParserID, 0)
}

func (s *VariableContext) SubscriptTail() ISubscriptTailContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubscriptTailContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubscriptTailContext)
}

func (s *VariableContext) ArgumentTail() IArgumentTailContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentTailContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentTailContext)
}

func (s *VariableContext) BACKSLASH() antlr.TerminalNode {
	return s.GetToken(FormulaParserBACKSLASH, 0)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (p *FormulaParser) Variable() (localctx IVariableContext) {
	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, FormulaParserRULE_variable)
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FormulaParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(61)
			p.Match(FormulaParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(62)
			p.SubscriptTail()
		}
		{
			p.SetState(63)
			p.ArgumentTail()
		}

	case FormulaParserBACKSLASH:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(65)
			p.Match(FormulaParserBACKSLASH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(66)
			p.Match(FormulaParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(67)
			p.SubscriptTail()
		}
		{
			p.SetState(68)
			p.ArgumentTail()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISubscriptTailContext is an interface to support dynamic dispatch.
type ISubscriptTailContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SUBSCRIPT() antlr.TerminalNode
	OPENCURLY() antlr.TerminalNode
	ID() antlr.TerminalNode
	CLOSECURLY() antlr.TerminalNode
	INTLIT() antlr.TerminalNode

	// IsSubscriptTailContext differentiates from other interfaces.
	IsSubscriptTailContext()
}

type SubscriptTailContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubscriptTailContext() *SubscriptTailContext {
	var p = new(SubscriptTailContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_subscriptTail
	return p
}

func InitEmptySubscriptTailContext(p *SubscriptTailContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_subscriptTail
}

func (*SubscriptTailContext) IsSubscriptTailContext() {}

func NewSubscriptTailContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubscriptTailContext {
	var p = new(SubscriptTailContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_subscriptTail

	return p
}

func (s *SubscriptTailContext) GetParser() antlr.Parser { return s.parser }

func (s *SubscriptTailContext) SUBSCRIPT() antlr.TerminalNode {
	return s.GetToken(FormulaParserSUBSCRIPT, 0)
}

func (s *SubscriptTailContext) OPENCURLY() antlr.TerminalNode {
	return s.GetToken(FormulaParserOPENCURLY, 0)
}

func (s *SubscriptTailContext) ID() antlr.TerminalNode {
	return s.GetToken(FormulaParserID, 0)
}

func (s *SubscriptTailContext) CLOSECURLY() antlr.TerminalNode {
	return s.GetToken(FormulaParserCLOSECURLY, 0)
}

func (s *SubscriptTailContext) INTLIT() antlr.TerminalNode {
	return s.GetToken(FormulaParserINTLIT, 0)
}

func (s *SubscriptTailContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubscriptTailContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubscriptTailContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.EnterSubscriptTail(s)
	}
}

func (s *SubscriptTailContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.ExitSubscriptTail(s)
	}
}

func (p *FormulaParser) SubscriptTail() (localctx ISubscriptTailContext) {
	localctx = NewSubscriptTailContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, FormulaParserRULE_subscriptTail)
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(72)
			p.Match(FormulaParserSUBSCRIPT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(73)
			p.Match(FormulaParserOPENCURLY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(74)
			p.Match(FormulaParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(75)
			p.Match(FormulaParserCLOSECURLY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(76)
			p.Match(FormulaParserSUBSCRIPT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(77)
			p.Match(FormulaParserOPENCURLY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(78)
			p.Match(FormulaParserINTLIT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(79)
			p.Match(FormulaParserCLOSECURLY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(80)
			p.Match(FormulaParserSUBSCRIPT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(81)
			p.Match(FormulaParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(82)
			p.Match(FormulaParserSUBSCRIPT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(83)
			p.Match(FormulaParserINTLIT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgumentTailContext is an interface to support dynamic dispatch.
type IArgumentTailContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPENPAREN() antlr.TerminalNode
	AllArgumentList() []IArgumentListContext
	ArgumentList(i int) IArgumentListContext
	CLOSEPAREN() antlr.TerminalNode
	SEMICOLON() antlr.TerminalNode

	// IsArgumentTailContext differentiates from other interfaces.
	IsArgumentTailContext()
}

type ArgumentTailContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentTailContext() *ArgumentTailContext {
	var p = new(ArgumentTailContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_argumentTail
	return p
}

func InitEmptyArgumentTailContext(p *ArgumentTailContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_argumentTail
}

func (*ArgumentTailContext) IsArgumentTailContext() {}

func NewArgumentTailContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentTailContext {
	var p = new(ArgumentTailContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_argumentTail

	return p
}

func (s *ArgumentTailContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentTailContext) OPENPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserOPENPAREN, 0)
}

func (s *ArgumentTailContext) AllArgumentList() []IArgumentListContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArgumentListContext); ok {
			len++
		}
	}

	tst := make([]IArgumentListContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArgumentListContext); ok {
			tst[i] = t.(IArgumentListContext)
			i++
		}
	}

	return tst
}

func (s *ArgumentTailContext) ArgumentList(i int) IArgumentListContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentListContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentListContext)
}

func (s *ArgumentTailContext) CLOSEPAREN() antlr.TerminalNode {
	return s.GetToken(FormulaParserCLOSEPAREN, 0)
}

func (s *ArgumentTailContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(FormulaParserSEMICOLON, 0)
}

func (s *ArgumentTailContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentTailContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentTailContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.EnterArgumentTail(s)
	}
}

func (s *ArgumentTailContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.ExitArgumentTail(s)
	}
}

func (p *FormulaParser) ArgumentTail() (localctx IArgumentTailContext) {
	localctx = NewArgumentTailContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, FormulaParserRULE_argumentTail)
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(87)
			p.Match(FormulaParserOPENPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(88)
			p.ArgumentList()
		}
		{
			p.SetState(89)
			p.Match(FormulaParserCLOSEPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(91)
			p.Match(FormulaParserOPENPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(92)
			p.ArgumentList()
		}
		{
			p.SetState(93)
			p.Match(FormulaParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(94)
			p.ArgumentList()
		}
		{
			p.SetState(95)
			p.Match(FormulaParserCLOSEPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgumentListContext is an interface to support dynamic dispatch.
type IArgumentListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	ArgumentListTail() IArgumentListTailContext

	// IsArgumentListContext differentiates from other interfaces.
	IsArgumentListContext()
}

type ArgumentListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentListContext() *ArgumentListContext {
	var p = new(ArgumentListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_argumentList
	return p
}

func InitEmptyArgumentListContext(p *ArgumentListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_argumentList
}

func (*ArgumentListContext) IsArgumentListContext() {}

func NewArgumentListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentListContext {
	var p = new(ArgumentListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_argumentList

	return p
}

func (s *ArgumentListContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentListContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ArgumentListContext) ArgumentListTail() IArgumentListTailContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentListTailContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentListTailContext)
}

func (s *ArgumentListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.EnterArgumentList(s)
	}
}

func (s *ArgumentListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.ExitArgumentList(s)
	}
}

func (p *FormulaParser) ArgumentList() (localctx IArgumentListContext) {
	localctx = NewArgumentListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, FormulaParserRULE_argumentList)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(100)
		p.expr(0)
	}
	{
		p.SetState(101)
		p.ArgumentListTail()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgumentListTailContext is an interface to support dynamic dispatch.
type IArgumentListTailContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COMMA() antlr.TerminalNode
	Expr() IExprContext
	ArgumentListTail() IArgumentListTailContext

	// IsArgumentListTailContext differentiates from other interfaces.
	IsArgumentListTailContext()
}

type ArgumentListTailContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentListTailContext() *ArgumentListTailContext {
	var p = new(ArgumentListTailContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_argumentListTail
	return p
}

func InitEmptyArgumentListTailContext(p *ArgumentListTailContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_argumentListTail
}

func (*ArgumentListTailContext) IsArgumentListTailContext() {}

func NewArgumentListTailContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentListTailContext {
	var p = new(ArgumentListTailContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_argumentListTail

	return p
}

func (s *ArgumentListTailContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentListTailContext) COMMA() antlr.TerminalNode {
	return s.GetToken(FormulaParserCOMMA, 0)
}

func (s *ArgumentListTailContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ArgumentListTailContext) ArgumentListTail() IArgumentListTailContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentListTailContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentListTailContext)
}

func (s *ArgumentListTailContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentListTailContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentListTailContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.EnterArgumentListTail(s)
	}
}

func (s *ArgumentListTailContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.ExitArgumentListTail(s)
	}
}

func (p *FormulaParser) ArgumentListTail() (localctx IArgumentListTailContext) {
	localctx = NewArgumentListTailContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, FormulaParserRULE_argumentListTail)
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FormulaParserCOMMA:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(103)
			p.Match(FormulaParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(104)
			p.expr(0)
		}
		{
			p.SetState(105)
			p.ArgumentListTail()
		}

	case FormulaParserSEMICOLON, FormulaParserCLOSEPAREN:
		p.EnterOuterAlt(localctx, 2)

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *FormulaParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *FormulaParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
