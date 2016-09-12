package main

import "strconv"

const (
	INTEGER = "INTEGER"
	MINUS   = "MINUS"
	PLUS    = "PLUS"
	EOF     = "EOF"
)

type Interpreter struct {
	text          string
	pos           int
	current_token *Token
	current_char  string
}

func NewInterpreter(t string) *Interpreter {
	return &Interpreter{text: t, pos: 0, current_token: nil, current_char: string(t[0])}
}

func (i *Interpreter) ErrorParsing() string {
	return "Error Parsing"
}

func (i *Interpreter) Advance() {
	i.pos++
	// Remove the 2 when empty spaces
	if i.pos > len(i.text)-2 {
		i.current_char = ""
	} else {
		i.current_char = string(i.text[i.pos])
	}
}

func (i *Interpreter) SkipWhitespace() {
	for i.current_char != "" && i.current_char == " " {
		i.Advance()
	}
}

func (i *Interpreter) Integer() int {
	result := ""
	for i.current_char != "" {
		if _, err := strconv.Atoi(i.current_char); err == nil {
			result += i.current_char
			i.Advance()
		}
		break
	}

	res, _ := strconv.Atoi(result)
	return res
}

func (i *Interpreter) NextToken() *Token {

	for i.current_char != "" {
		if i.current_char == " " {
			i.SkipWhitespace()
			//continue
		}
		if _, err := strconv.Atoi(i.current_char); err == nil {
			return NewToken(INTEGER, i.Integer())
		}

		if i.current_char == "+" {
			i.Advance()
			return NewToken(PLUS, "+")
		}

		if i.current_char == "-" {
			i.Advance()
			return NewToken(MINUS, "-")
		}
		panic(i.ErrorParsing())
	}
	return NewToken(EOF, nil)

}

func (i *Interpreter) eat(token_kind string) {
	if i.current_token.kind == token_kind {
		i.current_token = i.NextToken()
	} else {
		panic(i.ErrorParsing())
	}
}

func (i *Interpreter) expr() interface{} {
	var result interface{}
	i.current_token = i.NextToken()

	left := i.current_token
	i.eat(INTEGER)
	l, _ := left.value.(int)

	op := i.current_token
	if op.kind == PLUS {
		i.eat(PLUS)
	} else if op.kind == MINUS {
		i.eat(MINUS)
	}

	right := i.current_token
	i.eat(INTEGER)
	r, _ := right.value.(int)

	if op.kind == PLUS {
		result = l + r
	} else if op.kind == MINUS {
		result = l - r
	}

	return result
}
