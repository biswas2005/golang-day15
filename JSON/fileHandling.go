package JSON

import (
	"encoding/json"
	"fmt"
	"os"
)

type Folder struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	ID   int    `json:"id"`
}

func FileHandling() {

	folds := []Folder{{Name: "xyz", Age: 18, ID: 01},
		{Name: "abc", Age: 19, ID: 02}}
	file, err := os.Create("user.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	encode := json.NewEncoder(file)
	encode.SetIndent("", "  ")
	err = encode.Encode(&folds)
	if err != nil {
		fmt.Println("Error encoding:", err)
		return
	}

	file2, err := os.Open("user.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file2.Close()

	var decoded []Folder
	decoders := json.NewDecoder(file2)
	err = decoders.Decode(&decoded)
	if err != nil {
		fmt.Println("Error Decoding:", err)
		return
	}

	fmt.Println("All users:")
	for _, f := range folds {
		fmt.Printf("Name:%s,Age:%d,ID:%d\n", f.Name, f.Age, f.ID)

	}
}
