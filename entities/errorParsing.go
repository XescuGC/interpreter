package entities

type errorParsing struct {
	s string
}

func (e *errorParsing) Error() string {
	return e.s
}

func newErrorParsing() *errorParsing {
	return &errorParsing{s: "Error Parsing"}
}
