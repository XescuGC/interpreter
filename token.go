package main

type Token struct {
	kind  string
	value interface{}
}

func NewToken(k string, v interface{}) *Token {
	return &Token{kind: k, value: v}
}
