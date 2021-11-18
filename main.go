package main // 메인임을 알려줌

import (
	"fmt" // formatting위한

	"github.com/learngo/something"

	"strings"
)

type person struct {
	name    string
	age     int
	favFood []string // array표시는 []가 먼저나오고 타입
}

func canIDrink_switch(age int) bool {
	// switch age{
	// case 10:
	// 	return false
	// case 18:
	// 	return true
	// }
	// return false
	//ifelse문을 낭비를 막을 수 있고
	switch { // if처럼 변수설정을 할수 있다.
	case age < 18:
		return false
	case age >= 18:
		return true
	}
	return false
}
func canIDrink(age int) bool {

	if koreaAge := age + 2; koreaAge < 18 { //if안에 변수 선언 가능
		return false
	} else {
		return true
	}
}
func multiply(a int, b int) int {
	return a * b
}
func superAdd(numbers ...int) int {
	fmt.Println(numbers)
	fmt.Println("index, number := range numbers")
	for index, number := range numbers { //range는 enumerator임 numbersdml 인덱스와 값을 리턴함

		fmt.Println(index, number)
	}
	fmt.Println("i := 0; i < len(numbers); i++ ")
	for i := 0; i < len(numbers); i++ { //range는 enumerator임 numbersdml 인덱스와 값을 리턴함

		fmt.Println(numbers[i])
	}
	return 1
}

//다양한 것을 리턴하는
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}
func lenAndUpper_naked(name string) (length int, uppercase string) { //naked return
	// length :=1 //error 이미 리턴타입케이스에서 초기화됨
	defer fmt.Println("im done") // 함수가 종료하고 나서 할 작업을 실행
	//api에 응답을 보낸다던가
	// 파일을 생성하고 삭제하던가
	// 여러가지에응용이 가능함
	length = len(name)
	uppercase = strings.ToUpper(name)
	return //len(name), strings.ToUpper(name) //이미 변수로 리턴한다고 자동리턴하게함
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
	totalLength3, upperName3 := lenAndUpper_naked("dongunyun") // = 로하면 안됨

	fmt.Println(totalLength3, upperName3)
	//loop
	superAdd(1, 2, 3, 4, 5, 6, 7)
	fmt.Println(canIDrink(20))
	fmt.Println(canIDrink_switch(20))

	a := 2
	b := a //copy가 일어나고 있다.

	fmt.Println(a, b)
	a = 10 // b에 영향이 없음 // 깊은 복사
	fmt.Println(a, b)
	// 메모리를 보고 싶으면 & 을 붙인다.
	fmt.Println(&a, &b)
	b2 := &a
	fmt.Println(&a, b2) // 주소가 복사되었다.
	fmt.Println(a, b2)  // 주소가 복사되었다.
	//*은 훑어보는거지 주소가 가르키는 valuef를  훑어보다
	fmt.Println(*b2)
	a = 555
	fmt.Println(*b2)
	//반대로 b2를 이용해서 a를 바꿀수 있음
	*b2 = 2
	fmt.Println(a)
	names := [5]string{"a", "b", "c"} //array [길이] 타입 {요소,요소}
	names[2] = "as"                   // 0부터

	fmt.Println(names)
	names2 := []string{"a", "b", "c"} // slice라고 부르는 타입 길이 지정안함 계속 추가가능
	names2 = append(names2, "d")      // 1인자는 slice고
	fmt.Println(names2)
	dongun4 := map[string]string{"name": "dongun", "age": "30"} //map[key의 타입]valeu 타입, 다른타입을 구현하고싶으면 struct를 써야함
	fmt.Println(dongun4)
	for _, value := range dongun4 {
		fmt.Println(value)
	}
	favFood := []string{"kimchi", "ramen"}
	//person_ := person{"dongun", 32, favFood} //보기 어려움
	person_ := person{name: "dongun", age: 18, favFood: favFood}
	fmt.Println(person_)
	//클래스가없고 생성자가 없어서 직접실행시켜야함

}
