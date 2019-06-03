package main

import "fmt"

func main() {
	anonymousStruct := struct {
		anonymousField1  string
		anonymousField2  string
		importantMessage string
	}{
		anonymousField1:  "explosive value 1",
		anonymousField2:  "explosive value 2",
		importantMessage: "This struct will self detonate",
	}

	fmt.Println(anonymousStruct.anonymousField1)
	fmt.Println(anonymousStruct.anonymousField2)
	fmt.Println(anonymousStruct.importantMessage)
}
