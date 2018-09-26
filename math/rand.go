package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(rand.Float64())
	fmt.Println(rand.Float32())
	fmt.Println(rand.Int())
	fmt.Println(rand.Intn(10))
	fmt.Println(rand.Int31n(233))


	fmt.Println(rand.Perm(10))

	rand.Seed(time.Now().UnixNano())
	fmt.Println(time.Now().UnixNano())
	for {
		num := rand.Intn(10)
		if num == 3 {
			break
		}

		fmt.Println(num)
	}


	//r := rand.New(rand.NewSource(233))
	//r.Seed(time.Now().UnixNano())
	//fmt.Println(r.Int())
	//fmt.Println()
}
