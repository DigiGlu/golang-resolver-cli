package main

// import resolver "github.com/DigiGlu/golang-json-schema-ref-resolver/resolver"
import resolver "../golang-json-schema-ref-resolver/resolver"

func main() {
	// inJSON := []byte(`{"person": { "properties": { "firstName": { "type": "string" }, "lastName": { "type": "string" }}}, "employee": { "properties": { "$ref": "#/person/properties", "salary": { "type": "number" }}}}`)

	var jsonBytes = []byte(`
	{
		"key1":{
			"Item1": "Value1",
			"Item2": 10 },
		"key2":{
			"Item1": "Value2",
			"Item2": 20 },
		"key3":{
			"Item1": "Value3",
			"Item2": 4711 }
	}`)

	resolver.Dereference(jsonBytes)
}
