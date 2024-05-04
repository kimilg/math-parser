// Code generated from Formula.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Formula
import "github.com/antlr4-go/antlr/v4"

// BaseFormulaListener is a complete listener for a parse tree produced by FormulaParser.
type BaseFormulaListener struct{}

var _ FormulaListener = &BaseFormulaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFormulaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFormulaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFormulaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFormulaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterEquation is called when production equation is entered.
func (s *BaseFormulaListener) EnterEquation(ctx *EquationContext) {}

// ExitEquation is called when production equation is exited.
func (s *BaseFormulaListener) ExitEquation(ctx *EquationContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseFormulaListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseFormulaListener) ExitExpr(ctx *ExprContext) {}

// EnterBinaryOperator is called when production binaryOperator is entered.
func (s *BaseFormulaListener) EnterBinaryOperator(ctx *BinaryOperatorContext) {}

// ExitBinaryOperator is called when production binaryOperator is exited.
func (s *BaseFormulaListener) ExitBinaryOperator(ctx *BinaryOperatorContext) {}

// EnterFraction is called when production fraction is entered.
func (s *BaseFormulaListener) EnterFraction(ctx *FractionContext) {}

// ExitFraction is called when production fraction is exited.
func (s *BaseFormulaListener) ExitFraction(ctx *FractionContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseFormulaListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaseFormulaListener) ExitConstant(ctx *ConstantContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaseFormulaListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseFormulaListener) ExitVariable(ctx *VariableContext) {}

// EnterSubscriptTail is called when production subscriptTail is entered.
func (s *BaseFormulaListener) EnterSubscriptTail(ctx *SubscriptTailContext) {}

// ExitSubscriptTail is called when production subscriptTail is exited.
func (s *BaseFormulaListener) ExitSubscriptTail(ctx *SubscriptTailContext) {}

// EnterArgumentTail is called when production argumentTail is entered.
func (s *BaseFormulaListener) EnterArgumentTail(ctx *ArgumentTailContext) {}

// ExitArgumentTail is called when production argumentTail is exited.
func (s *BaseFormulaListener) ExitArgumentTail(ctx *ArgumentTailContext) {}

// EnterArgumentList is called when production argumentList is entered.
func (s *BaseFormulaListener) EnterArgumentList(ctx *ArgumentListContext) {}

// ExitArgumentList is called when production argumentList is exited.
func (s *BaseFormulaListener) ExitArgumentList(ctx *ArgumentListContext) {}

// EnterArgumentListTail is called when production argumentListTail is entered.
func (s *BaseFormulaListener) EnterArgumentListTail(ctx *ArgumentListTailContext) {}

// ExitArgumentListTail is called when production argumentListTail is exited.
func (s *BaseFormulaListener) ExitArgumentListTail(ctx *ArgumentListTailContext) {}

// EnterGeneralId is called when production generalId is entered.
func (s *BaseFormulaListener) EnterGeneralId(ctx *GeneralIdContext) {}

// ExitGeneralId is called when production generalId is exited.
func (s *BaseFormulaListener) ExitGeneralId(ctx *GeneralIdContext) {}

// EnterGeneralIntLit is called when production generalIntLit is entered.
func (s *BaseFormulaListener) EnterGeneralIntLit(ctx *GeneralIntLitContext) {}

// ExitGeneralIntLit is called when production generalIntLit is exited.
func (s *BaseFormulaListener) ExitGeneralIntLit(ctx *GeneralIntLitContext) {}