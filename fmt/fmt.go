package main

import (
	"fmt"
)

func main() {

	err := fmt.Errorf("open file failed.")
	fmt.Printf("%T %[1]v \n",err)
	//f, _ := os.OpenFile("a.txt",os.O_CREATE|os.O_APPEND|os.O_RDWR,0644)

	//fmt.Fprintln(f,"abc")
	//defer f.Close()
	//
	var name string
	var age int
	//fmt.Fscan(f,&name,&age)
	//fmt.Println(name,age)

	//var mounth, day int
	//n, err := fmt.Scanf("%d,%d",&mounth,&day)
	//fmt.Println(n,err,mounth,day)

	//fmt.Println(fmt.Sprintf("name: %s age: %s",name,age))

	str := "name:LuYoo age:18"
	fmt.Sscanf(str,"name:%s,age:%d",&name,&age)
	fmt.Println("parse",name,age)



}
