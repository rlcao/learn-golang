package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Name string `json:"name"`
	Body string
	Time int64
	Test string
}

func main() {
	msg := Message{"Test","This is a test message",1294706395881547000,"antest"}

	b,err := json.Marshal(msg)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n",string(b))

	var m interface{}
	err = json.Unmarshal([]byte(`{"Name":123}`),&m)
	if err != nil {
		fmt.Println(err)
	}
	ms := m.(map[string]interface{})
	val := ms["Name"].(float64)
	fmt.Println(val)


	f := float64(2.34)

	i := int(f)

	fmt.Printf("%d",i)
}
