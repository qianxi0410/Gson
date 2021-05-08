package gson

import "gson/parser"

type Gson struct {

}

func (g Gson) ParseJsonObject(jsonStr string) (map[string]interface{}, error) {
	p, err := parser.GetParser(jsonStr)

	if err != nil {
		return nil, err
	}
	return p.ParseJsonObject()
}

func (g Gson) ParseJsonArray(jsonStr string) ([]interface{}, error) {
	p, err := parser.GetParser(jsonStr)
	if err != nil {
		return nil, err
	}
	return p.ParseJsonArray()
}