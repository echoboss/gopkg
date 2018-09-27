package main

import (
	"time"
	"fmt"
	"encoding/json"
)

const DATETIME_LAYOUT = "2006-01-02 15:04:05"

func main() {
	newYear := time.Date(2019,time.January,1,0,0,0,0,time.UTC)
	printTime(time.Now().Add(1*time.Hour))
	printTime(time.Now().AddDate(1,2,3))
	now := time.Now()

	time.Sleep(50*time.Millisecond)

	fmt.Println(time.Since(now))
	fmt.Println(time.Now().Sub(now))
	fmt.Println(time.Now().UnixNano())


	fmt.Println(time.Parse(DATETIME_LAYOUT,"2018-01-01 00:12:12"))
	d, _ := time.ParseDuration("1h2m22s12ms")
	fmt.Println(d.Hours())
	fmt.Println(d.Minutes())
	fmt.Println(d.Seconds())
	fmt.Println(d.Nanoseconds())
	fmt.Println(d.String())

	fmt.Println(now.Day())

	l,_ := time.LoadLocation("PRC")
	fmt.Println(newYear.In(l))
	fmt.Println(newYear.In(time.Local))
	fmt.Println(newYear.Local())
	fmt.Println(time.Now().Local())
	fmt.Println(time.Now().ISOWeek())
	b, _ := json.Marshal(newYear)
	fmt.Printf("%s",b)
	fmt.Println(time.Now().Zone())


	timer := time.NewTimer(time.Second)
	<- timer.C

	ch1 := make(chan int)

	go func(chan int) {
		time.Sleep(3*time.Second)
		ch1 <- 0
	}(ch1)
	for {
		select {
		case <- ch1:
			fmt.Println("goroutine done")
			goto ForEnd
		case <- time.After(2*time.Second):
			fmt.Println("timeout1")


		}
	}
	ForEnd:


	f := func() {
		fmt.Println("timeout2")
	}
	time.AfterFunc(time.Second,f)
	time.Sleep(5*time.Second)


	ticker := time.NewTicker(time.Second)
	go func() {
		for t := range ticker.C {
			printTime(t)
		}
	}()

	<- time.After(3*time.Second)
	ticker.Stop()


	st := time.Tick(time.Second)

	for t := range st {
		printTime(t)
	}



}

func printTime(t time.Time)  {
	fmt.Println(t.Format(DATETIME_LAYOUT))
}