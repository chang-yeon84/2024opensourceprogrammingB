package main

import (
	"fmt"
	"os"
)

// func test(strs string) {
// func test(strs ...string, i int) { // error 가변 매개변수는 반드시 마지막에 와야함
func test(i int, strs ...string) { // 가변 매개변수 예시
	fmt.Println(strs)
}
func main() {
	//fmt.Println(os.Args, len(os.Args), reflect.TypeOf(os.Args))
	slices := os.Args[1:]
	fmt.Println(slices, slices[1])
	test(1, "123")
	test(-99, "abc")
	test(55)
	test(12, "123", "abc", "weq")
}
