package main

import (
	"fmt"
	"reflect"
)

type person struct {
	name string
	age int
}

func (p *person) somemethod(){
	fmt.Println(p.age)
	fmt.Println(*p)
}
func main() {
	var alex person
	alex.somemethod()
	num := new(int)
	fmt.Println(*num)
	i := make([]int, 2, 2)
	i = append(i, 3)
	fmt.Println(len(i))
	fmt.Println(cap(i))
	fmt.Println(i[2])
	var a int = 25
	//fmt.Println(a.(int))
	//fmt.Println(interface{}(a).(type))
	aa := reflect.TypeOf(a)
	bb := reflect.ValueOf(a)
	fmt.Println(aa,bb)
	T := make([]int,10)
	fmt.Println(T)

}