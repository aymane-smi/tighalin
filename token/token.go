package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	//Special
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//identifier + literals

	IDENT = "IDENT"
	INT   = "INT"

	//Operators

	ASSIGN = "ASSIGN"
	PLUS   = "PLUS"

	//Delimiters | Separators

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//Keywords

	FUNCTION = "FUNCTION"
	LET      = "LET"
	CONST    = "CONST"
)

var Keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}