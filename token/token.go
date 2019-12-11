package token

import "strings"

// Yell announces your message louder
func Yell(message string) string {
	return strings.ToUpper(message)
}



type TokenKind int 

const (
	ILLEGAL TokenKind = iota 
	EOF
	INCREMENT
	DECREMENT
)

type Token struct {
	Kind TokenKind 
	Literal string
}

var tokens = [...]string{
	ILLEGAL:	"ILLEGAL",
	EOF:		"EOF",
	INCREMENT:	"INCREMENT",
	DECREMENT:	"DECREMENT",
}
