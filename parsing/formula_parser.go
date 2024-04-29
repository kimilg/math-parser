// Code generated from Formula.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parsing // Formula
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
		"MULT", "DIV", "POW", "EQUAL", "SYMBOL", "ID", "LINE_COMMENT", "COMMENT",
		"INTLIT", "SINGLEINTLIT", "FLOATLIT", "WS",
	}
	staticData.RuleNames = []string{
		"program", "expr", "binaryOperator", "fraction", "constant", "variable",
		"subscriptTail", "argumentTail", "argumentList", "argumentListTail",
		"symbolList",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 25, 119, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 3, 1, 35, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		5, 1, 46, 8, 1, 10, 1, 12, 1, 49, 9, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 73, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 89, 8, 6, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 102,
		8, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 112, 8, 9,
		1, 10, 1, 10, 1, 10, 3, 10, 117, 8, 10, 1, 10, 0, 1, 2, 11, 0, 2, 4, 6,
		8, 10, 12, 14, 16, 18, 20, 0, 4, 1, 0, 14, 15, 1, 0, 12, 13, 1, 0, 12,
		17, 2, 0, 22, 22, 24, 24, 122, 0, 22, 1, 0, 0, 0, 2, 34, 1, 0, 0, 0, 4,
		50, 1, 0, 0, 0, 6, 52, 1, 0, 0, 0, 8, 61, 1, 0, 0, 0, 10, 72, 1, 0, 0,
		0, 12, 88, 1, 0, 0, 0, 14, 101, 1, 0, 0, 0, 16, 103, 1, 0, 0, 0, 18, 111,
		1, 0, 0, 0, 20, 116, 1, 0, 0, 0, 22, 23, 3, 2, 1, 0, 23, 24, 5, 17, 0,
		0, 24, 25, 3, 2, 1, 0, 25, 1, 1, 0, 0, 0, 26, 27, 6, 1, -1, 0, 27, 28,
		5, 5, 0, 0, 28, 29, 3, 2, 1, 0, 29, 30, 5, 6, 0, 0, 30, 35, 1, 0, 0, 0,
		31, 35, 3, 8, 4, 0, 32, 35, 3, 10, 5, 0, 33, 35, 3, 6, 3, 0, 34, 26, 1,
		0, 0, 0, 34, 31, 1, 0, 0, 0, 34, 32, 1, 0, 0, 0, 34, 33, 1, 0, 0, 0, 35,
		47, 1, 0, 0, 0, 36, 37, 10, 6, 0, 0, 37, 38, 5, 16, 0, 0, 38, 46, 3, 2,
		1, 6, 39, 40, 10, 5, 0, 0, 40, 41, 7, 0, 0, 0, 41, 46, 3, 2, 1, 6, 42,
		43, 10, 4, 0, 0, 43, 44, 7, 1, 0, 0, 44, 46, 3, 2, 1, 5, 45, 36, 1, 0,
		0, 0, 45, 39, 1, 0, 0, 0, 45, 42, 1, 0, 0, 0, 46, 49, 1, 0, 0, 0, 47, 45,
		1, 0, 0, 0, 47, 48, 1, 0, 0, 0, 48, 3, 1, 0, 0, 0, 49, 47, 1, 0, 0, 0,
		50, 51, 7, 2, 0, 0, 51, 5, 1, 0, 0, 0, 52, 53, 5, 11, 0, 0, 53, 54, 5,
		1, 0, 0, 54, 55, 5, 7, 0, 0, 55, 56, 3, 2, 1, 0, 56, 57, 5, 8, 0, 0, 57,
		58, 5, 7, 0, 0, 58, 59, 3, 2, 1, 0, 59, 60, 5, 8, 0, 0, 60, 7, 1, 0, 0,
		0, 61, 62, 7, 3, 0, 0, 62, 9, 1, 0, 0, 0, 63, 64, 5, 19, 0, 0, 64, 65,
		3, 12, 6, 0, 65, 66, 3, 14, 7, 0, 66, 73, 1, 0, 0, 0, 67, 68, 5, 11, 0,
		0, 68, 69, 5, 19, 0, 0, 69, 70, 3, 12, 6, 0, 70, 71, 3, 14, 7, 0, 71, 73,
		1, 0, 0, 0, 72, 63, 1, 0, 0, 0, 72, 67, 1, 0, 0, 0, 73, 11, 1, 0, 0, 0,
		74, 75, 5, 9, 0, 0, 75, 76, 5, 7, 0, 0, 76, 77, 3, 20, 10, 0, 77, 78, 5,
		8, 0, 0, 78, 89, 1, 0, 0, 0, 79, 80, 5, 9, 0, 0, 80, 81, 5, 7, 0, 0, 81,
		82, 5, 22, 0, 0, 82, 89, 5, 8, 0, 0, 83, 84, 5, 9, 0, 0, 84, 89, 5, 18,
		0, 0, 85, 86, 5, 9, 0, 0, 86, 89, 5, 23, 0, 0, 87, 89, 1, 0, 0, 0, 88,
		74, 1, 0, 0, 0, 88, 79, 1, 0, 0, 0, 88, 83, 1, 0, 0, 0, 88, 85, 1, 0, 0,
		0, 88, 87, 1, 0, 0, 0, 89, 13, 1, 0, 0, 0, 90, 91, 5, 5, 0, 0, 91, 92,
		3, 16, 8, 0, 92, 93, 5, 6, 0, 0, 93, 102, 1, 0, 0, 0, 94, 95, 5, 5, 0,
		0, 95, 96, 3, 16, 8, 0, 96, 97, 5, 4, 0, 0, 97, 98, 3, 16, 8, 0, 98, 99,
		5, 6, 0, 0, 99, 102, 1, 0, 0, 0, 100, 102, 1, 0, 0, 0, 101, 90, 1, 0, 0,
		0, 101, 94, 1, 0, 0, 0, 101, 100, 1, 0, 0, 0, 102, 15, 1, 0, 0, 0, 103,
		104, 3, 2, 1, 0, 104, 105, 3, 18, 9, 0, 105, 17, 1, 0, 0, 0, 106, 107,
		5, 2, 0, 0, 107, 108, 3, 2, 1, 0, 108, 109, 3, 18, 9, 0, 109, 112, 1, 0,
		0, 0, 110, 112, 1, 0, 0, 0, 111, 106, 1, 0, 0, 0, 111, 110, 1, 0, 0, 0,
		112, 19, 1, 0, 0, 0, 113, 114, 5, 18, 0, 0, 114, 117, 3, 20, 10, 0, 115,
		117, 1, 0, 0, 0, 116, 113, 1, 0, 0, 0, 116, 115, 1, 0, 0, 0, 117, 21, 1,
		0, 0, 0, 8, 34, 45, 47, 72, 88, 101, 111, 116,
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
	FormulaParserSYMBOL       = 18
	FormulaParserID           = 19
	FormulaParserLINE_COMMENT = 20
	FormulaParserCOMMENT      = 21
	FormulaParserINTLIT       = 22
	FormulaParserSINGLEINTLIT = 23
	FormulaParserFLOATLIT     = 24
	FormulaParserWS           = 25
)

// FormulaParser rules.
const (
	FormulaParserRULE_program          = 0
	FormulaParserRULE_expr             = 1
	FormulaParserRULE_binaryOperator   = 2
	FormulaParserRULE_fraction         = 3
	FormulaParserRULE_constant         = 4
	FormulaParserRULE_variable         = 5
	FormulaParserRULE_subscriptTail    = 6
	FormulaParserRULE_argumentTail     = 7
	FormulaParserRULE_argumentList     = 8
	FormulaParserRULE_argumentListTail = 9
	FormulaParserRULE_symbolList       = 10
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	EQUAL() antlr.TerminalNode

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) AllExpr() []IExprContext {
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

func (s *ProgramContext) Expr(i int) IExprContext {
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

func (s *ProgramContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(FormulaParserEQUAL, 0)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *FormulaParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FormulaParserRULE_program)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(22)
		p.expr(0)
	}
	{
		p.SetState(23)
		p.Match(FormulaParserEQUAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(24)
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
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(27)
			p.Match(FormulaParserOPENPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(28)
			p.expr(0)
		}
		{
			p.SetState(29)
			p.Match(FormulaParserCLOSEPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		{
			p.SetState(31)
			p.Constant()
		}

	case 3:
		{
			p.SetState(32)
			p.Variable()
		}

	case 4:
		{
			p.SetState(33)
			p.Fraction()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(47)
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
			p.SetState(45)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, FormulaParserRULE_expr)
				p.SetState(36)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(37)
					p.Match(FormulaParserPOW)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(38)
					p.expr(6)
				}

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, FormulaParserRULE_expr)
				p.SetState(39)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(40)
					_la = p.GetTokenStream().LA(1)

					if !(_la == FormulaParserMULT || _la == FormulaParserDIV) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(41)
					p.expr(6)
				}

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, FormulaParserRULE_expr)
				p.SetState(42)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(43)
					_la = p.GetTokenStream().LA(1)

					if !(_la == FormulaParserPLUS || _la == FormulaParserMINUS) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(44)
					p.expr(5)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(49)
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
		p.SetState(50)
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
		p.SetState(52)
		p.Match(FormulaParserBACKSLASH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(53)
		p.Match(FormulaParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(54)
		p.Match(FormulaParserOPENCURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(55)
		p.expr(0)
	}
	{
		p.SetState(56)
		p.Match(FormulaParserCLOSECURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(57)
		p.Match(FormulaParserOPENCURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(58)
		p.expr(0)
	}
	{
		p.SetState(59)
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
		p.SetState(61)
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
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FormulaParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(63)
			p.Match(FormulaParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(64)
			p.SubscriptTail()
		}
		{
			p.SetState(65)
			p.ArgumentTail()
		}

	case FormulaParserBACKSLASH:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(67)
			p.Match(FormulaParserBACKSLASH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(68)
			p.Match(FormulaParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(69)
			p.SubscriptTail()
		}
		{
			p.SetState(70)
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
	SymbolList() ISymbolListContext
	CLOSECURLY() antlr.TerminalNode
	INTLIT() antlr.TerminalNode
	SYMBOL() antlr.TerminalNode
	SINGLEINTLIT() antlr.TerminalNode

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

func (s *SubscriptTailContext) SymbolList() ISymbolListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISymbolListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISymbolListContext)
}

func (s *SubscriptTailContext) CLOSECURLY() antlr.TerminalNode {
	return s.GetToken(FormulaParserCLOSECURLY, 0)
}

func (s *SubscriptTailContext) INTLIT() antlr.TerminalNode {
	return s.GetToken(FormulaParserINTLIT, 0)
}

func (s *SubscriptTailContext) SYMBOL() antlr.TerminalNode {
	return s.GetToken(FormulaParserSYMBOL, 0)
}

func (s *SubscriptTailContext) SINGLEINTLIT() antlr.TerminalNode {
	return s.GetToken(FormulaParserSINGLEINTLIT, 0)
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
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(74)
			p.Match(FormulaParserSUBSCRIPT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(75)
			p.Match(FormulaParserOPENCURLY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(76)
			p.SymbolList()
		}
		{
			p.SetState(77)
			p.Match(FormulaParserCLOSECURLY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(79)
			p.Match(FormulaParserSUBSCRIPT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(80)
			p.Match(FormulaParserOPENCURLY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(81)
			p.Match(FormulaParserINTLIT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(82)
			p.Match(FormulaParserCLOSECURLY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(83)
			p.Match(FormulaParserSUBSCRIPT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(84)
			p.Match(FormulaParserSYMBOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(85)
			p.Match(FormulaParserSUBSCRIPT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(86)
			p.Match(FormulaParserSINGLEINTLIT)
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
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(90)
			p.Match(FormulaParserOPENPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(91)
			p.ArgumentList()
		}
		{
			p.SetState(92)
			p.Match(FormulaParserCLOSEPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(94)
			p.Match(FormulaParserOPENPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(95)
			p.ArgumentList()
		}
		{
			p.SetState(96)
			p.Match(FormulaParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(97)
			p.ArgumentList()
		}
		{
			p.SetState(98)
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
		p.SetState(103)
		p.expr(0)
	}
	{
		p.SetState(104)
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
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FormulaParserCOMMA:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(106)
			p.Match(FormulaParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(107)
			p.expr(0)
		}
		{
			p.SetState(108)
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

// ISymbolListContext is an interface to support dynamic dispatch.
type ISymbolListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SYMBOL() antlr.TerminalNode
	SymbolList() ISymbolListContext

	// IsSymbolListContext differentiates from other interfaces.
	IsSymbolListContext()
}

type SymbolListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySymbolListContext() *SymbolListContext {
	var p = new(SymbolListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_symbolList
	return p
}

func InitEmptySymbolListContext(p *SymbolListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FormulaParserRULE_symbolList
}

func (*SymbolListContext) IsSymbolListContext() {}

func NewSymbolListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SymbolListContext {
	var p = new(SymbolListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FormulaParserRULE_symbolList

	return p
}

func (s *SymbolListContext) GetParser() antlr.Parser { return s.parser }

func (s *SymbolListContext) SYMBOL() antlr.TerminalNode {
	return s.GetToken(FormulaParserSYMBOL, 0)
}

func (s *SymbolListContext) SymbolList() ISymbolListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISymbolListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISymbolListContext)
}

func (s *SymbolListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SymbolListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SymbolListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.EnterSymbolList(s)
	}
}

func (s *SymbolListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FormulaListener); ok {
		listenerT.ExitSymbolList(s)
	}
}

func (p *FormulaParser) SymbolList() (localctx ISymbolListContext) {
	localctx = NewSymbolListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, FormulaParserRULE_symbolList)
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FormulaParserSYMBOL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(113)
			p.Match(FormulaParserSYMBOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(114)
			p.SymbolList()
		}

	case FormulaParserCLOSECURLY:
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
