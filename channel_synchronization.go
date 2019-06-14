package main

import (
	"fmt"
	"time"
)

func work(has_done chan bool){
	fmt.Println("working....")
	// do something
	time.Sleep(time.Duration(5)*time.Second)
	fmt.Println("something has done")

	has_done <- true
}


func main() {
	done := make(chan bool, 1)

	go work(done)

	<-done
}
