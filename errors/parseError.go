package errors

type ParseError struct {
	error
}

func (p ParseError) Error() string {
	return "your json is parser failed!"
}

type TokenTypeParseError struct {
	error
}

func (p TokenTypeParseError) Error() string {
	return "some thing error when you parse your token type!!!"
}