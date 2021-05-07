package main

import (
	"fmt"
	"gson/parser"
	"gson/tokenizer"
)

func main() {
	list := tokenizer.GetTokenList("{\n    \"name\": \"qianxi2\",\n    \"age\": 21,\n    \"obj\": {\n        \"name\": \"inner\",\n        \"age\": false\n    },\n    \"sex\": \"male\"\n}")
	parser := parser.GetParser(list)
	object := parser.ParseJsonObject()

	fmt.Println(object)
}