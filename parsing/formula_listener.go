// Code generated from Formula.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Formula
import "github.com/antlr4-go/antlr/v4"

// FormulaListener is a complete listener for a parse tree produced by FormulaParser.
type FormulaListener interface {
	antlr.ParseTreeListener

	// EnterEquation is called when entering the equation production.
	EnterEquation(c *EquationContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterBinaryOperator is called when entering the binaryOperator production.
	EnterBinaryOperator(c *BinaryOperatorContext)

	// EnterFraction is called when entering the fraction production.
	EnterFraction(c *FractionContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterSubscriptTail is called when entering the subscriptTail production.
	EnterSubscriptTail(c *SubscriptTailContext)

	// EnterArgumentTail is called when entering the argumentTail production.
	EnterArgumentTail(c *ArgumentTailContext)

	// EnterArgumentList is called when entering the argumentList production.
	EnterArgumentList(c *ArgumentListContext)

	// EnterArgumentListTail is called when entering the argumentListTail production.
	EnterArgumentListTail(c *ArgumentListTailContext)

	// ExitEquation is called when exiting the equation production.
	ExitEquation(c *EquationContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitBinaryOperator is called when exiting the binaryOperator production.
	ExitBinaryOperator(c *BinaryOperatorContext)

	// ExitFraction is called when exiting the fraction production.
	ExitFraction(c *FractionContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitSubscriptTail is called when exiting the subscriptTail production.
	ExitSubscriptTail(c *SubscriptTailContext)

	// ExitArgumentTail is called when exiting the argumentTail production.
	ExitArgumentTail(c *ArgumentTailContext)

	// ExitArgumentList is called when exiting the argumentList production.
	ExitArgumentList(c *ArgumentListContext)

	// ExitArgumentListTail is called when exiting the argumentListTail production.
	ExitArgumentListTail(c *ArgumentListTailContext)
}
