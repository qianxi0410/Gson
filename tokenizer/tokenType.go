package tokenizer

type TokenType uint16

const (
	BEGIN_OBJECT = TokenType(1 << iota)
	END_OBJECT
	BEGIN_ARRAY
	END_ARRAY
	NULL
	NUMBER
	STRING
	BOOLEAN
	SEP_COLON
	SEP_COMMA
)