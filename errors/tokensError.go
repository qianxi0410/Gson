package errors
// some error when your parse your token

type ParseBoolError struct {
	error
}

func (p ParseBoolError) Error() string {
	return "some thing error when you parse your bool type!!!"
}

type ParseNumberError struct {
	error
}

func (p ParseNumberError) Error() string {
	return "some thing error when you parse your float type!!!"
}

type ParseNullError struct {
	error
}

func (p ParseNullError) Error() string {
	return "some thing error when you parse your null type!!!"
}

