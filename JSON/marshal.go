package JSON

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name     string `json:"name,omitempty"`
	Age      int    `json:"age"`
	Adult    bool   `json:"adult"`
	Password string `json:"-"`
}

func Marshal() {

	user := User{Name: "", Age: 20, Adult: false, Password: "qwert"}

	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error Marshaling:", err)
		return
	}
	fmt.Println(string(data))

	var u User
	err1 := json.Unmarshal(data, &u)
	if err1 != nil {
		fmt.Println("Error Unmarshaling:", err1)
		return
	}
	fmt.Println("Unmarshaling Data:", u)
}
