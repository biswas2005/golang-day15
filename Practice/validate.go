package Practice

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func Validate() {
	data := []byte(`[{"id":0,"name":"xyz","age":12},
	{"id":1,"name":"qwe","age":43}]`)

	var u []Person
	err := json.Unmarshal(data, &u)
	if err != nil {
		fmt.Println("Error decoding:", err)
		return
	}

	for i, val := range u {
		if val.ID == 0 {
			fmt.Printf("Invalid ID at index %d :%+v\n", i, val.ID)
			continue
		}
		if val.Name == " " {
			fmt.Printf("Invalid Name at index %d:%+v\n", i, val.Name)
			continue
		}
		if val.Age == 0 {
			fmt.Printf("Invalid age at index %d:%+v\n", i, val.Age)
		}
	}
	fmt.Println("Decoded:", u)

}
