package main

import (
	"errors"
	"fmt"

	genericutils "github.com/duyquang6/go-generic-utils"
)

func main() {
	optionExample1()
	eitherExample1()
}

func optionExample1() {
	fmt.Println("-------Option Example 1-------")
	valEx1 := "hello"
	opt1 := genericutils.Some(valEx1)
	fmt.Println(opt1.Get())
	fmt.Println(opt1.IsDefined())
	fmt.Println(opt1.GetOrElse("world"))

	opt2 := genericutils.None[string]()
	fmt.Println(opt2.IsDefined())
	fmt.Println(opt2.GetOrElse("world"))
}

func eitherExample1() {
	fmt.Println("-------Either Example 1-------")
	valEx1 := "hello"
	either1 := genericutils.Right[error, string](valEx1)
	fmt.Println(either1.IsRight())
	fmt.Println(either1.Get())
	fmt.Println(either1.GetOrElse("world"))

	opt2 := genericutils.Left[error, string](errors.New("example"))
	fmt.Println(opt2.IsLeft())
	fmt.Println(opt2.GetOrElse("world"))
}
