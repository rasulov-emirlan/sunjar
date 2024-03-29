package ast

import (
	"strings"

	"github.com/rasulov-emirlan/sunjar/src/token"
)

type (
	Node interface {
		TokenLiteral() string
		String() string
	}

	Statement interface {
		Node
		statementNode()
	}

	Expression interface {
		Node
		expressionNode()
	}
)

type Program struct {
	Statements []Statement
}

func (p Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p Program) String() string {
	var out string

	for _, s := range p.Statements {
		out += s.String()
	}

	return out
}

type VarStatement struct {
	Token token.Token // the token.VAR token
	Name  *Identifier
	Value Expression
}

func (vs *VarStatement) statementNode() {}

func (vs *VarStatement) TokenLiteral() string {
	return vs.Token.Literal
}

func (vs *VarStatement) String() string {
	var out string

	out += vs.TokenLiteral() + " "
	out += vs.Name.String()
	out += " = "

	if vs.Value != nil {
		out += vs.Value.String()
	}

	out += ";"

	return out
}

type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) String() string {
	return i.Value
}

type ReturnStatement struct {
	Token       token.Token // the 'return' token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (rs *ReturnStatement) String() string {
	var out string

	out += rs.TokenLiteral() + " "

	if rs.ReturnValue != nil {
		out += rs.ReturnValue.String()
	}

	out += ";"

	return out
}

type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}

func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}

func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}

type FloatLiteral struct {
	Token token.Token
	Value float64
}

func (fl *FloatLiteral) expressionNode() {}

func (fl *FloatLiteral) TokenLiteral() string {
	return fl.Token.Literal
}

func (fl *FloatLiteral) String() string {
	return fl.Token.Literal
}

type StringLiteral struct {
	Token token.Token
	Value string
}

func (sl *StringLiteral) expressionNode() {}

func (sl *StringLiteral) TokenLiteral() string {
	return sl.Token.Literal
}

func (sl *StringLiteral) String() string {
	return sl.Token.Literal
}

type PrefixExpression struct {
	Token    token.Token // The prefix token, e.g. !
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode() {}

func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

func (pe *PrefixExpression) String() string {
	var out string

	out += "("
	out += pe.Operator
	out += pe.Right.String()
	out += ")"

	return out
}

type InfixExpression struct {
	Token    token.Token // The operator token, e.g. +
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode() {}

func (ie *InfixExpression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *InfixExpression) String() string {
	var out string

	out += "("
	out += ie.Left.String()
	out += " " + ie.Operator + " "
	out += ie.Right.String()
	out += ")"

	return out
}

type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode() {}

func (b *Boolean) TokenLiteral() string {
	return b.Token.Literal
}

func (b *Boolean) String() string {
	return b.Token.Literal
}

type BlockStatement struct {
	Token      token.Token // the '{' token
	Statements []Statement
}

func (bs *BlockStatement) statementNode() {}

func (bs *BlockStatement) TokenLiteral() string {
	return bs.Token.Literal
}

func (bs *BlockStatement) String() string {
	var out string

	for _, s := range bs.Statements {
		out += s.String()
	}

	return out
}

type IfExpression struct {
	Token       token.Token // the 'if' token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode() {}

func (ie *IfExpression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *IfExpression) String() string {
	var out string

	out += "if"
	out += ie.Condition.String()
	out += " "
	out += ie.Consequence.String()

	if ie.Alternative != nil {
		out += "else "
		out += ie.Alternative.String()
	}

	return out
}

type FunctionLiteral struct {
	Token      token.Token // the 'fn' token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) expressionNode() {}

func (fl *FunctionLiteral) TokenLiteral() string {
	return fl.Token.Literal
}

func (fl *FunctionLiteral) String() string {
	var out string

	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out += fl.TokenLiteral()
	out += "("
	out += strings.Join(params, ", ")
	out += ")"
	out += fl.Body.String()

	return out
}

type CallExpression struct {
	Token     token.Token // the '(' token
	Function  Expression  // Identifier or FunctionLiteral
	Arguments []Expression
}

func (ce *CallExpression) expressionNode() {}

func (ce *CallExpression) TokenLiteral() string {
	return ce.Token.Literal
}

func (ce *CallExpression) String() string {
	var out string

	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	out += ce.Function.String()
	out += "("
	out += strings.Join(args, ", ")
	out += ")"

	return out
}

type ForExpression struct {
	Token     token.Token // the 'for' token
	Condition Expression
	Body      *BlockStatement
}

func (fe *ForExpression) expressionNode() {}

func (fe *ForExpression) TokenLiteral() string {
	return fe.Token.Literal
}

func (fe *ForExpression) String() string {
	var out string

	out += "for"
	out += fe.Condition.String()
	out += " "
	out += fe.Body.String()

	return out
}

type ArrayLiteral struct {
	Token    token.Token // the '[' token
	Elements []Expression
}

func (al *ArrayLiteral) expressionNode() {}

func (al *ArrayLiteral) TokenLiteral() string {
	return al.Token.Literal
}

func (al *ArrayLiteral) String() string {
	var out string

	elements := []string{}
	for _, e := range al.Elements {
		elements = append(elements, e.String())
	}

	out += "[" + strings.Join(elements, ", ") + "]"

	return out
}

type HashLiteral struct {
	Token token.Token // the '{' token
	Pairs map[Expression]Expression
}

func (hl *HashLiteral) expressionNode() {}

func (hl *HashLiteral) TokenLiteral() string {
	return hl.Token.Literal
}

func (hl *HashLiteral) String() string {
	var out string

	pairs := []string{}
	for key, value := range hl.Pairs {
		pairs = append(pairs, key.String()+":"+value.String())
	}

	out += "{" + strings.Join(pairs, ", ") + "}"

	return out
}

type IndexExpression struct {
	Token token.Token // the '[' token
	Left  Expression
	Index Expression
}

func (ie *IndexExpression) expressionNode() {}

func (ie *IndexExpression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *IndexExpression) String() string {
	var out string

	out += "(" + ie.Left.String() + "[" + ie.Index.String() + "])"

	return out
}
