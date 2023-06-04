package lexer

import (
	"testing"

	"github.com/aymane-smi/tighalin/token"
)


func testNextToken(t *testing.T){
	input := "=+(){},;"

	tests := []struct{
		expectedType token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.RBRACE, "{"},
		{token.LBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
	}

	l := New(input)

	for i, tt := range tests{
		tok := l.NextToken()

		if tok.Type != tt.expectedType{
			t.Fatalf("test[%d] - token Type is not supported. expected %q got %q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral{
			t.Fatalf("test[%d] - token Literal is not supported. exprected %q got %q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}