package tokenizer

import (
	"bytes"
	"gson/errors"
	"log"
	"strconv"
)

func GetTokenList(jsonStr string) (*TokenList, error) {
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
			beginArray(tokenList)
		case ']':
			endArray(tokenList)
		case 't', 'f':
			err := getBool(tokenList, reader, r)
			if err != nil {
				log.Fatal(err)
			}
		case 'n':
			err := getNull(tokenList, reader, r)
			if err != nil {
				return nil, err
			}
		case '"':
			getString(tokenList, reader)
		case '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			err := getNumber(tokenList, reader, r)
			if err != nil {
				return nil, err
			}
		case ':':
			getColon(tokenList)
		case ',':
			getComma(tokenList)
		}
	}
	// help gc
	reader = nil
	return tokenList, nil
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

func beginArray(list *TokenList) {
	list.Add(Token{
		TokenType: BEGIN_ARRAY,
		Value:     '[',
	})
}

func getString(list *TokenList, r *reader) {
	var buf bytes.Buffer
	for r.peek() != '"' {
		buf.WriteRune(r.read())
	}
	r.read()
	list.Add(Token{
		TokenType: STRING,
		Value:     buf.String(),
	})
}

func getNull(list *TokenList, r *reader, ch rune) error {
	nVal := string(ch)
	for i := 0; isNotEnd(r.peek()) && r.hasMore(); i++ {
		nVal += string(r.read())
	}
	if nVal != "null" {
		return errors.GetNullTokenError{}
	}
	list.Add(Token{
		TokenType: NULL,
		Value:     nil,
	})
	return nil
}

func getBool(list *TokenList, r *reader, ch rune) error {
	bVal := string(ch)
	if bVal == "t" {
		for i := 0; isNotEnd(r.peek()) && r.hasMore(); i++ {
			bVal += string(r.read())
		}
		
		if bVal != "true" {
			return errors.GetBoolTokenError{}
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
			return errors.GetBoolTokenError{}
		}
		
		list.Add(Token{
			TokenType: BOOLEAN,
			Value:     false,
		})
		return nil
	}
}

func getNumber(list *TokenList, r *reader, ch rune) error {
	nVal := string(ch)
	for isNotEnd(r.peek()) {
		nVal += string(r.read())
	}
	float, err := strconv.ParseFloat(nVal, 64)
	if err != nil {
		return errors.GetNumberTokenError{}
	}
	list.Add(Token{
		TokenType: NUMBER,
		Value:     float,
	})
	return nil
}

func getColon(list *TokenList) {
	list.Add(Token{
		TokenType: SEP_COLON,
		Value:     ':',
	})
}

func getComma(list *TokenList) {
	list.Add(Token{
		TokenType: SEP_COMMA,
		Value:     ',',
	})
}