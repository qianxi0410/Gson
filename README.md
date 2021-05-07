# Gson
A json parser written in golang(jsut for fun🆒)

## Data type mapping 
this is type mapping about json-type to go-type 

`number` => `float64`

`string` => `string`

`null` => `nil`

`array` => `[]interface{}`

`object` => `map[string][]interface{}`

`bool` => `bool`

so when you parse a json string like tihs:
```json
{
  // ...
}
```
you will get a map.

or when you parse a json string like tihs :
```json
[
  // ...
]
```
you will get a slice.

## usage
```go
func main() {
	s := "{// ...}" // your json string
	g := &gson.Gson{}
	object, _ := g.ParseJsonObject(s)
	ffmt.Puts(object)
}
```
or 
```go
func main() {
	s := "[// ...]" // your json string
	g := &gson.Gson{}
	arr, _ := g.ParseJsonArray(s)
	ffmt.Puts(arr)
}
```

## license
[MIT](./LICENSE)