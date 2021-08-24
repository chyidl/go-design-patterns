package main

import (
	"encoding/json"
	"fmt"
)

type MyObject struct {
	Number int    `json:"number"`
	Word   string `json:"word"`
}

func main() {
	packt := "packt"
	// convering json counterpart
	jsonPackt, ok := json.Marshal(packt)
	if ok != nil {
		panic("Could not marshal object")
	}
	fmt.Println(string(jsonPackt))

	object := MyObject{5, "Packt"}
	oJson, _ := json.Marshal(object)
	fmt.Printf("%s\n", oJson)

	jsonBytes := []byte(`{"number":5, "string":"Packt"}`)
	var object1 MyObject
	err := json.Unmarshal(jsonBytes, &object1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Number is %d, Word is %s\n", object1.Number, object1.Word)

	var dangerousObject map[string]interface{}
	err = json.Unmarshal(jsonBytes, &dangerousObject)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Number is %d, ", dangerousObject["number"])
	fmt.Printf("Word is %s\n", dangerousObject["string"])
	fmt.Printf("Error reference is %v\n", dangerousObject["nothing"])
}
