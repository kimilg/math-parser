// Code generated from Formula.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type FormulaLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var FormulaLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func formulalexerLexerInit() {
	staticData := &FormulaLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'\\frac'", "','", "'.'", "';'", "'('", "')'", "'{'", "'}'", "'_'",
		"'^'", "'\\'", "'+'", "'-'", "'*'", "'/'", "'**'", "'='",
	}
	staticData.SymbolicNames = []string{
		"", "", "COMMA", "DOT", "SEMICOLON", "OPENPAREN", "CLOSEPAREN", "OPENCURLY",
		"CLOSECURLY", "SUBSCRIPT", "SUPERSCRIPT", "BACKSLASH", "PLUS", "MINUS",
		"MULT", "DIV", "POW", "EQUAL", "SINGLEID", "ID", "LINE_COMMENT", "COMMENT",
		"SINGLEINTLIT", "INTLIT", "FLOATLIT", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "COMMA", "DOT", "SEMICOLON", "OPENPAREN", "CLOSEPAREN", "OPENCURLY",
		"CLOSECURLY", "SUBSCRIPT", "SUPERSCRIPT", "BACKSLASH", "PLUS", "MINUS",
		"MULT", "DIV", "POW", "EQUAL", "SINGLEID", "ID", "LINE_COMMENT", "COMMENT",
		"SINGLEINTLIT", "INTLIT", "FLOATLIT", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 25, 159, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4,
		1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1,
		15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 4, 18, 95, 8, 18, 11, 18,
		12, 18, 96, 1, 18, 1, 18, 4, 18, 101, 8, 18, 11, 18, 12, 18, 102, 3, 18,
		105, 8, 18, 1, 19, 1, 19, 1, 19, 1, 19, 5, 19, 111, 8, 19, 10, 19, 12,
		19, 114, 9, 19, 1, 19, 3, 19, 117, 8, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1,
		20, 1, 20, 1, 20, 1, 20, 5, 20, 127, 8, 20, 10, 20, 12, 20, 130, 9, 20,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 4, 22, 141,
		8, 22, 11, 22, 12, 22, 142, 1, 23, 1, 23, 1, 23, 5, 23, 148, 8, 23, 10,
		23, 12, 23, 151, 9, 23, 1, 24, 4, 24, 154, 8, 24, 11, 24, 12, 24, 155,
		1, 24, 1, 24, 2, 112, 128, 0, 25, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6,
		13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31,
		16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49,
		25, 1, 0, 4, 3, 0, 65, 90, 97, 122, 124, 124, 1, 0, 48, 57, 1, 0, 49, 57,
		3, 0, 9, 10, 13, 13, 32, 32, 167, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0,
		5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0,
		13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0,
		0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0,
		0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0,
		0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1,
		0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 1, 51,
		1, 0, 0, 0, 3, 57, 1, 0, 0, 0, 5, 59, 1, 0, 0, 0, 7, 61, 1, 0, 0, 0, 9,
		63, 1, 0, 0, 0, 11, 65, 1, 0, 0, 0, 13, 67, 1, 0, 0, 0, 15, 69, 1, 0, 0,
		0, 17, 71, 1, 0, 0, 0, 19, 73, 1, 0, 0, 0, 21, 75, 1, 0, 0, 0, 23, 77,
		1, 0, 0, 0, 25, 79, 1, 0, 0, 0, 27, 81, 1, 0, 0, 0, 29, 83, 1, 0, 0, 0,
		31, 85, 1, 0, 0, 0, 33, 88, 1, 0, 0, 0, 35, 90, 1, 0, 0, 0, 37, 104, 1,
		0, 0, 0, 39, 106, 1, 0, 0, 0, 41, 122, 1, 0, 0, 0, 43, 136, 1, 0, 0, 0,
		45, 138, 1, 0, 0, 0, 47, 144, 1, 0, 0, 0, 49, 153, 1, 0, 0, 0, 51, 52,
		5, 92, 0, 0, 52, 53, 5, 102, 0, 0, 53, 54, 5, 114, 0, 0, 54, 55, 5, 97,
		0, 0, 55, 56, 5, 99, 0, 0, 56, 2, 1, 0, 0, 0, 57, 58, 5, 44, 0, 0, 58,
		4, 1, 0, 0, 0, 59, 60, 5, 46, 0, 0, 60, 6, 1, 0, 0, 0, 61, 62, 5, 59, 0,
		0, 62, 8, 1, 0, 0, 0, 63, 64, 5, 40, 0, 0, 64, 10, 1, 0, 0, 0, 65, 66,
		5, 41, 0, 0, 66, 12, 1, 0, 0, 0, 67, 68, 5, 123, 0, 0, 68, 14, 1, 0, 0,
		0, 69, 70, 5, 125, 0, 0, 70, 16, 1, 0, 0, 0, 71, 72, 5, 95, 0, 0, 72, 18,
		1, 0, 0, 0, 73, 74, 5, 94, 0, 0, 74, 20, 1, 0, 0, 0, 75, 76, 5, 92, 0,
		0, 76, 22, 1, 0, 0, 0, 77, 78, 5, 43, 0, 0, 78, 24, 1, 0, 0, 0, 79, 80,
		5, 45, 0, 0, 80, 26, 1, 0, 0, 0, 81, 82, 5, 42, 0, 0, 82, 28, 1, 0, 0,
		0, 83, 84, 5, 47, 0, 0, 84, 30, 1, 0, 0, 0, 85, 86, 5, 42, 0, 0, 86, 87,
		5, 42, 0, 0, 87, 32, 1, 0, 0, 0, 88, 89, 5, 61, 0, 0, 89, 34, 1, 0, 0,
		0, 90, 91, 7, 0, 0, 0, 91, 36, 1, 0, 0, 0, 92, 94, 3, 21, 10, 0, 93, 95,
		7, 0, 0, 0, 94, 93, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 94, 1, 0, 0, 0,
		96, 97, 1, 0, 0, 0, 97, 105, 1, 0, 0, 0, 98, 100, 7, 0, 0, 0, 99, 101,
		7, 0, 0, 0, 100, 99, 1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102, 100, 1, 0,
		0, 0, 102, 103, 1, 0, 0, 0, 103, 105, 1, 0, 0, 0, 104, 92, 1, 0, 0, 0,
		104, 98, 1, 0, 0, 0, 105, 38, 1, 0, 0, 0, 106, 107, 5, 47, 0, 0, 107, 108,
		5, 47, 0, 0, 108, 112, 1, 0, 0, 0, 109, 111, 9, 0, 0, 0, 110, 109, 1, 0,
		0, 0, 111, 114, 1, 0, 0, 0, 112, 113, 1, 0, 0, 0, 112, 110, 1, 0, 0, 0,
		113, 116, 1, 0, 0, 0, 114, 112, 1, 0, 0, 0, 115, 117, 5, 13, 0, 0, 116,
		115, 1, 0, 0, 0, 116, 117, 1, 0, 0, 0, 117, 118, 1, 0, 0, 0, 118, 119,
		5, 10, 0, 0, 119, 120, 1, 0, 0, 0, 120, 121, 6, 19, 0, 0, 121, 40, 1, 0,
		0, 0, 122, 123, 5, 47, 0, 0, 123, 124, 5, 42, 0, 0, 124, 128, 1, 0, 0,
		0, 125, 127, 9, 0, 0, 0, 126, 125, 1, 0, 0, 0, 127, 130, 1, 0, 0, 0, 128,
		129, 1, 0, 0, 0, 128, 126, 1, 0, 0, 0, 129, 131, 1, 0, 0, 0, 130, 128,
		1, 0, 0, 0, 131, 132, 5, 42, 0, 0, 132, 133, 5, 47, 0, 0, 133, 134, 1,
		0, 0, 0, 134, 135, 6, 20, 0, 0, 135, 42, 1, 0, 0, 0, 136, 137, 7, 1, 0,
		0, 137, 44, 1, 0, 0, 0, 138, 140, 7, 2, 0, 0, 139, 141, 7, 1, 0, 0, 140,
		139, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 140, 1, 0, 0, 0, 142, 143,
		1, 0, 0, 0, 143, 46, 1, 0, 0, 0, 144, 145, 3, 45, 22, 0, 145, 149, 5, 46,
		0, 0, 146, 148, 7, 1, 0, 0, 147, 146, 1, 0, 0, 0, 148, 151, 1, 0, 0, 0,
		149, 147, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150, 48, 1, 0, 0, 0, 151, 149,
		1, 0, 0, 0, 152, 154, 7, 3, 0, 0, 153, 152, 1, 0, 0, 0, 154, 155, 1, 0,
		0, 0, 155, 153, 1, 0, 0, 0, 155, 156, 1, 0, 0, 0, 156, 157, 1, 0, 0, 0,
		157, 158, 6, 24, 0, 0, 158, 50, 1, 0, 0, 0, 10, 0, 96, 102, 104, 112, 116,
		128, 142, 149, 155, 1, 6, 0, 0,
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

// FormulaLexerInit initializes any static state used to implement FormulaLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewFormulaLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func FormulaLexerInit() {
	staticData := &FormulaLexerLexerStaticData
	staticData.once.Do(formulalexerLexerInit)
}

// NewFormulaLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewFormulaLexer(input antlr.CharStream) *FormulaLexer {
	FormulaLexerInit()
	l := new(FormulaLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &FormulaLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Formula.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// FormulaLexer tokens.
const (
	FormulaLexerT__0         = 1
	FormulaLexerCOMMA        = 2
	FormulaLexerDOT          = 3
	FormulaLexerSEMICOLON    = 4
	FormulaLexerOPENPAREN    = 5
	FormulaLexerCLOSEPAREN   = 6
	FormulaLexerOPENCURLY    = 7
	FormulaLexerCLOSECURLY   = 8
	FormulaLexerSUBSCRIPT    = 9
	FormulaLexerSUPERSCRIPT  = 10
	FormulaLexerBACKSLASH    = 11
	FormulaLexerPLUS         = 12
	FormulaLexerMINUS        = 13
	FormulaLexerMULT         = 14
	FormulaLexerDIV          = 15
	FormulaLexerPOW          = 16
	FormulaLexerEQUAL        = 17
	FormulaLexerSINGLEID     = 18
	FormulaLexerID           = 19
	FormulaLexerLINE_COMMENT = 20
	FormulaLexerCOMMENT      = 21
	FormulaLexerSINGLEINTLIT = 22
	FormulaLexerINTLIT       = 23
	FormulaLexerFLOATLIT     = 24
	FormulaLexerWS           = 25
)
