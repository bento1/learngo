package something

import "fmt"

func SayHello() { // 패키지에 만든 함수는 대문자로 시작해야 호출이 가능
	fmt.Println("hello")
}
func SayBye() {
	fmt.Println("Bye")
}
