package main

import (
	"gopkg.in/ffmt.v1"
	"gson/gson"
)

func main() {
	s := "{\n    \"name\": \"三班\",\n    \"students\": [\n        {\n            \"age\": 25,\n            \"gender\": \"female\",\n            \"grades\": \"三班\",\n            \"name\": \"露西\",\n            \"score\": {\n                \"网络协议\": 98,\n                \"JavaEE\": 92,\n                \"计算机基础\": 93\n            },\n            \"weight\": 51.3\n        },\n        {\n            \"age\": 26,\n            \"gender\": \"male\",\n            \"grades\": \"三班\",\n            \"name\": \"杰克\",\n            \"score\": {\n                \"网络安全\": 75,\n                \"Linux操作系统\": 81,\n                \"计算机基础\": 92\n            },\n            \"weight\": 66.5\n        },\n        {\n            \"age\": 25,\n            \"gender\": \"female\",\n            \"grades\": \"三班\",\n            \"name\": \"莉莉\",\n            \"score\": {\n                \"网络安全\": 95,\n                \"Linux操作系统\": 98,\n                \"SQL数据库\": 88,\n                \"数据结构\": 89\n            },\n            \"weight\": 55\n        }\n    ]\n}"
	g := &gson.Gson{}
	object, _ := g.ParseJsonObject(s)
	ffmt.Puts(object)
}