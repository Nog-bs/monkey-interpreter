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
