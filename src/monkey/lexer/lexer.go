/*
 * Lexer module that performs the lexical analysis of the given source code.
 *
 * @author Edro Gonzales
 * @version 2024
 */

package lexer

import (
	"monkey/token"
)

/*
 * Lexer represents a lexical analyzer for the Monkey programming language.
 * It's responsible for breaking down the input source code into a stream of tokens.
 */
type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

/*
 * New creates a new Lexer instance.
 * 
 * @param input  the input string to be analyzed
 * @return a new Lexer ready to tokenize the input
 */
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

/*
 * readChar reads the next character from the input and advances the lexer's positions.
 * If the end of the input is reached, the 'ch' field is set to 0.
 * 
 * @param l is a pointer towards the lexer instance
 */
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

/*
 * Returns the current token being looked at and then moves the pointer of the current Lexer to next.
 *
 * @param l is a Lexer that is evaluating the source code
 * @return tok which is the current token to be analyzed
 */
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhiteSpace()

	// deciding which tokenType this current token is and then providing it the literal passed in
	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Type = token.EOF
		tok.Literal = ""
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			// we can return early because readIdentifier will move position of Lexer internally
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	// moves pointer to next token
	l.readChar()
	// returns current token
	return tok
}

func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

/*
 * Will find within an lexer the whole piece of a word and return that portion
 * 
 * @param l a lexer that is performing lexical analysis on a piece of source code
 * @return input a word
 */
func (l *Lexer) readIdentifier() string {
	startPosition := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[startPosition: l.position]
}

/**
 * Determines if ascii code of given byte is within an alphanumeric value.
 *
 * @param ch is a byte that represents the ascii code value of the current char
 * @return bool true if bye is a letter and false if not
 */
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

/*
 * Generates a new token from the struct made earlier.
 *
 * @param tokenType tokenType of the token
 * @param ch the numerical character representation of the literal token
 * @return token that is of type Token
 */
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}