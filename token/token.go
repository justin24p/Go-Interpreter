// what are packages in Go?
// packages are a way to group related files similar to modules in javascript
// way to export and import code
package token

// what is a struct in go?
// groups vars under a single name similar to classes but have no inheritence or polymorphism

// a pacakage part of the standard library
type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	// indentifies + literals 
	IDENT = "IDENT"
	INT = "INT"

	//Operators
	ASSIGN = "="
	PLUS = "+"

	// Delimiters

	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = "("
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCITON = "FUNCTION"
	LET = "LET"	
)

type Token struct {
	Type TokenType
	Literal string
}

