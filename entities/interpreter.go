package entities

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

func (i *Interpreter) advance() {
	i.pos++
	// Remove the 2 when empty spaces
	if i.pos > len(i.text)-2 {
		i.current_char = ""
	} else {
		i.current_char = string(i.text[i.pos])
	}
}

func (i *Interpreter) skipWhitespace() {
	for i.current_char != "" && i.current_char == " " {
		i.advance()
	}
}

func (i *Interpreter) integer() int {
	result := ""
	for i.current_char != "" {
		if _, err := strconv.Atoi(i.current_char); err == nil {
			result += i.current_char
			i.advance()
		}
		break
	}

	res, _ := strconv.Atoi(result)
	return res
}

func (i *Interpreter) nextToken() (*Token, error) {

	for i.current_char != "" {
		if i.current_char == " " {
			i.skipWhitespace()
		}
		if _, err := strconv.Atoi(i.current_char); err == nil {
			return NewToken(INTEGER, i.integer()), nil
		}

		if i.current_char == "+" {
			i.advance()
			return NewToken(PLUS, "+"), nil
		}

		if i.current_char == "-" {
			i.advance()
			return NewToken(MINUS, "-"), nil
		}
		return nil, newErrorParsing()
	}
	return NewToken(EOF, nil), nil

}

func (i *Interpreter) eat(token_kind string) error {
	var err error
	if i.current_token.kind == token_kind {
		i.current_token, err = i.nextToken()
		if err != nil {
			return err
		}
		return nil
	} else {
		return newErrorParsing()
	}
}

func (i *Interpreter) Expr() interface{} {
	var result int
	var err error

	i.current_token, err = i.nextToken()
	if err != nil {
		panic(err)
	}

	left := i.current_token
	err = i.eat(INTEGER)
	if err != nil {
		panic(err)
	}
	l, _ := left.value.(int)

	op := i.current_token
	if op.kind == PLUS {
		err = i.eat(PLUS)
		if err != nil {
			panic(err)
		}
	} else if op.kind == MINUS {
		err = i.eat(MINUS)
		if err != nil {
			panic(err)
		}
	}

	right := i.current_token
	err = i.eat(INTEGER)
	if err != nil {
		panic(err)
	}
	r, _ := right.value.(int)

	if op.kind == PLUS {
		result = l + r
	} else if op.kind == MINUS {
		result = l - r
	}

	for {
		ope := i.current_token
		if ope.kind == EOF {
			break
		}
		if ope.kind == PLUS {
			err = i.eat(PLUS)
			if err != nil {
				break
			}
		} else if ope.kind == MINUS {
			err = i.eat(MINUS)
			if err != nil {
				break
			}
		}

		other := i.current_token
		err = i.eat(INTEGER)
		if err != nil {
			panic(err)
		}
		o, _ := other.value.(int)

		if ope.kind == PLUS {
			result += o
		} else if ope.kind == MINUS {
			result -= o
		}
	}

	return result
}
