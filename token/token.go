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

	ASSIGN  = "+"
	PLUS    = "+"
	MINUS   = "-"
	BANG    = "!"
	ASTRISK = "*"
	SLASH   = "/"
	LT      = "<"
	GT      = ">"

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

// keywords map
var Keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

//look if the given string is a valid keyword

func LookupIdent(ident string) TokenType {
	if tok, ok := Keywords[ident]; ok {
		return tok
	}
	return IDENT
}