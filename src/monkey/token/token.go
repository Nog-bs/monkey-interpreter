/*
 * Token module that handles all declarations related to token data
 *
 * @author Edro Gonzales
 * @version 2024
 */

package token

type TokenType string

/*
 * Token to be evaluated during the tokenizing process
 * Contains a Type and Literal of which represent the type and the literal representation
 */
type Token struct {
	Type    TokenType
	Literal string
}

/*
 * Global Token Declarations
 */
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	ASSIGN = "="
	PLUS   = "+"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"
)

/*
 * Keywords map that will allow fast look up.
 */
var keywords = map[string]TokenType {
	"fn": FUNCTION,
	"let": LET,
}

/*
 * Looks up identity of keyword based on establish map in the module.
 * 
 * @param ident is a string that is the given keyword or identifier to be evaluated
 * @return returns the given TokenType
 */
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}