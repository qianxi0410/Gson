package gson

import "gson/parser"

type Gson struct {

}

func (g Gson) ParseJsonObject(jsonStr string) (map[string]interface{}, error) {
	p := parser.GetParser(jsonStr)
	return p.ParseJsonObject()
}

func (g Gson) ParseJsonArray(jsonStr string) ([]interface{}, error) {
	p := parser.GetParser(jsonStr)
	return p.ParseJsonArray()
}