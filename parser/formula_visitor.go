// Code generated from Formula.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Formula
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by FormulaParser.
type FormulaVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by FormulaParser#equation.
	VisitEquation(ctx *EquationContext) interface{}

	// Visit a parse tree produced by FormulaParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by FormulaParser#binaryOperator.
	VisitBinaryOperator(ctx *BinaryOperatorContext) interface{}

	// Visit a parse tree produced by FormulaParser#fraction.
	VisitFraction(ctx *FractionContext) interface{}

	// Visit a parse tree produced by FormulaParser#constant.
	VisitConstant(ctx *ConstantContext) interface{}

	// Visit a parse tree produced by FormulaParser#variable.
	VisitVariable(ctx *VariableContext) interface{}

	// Visit a parse tree produced by FormulaParser#subscriptTail.
	VisitSubscriptTail(ctx *SubscriptTailContext) interface{}

	// Visit a parse tree produced by FormulaParser#argumentTail.
	VisitArgumentTail(ctx *ArgumentTailContext) interface{}

	// Visit a parse tree produced by FormulaParser#argumentList.
	VisitArgumentList(ctx *ArgumentListContext) interface{}

	// Visit a parse tree produced by FormulaParser#argumentListTail.
	VisitArgumentListTail(ctx *ArgumentListTailContext) interface{}

	// Visit a parse tree produced by FormulaParser#generalId.
	VisitGeneralId(ctx *GeneralIdContext) interface{}

	// Visit a parse tree produced by FormulaParser#generalIntLit.
	VisitGeneralIntLit(ctx *GeneralIntLitContext) interface{}
}
