package errors
// some error when your parse your token

type GetBoolTokenError struct {
	error
}

func (p GetBoolTokenError) Error() string {
	return "some thing error when you get your bool token!!!"
}

type GetNumberTokenError struct {
	error
}

func (p GetNumberTokenError) Error() string {
	return "some thing error when you get your float token!!!"
}

type GetNullTokenError struct {
	error
}

func (p GetNullTokenError) Error() string {
	return "some thing error when you get your null token!!!"
}

