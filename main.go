package main

import (
	"fmt"
	"gson/parser"
	"gson/tokenizer"
)

func main() {
	list := tokenizer.GetTokenList("{\n    \"name\": \"qianxi2\",\n    \"age\": 21,\n    \"obj\": {\n        \"name\": \"inner\",\n        \"age\": false\n    }\n}")
	//list := tokenizer.GetTokenList("{\n \"obj\": {\n        \"name\": \"inner\",\n        \"age\": false\n    }\n}")
	//list.ToString()
	parser := parser.GetParser(list)
	object := parser.ParseJsonObject()

	fmt.Println(object["name"])
	fmt.Println(object["age"])
	fmt.Println(object["obj"])
	fmt.Println(object)
}