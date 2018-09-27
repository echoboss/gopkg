package main

import (
	"encoding/json"
	"fmt"
	"bytes"
)

type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Tags []string `json:"tags"`
}



func main() {
	//u := User{Name:"LuYoo",Age:18,Tags:[]string{"go","python","linux"}}
	//b, _ := json.Marshal(u)
	//b, _ := json.MarshalIndent(u,"","\t")

	//var user User
	//json.Unmarshal(b,&user)
	//fmt.Printf("%+v\n",user)
	//
	//s := `{"name":"LuYoo","age":18}`
	//
	//dst := bytes.NewBuffer(nil)
	//json.Indent(dst,[]byte(s),"","\t")
	//fmt.Println(dst.String())
	//
	//streamJSON := `{"name":"pike","age":20,"tags":["C","C++","Go"]}{"name":"ken","age":20,"tags":["C","C++","Go"]}`
	//
	//dec := json.NewDecoder(strings.NewReader(streamJSON))
	//
	//for {
	//	var uu User
	//	err := dec.Decode(&uu)
	//	if err == io.EOF {
	//		break
	//	}
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	fmt.Printf("%+v\n",uu)
	//}

	byt := bytes.NewBuffer(nil)
	//w := bufio.NewWriter(byt)
	//
	//enc := json.NewEncoder(w)
	//
	//if err := enc.Encode(u);err != nil {
	//	panic(err)
	//}
	//fmt.Println(w.Flush())
	//fmt.Printf("%s",byt.String())

	//json.Compact(byt,b)
	//fmt.Printf("%s",byt.String())


	s := `{"name":"<script>alert('f**k')</script>", "age:1&1"}`
	json.HTMLEscape(byt, []byte(s))
	fmt.Println(byt.String())


}
