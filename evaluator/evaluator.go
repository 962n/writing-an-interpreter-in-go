package evaluator

import (
	"github.com/962n/writing-an-interpreter-in-go/ast"
	"github.com/962n/writing-an-interpreter-in-go/object"
)

func Eval(anode ast.Node) object.Object {
	switch node := anode.(type) {
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	}
	return nil
}

func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object
	for _ , stmt := range stmts {
		result = Eval(stmt)
	}
	return result
}