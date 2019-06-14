package main

import "fmt"

func main(){

	massages := make(chan string)

	go func(){
		massages <- "ping baidu.com"
	}()

	msg := <- massages
	fmt.Println(msg)

}