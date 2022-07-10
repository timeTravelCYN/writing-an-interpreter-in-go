package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	// Special tokens
	ILLEGAL = "ILLEGAL" // Illegal token
	EOF = "EOF" // End of file

	// Identifiers + literals 标识符+字面量
	IDENT = "IDENT" // Identifier
	INT = "INT" // Integer literal

	// Operators 操作符
	ASSIGN = "=" // Assignment operator
	PLUS = "+" // Plus operator

	// Delimiters 分隔符
	COMMA = "," // Comma operator
	SEMICOLON = ";" // Semicolon operator

	LPAREN = "(" // Left parenthesis
	RPAREN = ")" // Right parenthesis
	LBRACE = "{" // Left brace
	RBRACE = "}" // Right brace

	// Keywords 关键字
	FUNCTION = "FUNCTION" // Function keyword
	LET = "LET" // Let keyword
)
