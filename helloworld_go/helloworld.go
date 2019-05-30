package main

import (
	"os"
	"fmt"
	"github.com/ohhyel/stringutil_go"
)


type person struct {
	minAge int
	name string
	age int
}

func (p person) Hello()  {
	fmt.Printf("Hello. My name is %s\n", p.name)
}

// function (receiver) FunctionName (param) returnType {}
func (p person) MyAge(){
	fmt.Printf("My age is %d\n", p.age)
}

func main() {
	fmt.Println("hello world!!")
	str := stringutil_go.Reverse("TEST REVERSE!!")
	fmt.Println(str)

	// struct ( 생성자가 없어서 초기화 하여 사용 하거나 New 함수를 만들어서 사용
	john := person{minAge:0, name:"john doe", age: 20}
	john.Hello()
	john.MyAge()

	// go routine
	for i := 0; i < 10 ; i++ {
		go goRoutine(i)
	}

	// go routine test
	var input string
	fmt.Scanln(&input)




	// error test
	f, error := os.Open("filename.txt")
	if error != nil {
		fmt.Println("File oepn error : ", error.Error())
		os.Exit(1)
	}
	fmt.Println("File oepn: ", f.Name())

}
func goRoutine(id int) {
	for i := 0; i< 10 ; i++  {
		fmt.Println("id: ", id , " - ", i)
	}
}
