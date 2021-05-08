package parser

import (
	"gson/errors"
	"gson/tokenizer"
	"log"
)

type Parser struct {
	tokenizer.TokenList
}

func GetParser(jsonStr string) (*Parser, error) {
	list, err := tokenizer.GetTokenList(jsonStr)
	if err != nil {
		return nil, err
	}
	list.ReSet()
	return &Parser{*list}, nil
}

func (p *Parser) ParseJsonObject() (map[string]interface{}, error) {
	var jsonObj map[string]interface{}
	expectToken := tokenizer.BEGIN_OBJECT

	var key string

	for p.HasMore() {
		token := p.Peek()
		tokenType, tokenValue := token.TokenType, token.Value

		switch tokenType {
		case tokenizer.BEGIN_OBJECT:
			if err := checkTokenType(tokenType, expectToken); err != nil {
				return nil, err
			}
			if key != "" {
				object, err := p.ParseJsonObject()
				if err != nil {
					return nil, err
				}
				jsonObj[key] = object
				key = ""
				expectToken = tokenizer.END_OBJECT | tokenizer.SEP_COMMA
			} else {
				p.Read()
				jsonObj = make(map[string]interface{})
				expectToken = tokenizer.END_OBJECT | tokenizer.STRING
			}
		case tokenizer.END_OBJECT:
			if err := checkTokenType(tokenType, expectToken); err != nil {
				return nil, err
			}
			p.Read()
			return jsonObj, nil
		case tokenizer.BEGIN_ARRAY:
			if err := checkTokenType(tokenType, expectToken); err != nil {
				log.Fatal(err)
			}
			array, err := p.ParseJsonArray()
			if err != nil {
				 return nil, err
			}
			jsonObj[key] = array
			key = ""
			expectToken = tokenizer.END_OBJECT | tokenizer.SEP_COMMA
		case tokenizer.END_ARRAY: // never
			panic("error when your parse json object !")
		case tokenizer.STRING:
			if err := checkTokenType(tokenType, expectToken); err != nil {
				return nil, err
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
				return nil, err
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
				return nil, err
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
				return nil, err
			}
			p.Read()
			expectToken = tokenizer.STRING
		case tokenizer.SEP_COLON:
			if err := checkTokenType(tokenType, expectToken); err != nil {
				return nil, err
			}
			p.Read()
			expectToken = tokenizer.BOOLEAN | tokenizer.STRING |
							tokenizer.NUMBER | tokenizer.NULL |
							tokenizer.BEGIN_OBJECT | tokenizer.BEGIN_ARRAY
		}
	}
	return nil, errors.ParseError{}
}

func checkTokenType(real, expect tokenizer.TokenType) error {
	if real & expect == 0 {
		switch real {
		case tokenizer.BEGIN_OBJECT, tokenizer.END_OBJECT:
			return errors.ObjectTypeParseError{}
		case tokenizer.BEGIN_ARRAY, tokenizer.END_ARRAY:
			return errors.ArrayTypeParseError{}
		case tokenizer.STRING:
			return errors.StringTypeParseError{}
		case tokenizer.NUMBER:
			return errors.NumberTypeParseError{}
		case tokenizer.NULL:
			return errors.NullTypeParseError{}
		case tokenizer.SEP_COLON:
			return errors.ColonTypeParseError{}
		case tokenizer.SEP_COMMA:
			return errors.CommaTypeParseError{}
		case tokenizer.BOOLEAN:
			return errors.BoolTypeParseError{}
		}
	}
	return nil
}

func (p *Parser) ParseJsonArray() ([]interface{}, error) {
	var jsonArray []interface{}
	expectToken := tokenizer.BEGIN_ARRAY
	var flag bool = false


	for p.HasMore() {
		token := p.Peek()
		tokenType, tokenValue := token.TokenType, token.Value

		switch tokenType {
		case tokenizer.BEGIN_ARRAY:
			if err := checkTokenType(tokenType, expectToken); err != nil {
				return nil, err
			}
			if !flag {
				p.Read()
				jsonArray = make([]interface{}, 0, 1 << 6)
				flag = true
				expectToken = tokenizer.END_ARRAY | tokenizer.STRING |
							tokenizer.BEGIN_ARRAY | tokenizer.NUMBER |
							tokenizer.NULL | tokenizer.BOOLEAN |
							tokenizer.BEGIN_OBJECT
			} else {
				array, err := p.ParseJsonArray()
				if err != nil {
					return nil, err
				}
				jsonArray = append(jsonArray, array)
				expectToken = tokenizer.SEP_COMMA | tokenizer.END_ARRAY
			}
		case tokenizer.END_ARRAY:
			if err := checkTokenType(tokenType, expectToken); err != nil {
				return nil, err
			}
			p.Read()
			return jsonArray, nil
		case tokenizer.BEGIN_OBJECT:
			if err := checkTokenType(tokenType, expectToken); err != nil {
				return nil, err
			}
			object, err := p.ParseJsonObject()
			if err != nil {
				return nil, err
			}
			jsonArray = append(jsonArray, object)
			expectToken = tokenizer.SEP_COMMA | tokenizer.END_ARRAY
		case tokenizer.END_OBJECT:
			panic("error when your parse json array !")
		case tokenizer.STRING:
			if err := checkTokenType(tokenType, expectToken); err != nil {
				return nil, err
			}
			p.Read()
			jsonArray = append(jsonArray, tokenValue.(string))
			expectToken = tokenizer.SEP_COMMA | tokenizer.END_ARRAY
		case tokenizer.NUMBER:
			if err := checkTokenType(tokenType, expectToken); err != nil {
				return nil, err
			}
			p.Read()
			jsonArray = append(jsonArray, tokenValue.(float64))
			expectToken = tokenizer.SEP_COMMA | tokenizer.END_ARRAY
		case tokenizer.NULL:
			if err := checkTokenType(tokenType, expectToken); err != nil {
				return nil, err
			}
			p.Read()
			jsonArray = append(jsonArray, nil)
			expectToken = tokenizer.SEP_COMMA | tokenizer.END_ARRAY
		case tokenizer.BOOLEAN:
			if err := checkTokenType(tokenType, expectToken); err != nil {
				return nil, err
			}
			p.Read()
			jsonArray = append(jsonArray, tokenValue.(bool))
			expectToken = tokenizer.SEP_COMMA | tokenizer.END_ARRAY
		case tokenizer.SEP_COMMA:
			if err := checkTokenType(tokenType, expectToken); err != nil {
				return nil, err
			}
			p.Read()
			expectToken = tokenizer.NULL | tokenizer.STRING |
							tokenizer.NUMBER | tokenizer.BEGIN_OBJECT |
							tokenizer.BEGIN_ARRAY | tokenizer.BOOLEAN
		case tokenizer.SEP_COLON: // never
			panic("error when your parse json array !")
		}
	}

	return nil, errors.ParseError{}
}