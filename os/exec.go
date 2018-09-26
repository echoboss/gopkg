package main


import (
	"os/exec"
	"fmt"
	"bytes"
)

func main() {
	//args := []string{"en0",}
	//cmd := exec.Command("ifconfig",args...)
	//out, err := cmd.Output()
	//fmt.Println(string(out),err)


	c := exec.Command("echo","abc")
	var buf bytes.Buffer
	c.Stdout = &buf
	fmt.Println(c.Run())
	fmt.Println(buf.String())

	c.Start()
	c.Wait()
	fmt.Println(exec.LookPath("go"))
}
