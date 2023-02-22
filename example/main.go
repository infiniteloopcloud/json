package main

import (
	"fmt"

	"github.com/infiniteloopcloud/json"
)

type nullString struct {
	Valid  bool
	String string
}

func (n nullString) IsZero() bool {
	return !n.Valid
}

type RandomDataSet struct {
	Data1 nullString `json:"data1,omitempty"`
	Data2 nullString `json:"data2,omitempty"`
}

func main() {
	result, err := json.Marshal(RandomDataSet{Data1: nullString{Valid: false, String: "test"}})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(result))
	// Output: {}

	result, err = json.Marshal(RandomDataSet{Data1: nullString{Valid: true, String: "test"}})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(result))
	// Output: {"data1":{"Valid":true,"String":"test"}}
}
