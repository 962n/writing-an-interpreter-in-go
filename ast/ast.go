package ast

import (
	"bytes"
	"github.com/962n/writing-an-interpreter-in-go/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

// Statement : 文
type Statement interface {
	Node
	statementNode()
}

// Expression : 式
type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {

}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")
	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {
	// おや、let 文の識別子は値を生成しなかったのではない?
	//なぜ Expressionなのだろうか? これは簡単のためだ。
	//Monkey プログラムの他の部分では識 別子は値を生成する。
	//例えばlet x = valueProducingIdentifier;という具合だ。

}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) String() string {
	return i.Value
}

type ReturnStatement struct {
	Token       token.Token // 'return'トークン
	ReturnValue Expression
}

func (r *ReturnStatement) statementNode() {
}

func (r *ReturnStatement) TokenLiteral() string { return r.Token.Literal }

func (r *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(r.ReturnValue.TokenLiteral() + " ")
	if r.ReturnValue != nil {
		out.WriteString(r.ReturnValue.String())
	}
	out.WriteString(";")
	return r.Token.Literal
}

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {
}

func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
