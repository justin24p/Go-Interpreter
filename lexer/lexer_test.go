package lexer

import (
	"testing"

	"monkey/token"
)

// what is a pointer in go?
// a pointer is a variable that stores the memory address of another variable
// pointers allow you to reference a var or struct from multiple places in code without
// copying data

// test funcitons must start with the word Test
// the t *testing.T is a pointer to the testing.T struct from the testing package

// anytime a funciton uses pointers such as (structProperty *StructName) when passing it down
// i must use &structInstance
func TestNextToken(t *testing.T) {
	// := declare and intilaize and infer variable
	// can only be used within a funciton
	input := `=+(){};`

	// a struct is a similar to a class where you can declare variables no oop properties tho
	tests := []struct {
		expectedType token.TokenType
		expectedLiteral string
	} {
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE,"{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON,";"},
		{token.EOF, ""},

	}
	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype.wrong. expected-%q, got=%",i,&tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("test[%d] - literal wrong, expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}