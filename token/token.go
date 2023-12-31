package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT  = "IDENT"  // add, foobar, x, y, ...
	INT    = "INT"    // 1343456
	STRING = "STRING" // "foobar"

	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"

	EQ     = "=="
	NotEq  = "!="
	RETURN = "->"

	AND = "&&"
	OR  = "||"

	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LBRACKET  = "["
	RBRACKET  = "]"
	COLON     = ":"

	FUNCTION = "FK"
	LET      = "Cong"
	SET      = "SET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
)

var Keywords = map[string]TokenType{
	"fk":    FUNCTION,
	"cong":  LET,
	"set":   SET,
	"true":  TRUE,
	"false": FALSE,
	"if":    IF,
	"else":  ELSE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := Keywords[ident]; ok {
		return tok
	}
	return IDENT
}
