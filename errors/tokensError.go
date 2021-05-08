package errors
// some error when your parse your token

type GetBoolTokenError struct {
	error
}

func (p GetBoolTokenError) Error() string {
	return "some thing error when you parse your bool type!!!"
}

type GetNumberTokenError struct {
	error
}

func (p GetNumberTokenError) Error() string {
	return "some thing error when you parse your float type!!!"
}

type GetNullTokenError struct {
	error
}

func (p GetNullTokenError) Error() string {
	return "some thing error when you parse your null type!!!"
}

