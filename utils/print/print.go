package print

import (
	"encoding/json"
	"fmt"
)

// Pretty json으로 변경해서 프린트
func Pretty(v interface{}) {
	b, _ := json.MarshalIndent(v, "", "  ")
	println(string(b))
}

// Struct 프린트
func Struct(v interface{}) {
	fmt.Println("____________________________________")
	fmt.Printf("%+v\n", v)
	fmt.Println("____________________________________")
}
