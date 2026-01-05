package Practice

import (
	"encoding/json"
	"fmt"
)

func Dynamic() {

	data := []byte(`{
	"ID":1,
	"name":"abde",
	"active":true,
	"skills":["Go","Rust","Python"],
	"profile":{"age":21,"city":"delhi"}
	}`)

	var result map[string]interface{}
	err := json.Unmarshal(data, &result)
	if err != nil {
		fmt.Println("Error Unmarshaling:", err)
		return
	}
	fmt.Println("Decoded:", result)

	fmt.Println("ID:", result["ID"])
	fmt.Println("Name:", result["name"])
	fmt.Println("Active:", result["active"])

	if profile, ok := result["profile"].(map[string]interface{}); ok {
		fmt.Println("Age:", profile["age"])
		fmt.Println("City:", profile["city"])
	}

	if skills, ok := result["skills"].([]interface{}); ok {
		fmt.Println("Skills")
		for _, s := range skills {
			fmt.Println("--", s)
		}
	}
}
