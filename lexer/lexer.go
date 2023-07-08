package lexer

import (
	"github.com/aymane-smi/tighalin/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

// create new Lexer object
func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}
	l.ReadChar()
	return l
}

// read char from Lexer as a method
func (l *Lexer) ReadChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

//get the token from the current char of the Lexer and increment reading position

func (l *Lexer) NextToken() (token.Token){
	var tok token.Token
	
	switch l.ch {
		case '=':
			tok = newToken(token.ASSIGN, l.ch)
		case '+':
			tok = newToken(token.PLUS, l.ch)
		case  ',':
			tok = newToken(token.COMMA, l.ch)
		case ';':
			tok = newToken(token.SEMICOLON, l.ch)
		case '(':
			tok = newToken(token.LPAREN, l.ch)
		case ')':
			tok = newToken(token.RPAREN, l.ch)
		case '{':
			tok = newToken(token.LBRACE, l.ch)
		case '}':
			tok = newToken(token.RBRACE, l.ch)
		case 0:
			tok.Literal = ""
			tok.Type = token.EOF
		default:
			if isValidCharVar(l.ch){
				tok.Literal = l.readIdentifier()
				return tok
			}else{
				tok = newToken(token.ILLEGAL, l.ch)
			}
	}

	l.ReadChar()
	return tok
}

//read identifier

func (l* Lexer) readIdentifier() string{
	position := l.position
	for isValidCharVar(l.ch){
		l.ReadChar()
	}
	return l.input[position:l.position]
}

//check if a char is valid to be in variable name

func isValidCharVar(ch byte) bool{
	return  isLetter(ch) && isNumber(ch) && isSpecial(ch)
}

//helpers for isValidCharVar function

func isLetter(ch byte) bool{
	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z'
}

func isNumber(ch byte) bool{
	var tmp int = int(ch - '0')
	return tmp >= 0 && tmp <= 9
}

func isSpecial(ch byte) bool{
	var chars string = "_@"
	result := false
	for i := 0; i < len(chars); i++{
		if chars[i] == ch{
			result = true
		}
	}
	return result
}



//get the object of the given literal

func newToken(tokenType token.TokenType, ch byte) token.Token{
	return token.Token{
		Type: tokenType,
		Literal: string(ch),
	}
}