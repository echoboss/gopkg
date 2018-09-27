package main

import (
	"os"
	"encoding/gob"
	"fmt"
)

func enc()  {
	info := map[string]string{
		"name":  "xichen",
		"age": "24",
	}
	name := "test.gob"
	File, _ := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0644)
	defer File.Close()
	enc := gob.NewEncoder(File)
	if err := enc.Encode(info); err != nil {
		fmt.Println(err)
	}
}
func dec()  {
	var M map[string]string
	File, _ := os.Open("test.gob")
	D := gob.NewDecoder(File)
	D.Decode(&M)
	fmt.Println(M)
}

func main() {
	//enc()
	dec()
}
