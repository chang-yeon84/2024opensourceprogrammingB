package main

import "fmt"

func main() {
	var emptySlice []bool
	//emptySlice = make([]bool, 5)
	fmt.Printf("%#v %d\n", emptySlice, len(emptySlice)) // slice zero value (nil), 0
	if len(emptySlice) == 0 {
		emptySlice = append(emptySlice, true)
	}
	fmt.Printf("%#v %d\n", emptySlice, len(emptySlice)) // []bool{true}, 1

	gpas := [5]float64{3.5, 4.1, 4.5, 3.9, 4.23} // array := array literal
	gpa_slice := gpas[1:4]                       // [4.1, 4.5, 3.9]
	//gpa_slice[1] = 2.71
	gpas[2] = 2.71
	//gpa_slice = append(gpa_slice, 4.3)
	gpa_slice = append(gpa_slice, 4.3, 5.55)
	fmt.Println(len(gpa_slice), gpa_slice, gpas)
}
