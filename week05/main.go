package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {
	//var i int
	//i = 55

	//var i int = 55
	var f float64 = 1.92

	//f := 1.92
	i := "55"
	fmt.Println(strings.Title("kim hana"))
	fmt.Printf("%f %f\n", f, math.Ceil(f))
	fmt.Println(reflect.TypeOf(f), reflect.TypeOf(i))
	fmt.Println("i is", f)
	fmt.Println("i is", i)
	fmt.Print("i is ", i, "\n")
	fmt.Println("i is", i)
	//fmt.Printf("i is %d\n", i)
	fmt.Printf("i is %s\n", i)

}
