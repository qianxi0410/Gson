package main

import (
	"gopkg.in/ffmt.v1"
	"gson/parser"
	"gson/tokenizer"
)

func main() {
	//list := tokenizer.GetTokenList(" {  \nsuccess : 0,  \nerrorMsg :` \"错误消息\",  \ndata : {  \n\ttotal : \"总记录数\",  \n\trows : [\n\t\t\t {  \n\t\t\t\tid : \"任务ID\",  \n\t\t\t\tworkName : \"任务名称\",  \n\t\t\t\tassigneeName : \"经办人姓名\",  \n\t\t\t\tname : \"流程步骤名称\",  \n\t\t\t\tprocessInstanceInitiatorName : \"发起人\",  \n\t\t\t\tprocessInstanceStartTime : \"发起时间\",  \n\t\t\t\tcreateTime : \"到达时间\",  \n\t\t\t\tdueDate : \"截止时间\"  \n\t\t\t}, \n\t\t\t{  \n\t\t\t\tid : \"ID\",  \n\t\t\t\tworkName : \"名称\",  \n\t\t\t\tassigneeName : \"经办人\",  \n\t\t\t\tname : \"流程\",  \n\t\t\t\tprocessInstanceInitiatorName : \"发起人\",  \n\t\t\t\tprocessInstanceStartTime : \"发起\",  \n\t\t\t\tcreateTime : \"到达\",  \n\t\t\t\tdueDate : \"截止\"  \n\t\t\t} \n\t\t]  \n\t}  \n}")
	list := tokenizer.GetTokenList("{\n    \"name\": \"三班\",\n    \"students\": [\n        {\n            \"age\": 25,\n            \"gender\": \"female\",\n            \"grades\": \"三班\",\n            \"name\": \"露西\",\n            \"score\": {\n                \"网络协议\": 98,\n                \"JavaEE\": 92,\n                \"计算机基础\": 93\n            },\n            \"weight\": 51.3\n        },\n        {\n            \"age\": 26,\n            \"gender\": \"male\",\n            \"grades\": \"三班\",\n            \"name\": \"杰克\",\n            \"score\": {\n                \"网络安全\": 75,\n                \"Linux操作系统\": 81,\n                \"计算机基础\": 92\n            },\n            \"weight\": 66.5\n        },\n        {\n            \"age\": 25,\n            \"gender\": \"female\",\n            \"grades\": \"三班\",\n            \"name\": \"莉莉\",\n            \"score\": {\n                \"网络安全\": 95,\n                \"Linux操作系统\": 98,\n                \"SQL数据库\": 88,\n                \"数据结构\": 89\n            },\n            \"weight\": 55\n        }\n    ]\n}")
	parser := parser.GetParser(list)
	object := parser.ParseJsonObject()
	ffmt.Puts(object)
}