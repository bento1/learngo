package main // 메인임을 알려줌

import (
	"fmt" // formatting위한

	"github.com/learngo/something"

	"strings"
)

func multiply(a int, b int) int {
	return a * b
}

//다양한 것을 리턴하는
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}
func repeatMe(works ...string) { //...은 변수를 무제한으로 받을 수 있는 방법
	fmt.Println(works)
}
func main() { // 진입정
	fmt.Println("hello world") // 왜 Println은 대문자로 시작인가? export 를위해서
	something.SayHello()
	//const name= "dongun"// type을 써줘야함
	const name string = "dongun"
	// name="dongun2" 불가능
	fmt.Println(name)
	var name2 string = "dongun"
	name2 = "dongun2"
	fmt.Println(name2) //가능 너무 길기떄문에
	name3 := "dongun"  // go가 type을 찾아준다. 축약형은 var만 가능하다.
	fmt.Println(name3)
	name3 = "dongun2"
	fmt.Println(name3)
	fmt.Println(multiply(2, 2))
	totalLength, upperName := lenAndUpper("dongun") // 사용하지 않으면  에러남  하나만 받으면 안됨
	totalLength2, _ := lenAndUpper("dongun")        // _ 는 변수를 무시하는 방법이다.
	fmt.Println(totalLength, upperName)             // 사용해야함
	fmt.Println(totalLength2)                       // 사용해야함
	repeatMe("1", "2", "3", "4")
}
