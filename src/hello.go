// hello.go
package main

import (
	"time"
	"fmt"
	"calculate"
)

//struct
type Student struct{
	age int
}

func (s *Student) get() int{
	return s.age
}

func (s *Student) put(v int){
	s.age = v
}

//interface
type I interface{
	get() int
	put(int)
}

func myTest(i I){
	i.put(10)
	fmt.Println("i.get=",i.get())
}

func ready(name string,n int){
	time.Sleep(time.Second * time.Duration(n));
	fmt.Println(name," is ok!!!");
}

func main() {

	data := 5
	
	fmt.Println("HanNuo(",data,")=",calculate.HanNuo(data))
	
	var person Student
	myTest(&person)
	
	go ready("Coffee",2)
	go ready("Tea",3)
	fmt.Println("I am waiting!")
	time.Sleep(time.Second * 5)
	

}