package ast

import "github.com/962n/writing-an-interpreter-in-go/token"

type Node interface {
	TokenLiteral() string
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
