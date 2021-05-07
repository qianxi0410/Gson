package tokenizer

import "fmt"

type Token struct {
	TokenType TokenType
	Value     interface{}
	next      *Token
}

type TokenList struct {
	// an list of token
	tokens *Token

	cur *Token
}

func (t *TokenList) Add(token Token)  {
	t.cur.next = &token
	t.cur = t.cur.next
}

func (t *TokenList) Read() Token {
	token := Token{
		TokenType: t.cur.TokenType,
		Value:     t.cur.Value,
	}
	t.cur = t.cur.next
	return token
}

func (t *TokenList) Peek() Token {
	token := Token{
		TokenType: t.cur.TokenType,
		Value:     t.cur.Value,
	}
	//t.cur = t.cur.next
	return token
}

func (t *TokenList) HasMore() bool {
	return t.cur != nil
}

func (t *TokenList) ReSet()  {
	t.cur = t.tokens.next
}

func (t TokenList) ToString()  {
	t.ReSet()
	for t.HasMore() {
		fmt.Printf("%v\n", t.cur.Value)
		t.cur = t.cur.next
	}
}