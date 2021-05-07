package parser

import (
	"gson/errors"
	"gson/tokenizer"
	"log"
)

type Parser struct {
	tokenizer.TokenList
}

func GetParser(list *tokenizer.TokenList) *Parser {
	list.ReSet()
	return &Parser{*list}
}

func (p *Parser) ParseJsonObject() map[string]interface{} {
	var jsonObj map[string]interface{}
	expectToken := tokenizer.BEGIN_OBJECT

	var key string

	for p.HasMore() {
		token := p.Peek()
		tokenType, tokenValue := token.TokenType, token.Value

		switch tokenType {
		case tokenizer.BEGIN_OBJECT:
			if err := checkTokenType(tokenType, expectToken); err != nil {
				log.Fatal(err)
			}
			if key != "" {
				jsonObj[key] = p.ParseJsonObject()
				expectToken = tokenizer.END_OBJECT
			} else {
				p.Read()
				jsonObj = make(map[string]interface{})
				expectToken = tokenizer.END_OBJECT | tokenizer.STRING
			}
		case tokenizer.END_OBJECT:
			if err := checkTokenType(tokenType, expectToken); err != nil {
				log.Fatal(err)
			}
			p.Read()
			key = ""
			expectToken = tokenizer.END_OBJECT | tokenizer.SEP_COMMA
			return jsonObj
		case tokenizer.BEGIN_ARRAY:
			if err := checkTokenType(tokenType, expectToken); err != nil {
				log.Fatal(err)
			}
			p.Read()
			jsonObj[key] = ParseJsonArray()
			expectToken = tokenizer.END_ARRAY
		case tokenizer.END_ARRAY:
			if err := checkTokenType(tokenType, expectToken); err != nil {
				log.Fatal(err)
			}
			p.Read()
			key = ""
		case tokenizer.STRING:
			if err := checkTokenType(tokenType, expectToken); err != nil {
				log.Fatal(err)
			}
			p.Read()
			if key != "" {
				jsonObj[key] = tokenValue
				key = ""
				expectToken = tokenizer.END_OBJECT | tokenizer.SEP_COMMA
			} else {
				key = tokenValue.(string)
				expectToken = tokenizer.SEP_COLON
			}
		case tokenizer.NUMBER:
			if err := checkTokenType(tokenType, expectToken); err != nil {
				log.Fatal(err)
			}
			p.Read()
			if key == "" {
				panic("invalid json key!")
			} else {
				jsonObj[key] = tokenValue.(float64)
				key = ""
				expectToken = tokenizer.END_OBJECT | tokenizer.SEP_COMMA
			}
		case tokenizer.NULL:
			if err := checkTokenType(tokenType, expectToken); err != nil {
				log.Fatal(err)
			}
			p.Read()
			if key == "" {
				panic("invalid json key!")
			} else {
				jsonObj[key] = nil
				key = ""
				expectToken = tokenizer.END_OBJECT | tokenizer.SEP_COMMA
			}
		case tokenizer.BOOLEAN:
			if err := checkTokenType(tokenType, expectToken); err != nil {
				log.Fatal(err)
			}
			p.Read()
			if key == "" {
				panic("invalid json key!")
			} else {
				jsonObj[key] = tokenValue
				key = ""
				expectToken = tokenizer.END_OBJECT | tokenizer.SEP_COMMA
			}
		case tokenizer.SEP_COMMA:
			if err := checkTokenType(tokenType, expectToken); err != nil {
				log.Fatal(err)
			}
			p.Read()
			expectToken = tokenizer.END_OBJECT | tokenizer.STRING
		case tokenizer.SEP_COLON:
			if err := checkTokenType(tokenType, expectToken); err != nil {
				log.Fatal(err)
			}
			p.Read()
			expectToken = tokenizer.BOOLEAN | tokenizer.STRING | tokenizer.NUMBER | tokenizer.NULL | tokenizer.BEGIN_OBJECT | tokenizer.BEGIN_ARRAY
		}
	}
	return nil
}

func checkTokenType(real, expect tokenizer.TokenType) error {
	if real & expect == 0 {
		return errors.TokenTypeParseError{}
	}
	return nil
}

func ParseJsonArray() []interface{} {
	return nil
}