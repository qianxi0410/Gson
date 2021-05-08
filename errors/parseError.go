package errors

type ParseError struct {
	error
}

func (p ParseError) Error() string {
	return "there must be something wrong with your json :("
}

type ObjectTypeParseError struct {
	error
}

func (p ObjectTypeParseError) Error() string {
	return "i don't need an object!!!"
}

type ArrayTypeParseError struct {
	error
}

func (p ArrayTypeParseError) Error() string {
	return "i don't need an array!!!"
}

type NumberTypeParseError struct {
	error
}

func (p NumberTypeParseError) Error() string {
	return "i don't need an number!!!"
}

type StringTypeParseError struct {
	error
}

func (p StringTypeParseError) Error() string {
	return "i don't need an string!!!"
}

type NullTypeParseError struct {
	error
}

func (p NullTypeParseError) Error() string {
	return "i don't need an null!!!"
}

type BoolTypeParseError struct {
	error
}

func (p BoolTypeParseError) Error() string {
	return "i don't need an bool!!!"
}

type CommaTypeParseError struct {
	error
}

func (p CommaTypeParseError) Error() string {
	return "i don't need a comma!!!"
}

type ColonTypeParseError struct {
	error
}

func (p ColonTypeParseError) Error() string {
	return "i don't need a colon!!!"
}