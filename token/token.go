package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string // can change to int or byte later when attempting to optimize
}

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"
	
	IDENT = "IDENT"
	INT = "INT"
	
	ASSIGN = "="
	PLUS = "+"
	
	COMMA = ","
	SEMICOLON = ";"
	
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	
	FUNCTION = "FUNCTION"
	LET = "LET"
)
// lexer can later have a buffer or save feature, maybe a cache?
// can also include line number and filenames within the process of tokenizing