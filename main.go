package main

import (
	"fmt"
	"gson/parser"
	"gson/tokenizer"
)

func main() {
	list := tokenizer.GetTokenList("{\n    \"name\": \"刘毅涵\",\n    \"age\": 21\n}")

	parser := parser.GetParser(list)
	object := parser.ParseJsonObject()
	fmt.Println(object["name"])
	fmt.Println(object["age"])
}