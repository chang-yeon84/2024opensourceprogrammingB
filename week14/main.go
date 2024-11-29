package main

import "fmt"

type student struct {
	id   int
	name string
	gpa  float32
}

func main() {
	var student1 student
	student1.id = 20241111
	student1.name = "heung min"
	fmt.Println(student1.id)
	student1.gpa = 4.5
	var student2 student
	student2.id = 20241234
	student2.name = "seo jihey"
	student2.gpa = 4.4
	fmt.Println(student2.gpa)
}
