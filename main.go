package main

import (
	"gopkg.in/ffmt.v1"
	"gson/gson"
	"log"
)

func main() {
	s := "[     {\"Name\": \"Platypus\", \"Order\": \"Monotremata\"},     {\"Name\": \"Quoll\",    \"Order\": \"Dasyuromorphia\"} ]";
	g := &gson.Gson{}
	arr, error := g.ParseJsonArray(s)
	if error != nil {
		log.Fatal(error)
	}
	ffmt.Puts(arr[0])
}