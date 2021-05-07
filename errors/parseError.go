package errors

type ParseBoolError struct {
	error
}

func (p ParseBoolError) Error() string {
	return "some thing error when you parse your bool type!!!"
}

type ParseFloatError struct {
	error
}

func (p ParseFloatError) Error() string {
	return "some thing error when you parse your float type!!!"
}

type ParseNullError struct {
	error
}

func (p ParseNullError) Error() string {
	return "some thing error when you parse your null type!!!"
}

type TokenTypeParseError struct {
	error
}

func (p TokenTypeParseError) Error() string {
	return "some thing error when you parse your token type!!!"
}