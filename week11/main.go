package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isPrime(n int) bool {
	if n <= 1 {
		//isPrime = false
		return false
	} else if n == 2 {
		//isPrime = true
		return true
	} else if n%2 == 0 {
		//isPrime = false
		return false
	} else {
		j := 3
		for j*j <= n {
			if n%j == 0 {
				// isPrime = false
				// break
				return false
			}
			//fmt.Printf("%d ", j)
			j = j + 2
		}
	}
	return true
}

func getInteger() int {
	in := bufio.NewReader(os.Stdin)
	a, err := in.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	a = strings.TrimSpace(a)
	number, err := strconv.Atoi(a)
	if err != nil {
		log.Fatal(err)
	}
	return number

}
func main() {
	fmt.Print("Input First(start) number : ")
	n1 := getInteger()
	fmt.Print("Input second(end)number : ")
	n2 := getInteger()
	for i := n1; i <= n2; i++ {
		if isPrime(i) {
			fmt.Printf("%d ", i)
		}
	}
}
