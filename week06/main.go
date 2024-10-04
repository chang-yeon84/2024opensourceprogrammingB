package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := 13
	var f float64 = 12.9 //f := 12.9

	fmt.Printf("value i : %d, value f : %f\n", i, f)
	fmt.Printf("%d * %f = %d\n", i, f, i*int(f))
	fmt.Println(reflect.TypeOf(f), reflect.TypeOf(i))
}
