package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func(){
		time.Sleep(time.Duration(2) * time.Second)
		c1 <- "func1: has slept two seconds!"
	}()

	go func(){
		time.Sleep(time.Duration(2) * time.Second)
		c2 <- "func2: has slept two seconds"
	}()

	for i := 0; i < 2; i++{
		select {
			case msg1 := <- c1:
				fmt.Println("received from c1", msg1)
			case msg2 := <- c2:
				fmt.Println("received from c2", msg2)
		}
	}
}
