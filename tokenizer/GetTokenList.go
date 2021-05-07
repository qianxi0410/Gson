package tokenizer

import (
	"gson/errors"
	"log"
	"strconv"
)

func GetTokenList(jsonStr string) *TokenList {
	tokenList := getTokenList()
	reader := convStr2Reader(jsonStr)

	for reader.hasMore() {
		r := reader.read()
		switch r {
		case '{':
			beginObject(tokenList)
		case '}':
			endObject(tokenList)
		case '[':
			beginArray(tokenList, reader)
		case ']':
			endArray(tokenList)
		case 't', 'f':
			err := Bool(tokenList, reader, r)
			if err != nil {
				log.Fatal(err)
			}
		case 'n':
			err := Null(tokenList, reader, r)
			if err != nil {
				log.Fatal(err)
			}
		case '"':
			String(tokenList, reader, r)
		case '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			err := number(tokenList, reader, r)
			if err != nil {
				log.Fatal(err)
			}
		case ':':
			Colon(tokenList)
		case ',':
			Comma(tokenList)
		}
	}
	// help gc
	reader = nil
	return tokenList
}

func endObject(list *TokenList) {
	list.Add(Token{
		TokenType: END_OBJECT,
		Value:     '}',
	})
}

func beginObject(list *TokenList) {
	list.Add(Token{
		TokenType: BEGIN_OBJECT,
		Value:     '{',
	})
}

func endArray(list *TokenList) {
	list.Add(Token{
		TokenType: END_ARRAY,
		Value:     ']',
	})
}

func beginArray(list *TokenList, r *reader) {
	list.Add(Token{
		TokenType: BEGIN_ARRAY,
		Value:     '[',
	})
}

func String(list *TokenList, r *reader, ch rune) {
	sVal := ""

	for r.peek() != '"' {
		sVal += string(r.read())
	}
	r.read()
	list.Add(Token{
		TokenType: STRING,
		Value:     sVal,
	})
}

func Null(list *TokenList, r *reader, ch rune) error {
	nVal := string(ch)
	for i := 0; isNotEnd(r.peek()) && r.hasMore(); i++ {
		nVal += string(r.read())
	}
	if nVal != "null" {
		return errors.ParseNullError{}
	}
	list.Add(Token{
		TokenType: NULL,
		Value:     nil,
	})
	return nil
}

func Bool(list *TokenList, r *reader, ch rune) error {
	bVal := string(ch)
	if bVal == "t" {
		for i := 0; isNotEnd(r.peek()) && r.hasMore(); i++ {
			bVal += string(r.read())
		}
		
		if bVal != "true" {
			return errors.ParseBoolError{}
		}
		
		list.Add(Token{
			TokenType: BOOLEAN,
			Value:     true,
		})
		return nil
	} else {
		for i := 0; isNotEnd(r.peek()) && r.hasMore(); i++ {
			bVal += string(r.read())
		}
		
		if bVal != "false" {
			return errors.ParseBoolError{}
		}
		
		list.Add(Token{
			TokenType: BOOLEAN,
			Value:     false,
		})
		return nil
	}
}

func number(list *TokenList, r *reader, ch rune) error {
	nVal := string(ch)
	for isNotEnd(r.peek()) {
		nVal += string(r.read())
	}
	float, err := strconv.ParseFloat(nVal, 32)
	if err != nil {
		return errors.ParseFloatError{}
	}
	list.Add(Token{
		TokenType: NUMBER,
		Value:     float,
	})
	return nil
}

func Colon(list *TokenList) {
	list.Add(Token{
		TokenType: SEP_COLON,
		Value:     ':',
	})
}

func Comma(list *TokenList) {
	list.Add(Token{
		TokenType: SEP_COMMA,
		Value:     ',',
	})
}