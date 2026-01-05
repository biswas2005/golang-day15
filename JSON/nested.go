package JSON

import (
	"encoding/json"
	"fmt"
)

type Profile struct {
	Age  int    `json:"age"`
	City string `json:"city"`
}
type Person struct {
	ID    int     `json:"id"`
	About Profile `json:"about"`
}

func Nested() {

	data := (`{
	"id":1,"about":{
	"age":20,"city":"hyd"}
	}`)

	var nests Person
	err := json.Unmarshal([]byte(data), &nests)
	if err != nil {
		fmt.Println("Error Unmarshalling:", err)
		return
	}
	fmt.Println("Data:", nests)

	fmt.Printf("Person ID:%d\n", nests.ID)
	fmt.Printf("Person age:%d\n", nests.About.Age)
	fmt.Printf("Person City:%s\n", nests.About.City)

}
