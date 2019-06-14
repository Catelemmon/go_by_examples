package main

import (
	"errors"
	"fmt"
)

func func1(i int)(int, error){
	if i == 42{
		return -1, errors.New("we can't work with 42")
	}
	return i + 3, nil
}

type argError struct {
	arg int
	prob string
}

// channel
func (e *argError) Error() string{
	return fmt.Sprintf( "%d - %s", e.arg, e.prob)
}

func func2(arg int) (int, error) {
	if arg == 42{
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	for _, i := range []int{7, 42}{
		if r, e := func1(i); e != nil{
			fmt.Println("func1 failed: ", e)
		} else{
			fmt.Println("fun1 worked: ", r)
		}
	}

	for _, i := range []int{7, 42}{
		if r, e := func2(i); e != nil {
			fmt.Println("func2 failed: ", e)
		}else {
			fmt.Println("func2 worked: ", r)
		}
	}

	_, e := func2(42)
	if ae, ok := e.(*argError); ok{
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}