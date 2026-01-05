package JSON

import (
	"encoding/json"
	"fmt"
)

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func ArraySlice() {
	data := []byte(`[
	{"id":1,"name":"abc"},
	{"id":2,"name":"xyz"}
	]`)

	var items []Item
	err := json.Unmarshal(data, &items)
	if err != nil {
		fmt.Println("Error Unmarshalling:", err)
		return
	}
	fmt.Println("Slice of structs:", items)
	for _, item := range items {
		fmt.Printf("ID:%d,Name:%s\n", item.ID, item.Name)
	}
}
