package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	//fmt.Printf("%f\n", math.Sqrt(16.0))
	fmt.Print("점수 입력 : ")
	in := bufio.NewReader(os.Stdin)
	i, err := in.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	i = strings.TrimSpace(i)
	n, err := strconv.Atoi(i)
	if err != nil {
		log.Fatal(err)
	}

	var isprime bool = true

	if n <= 1 {
		isprime = false
	} else if n == 2 {
		isprime = true
	} else if n%2 == 0 { //2를 제외한 짝수는 모두 소수가 아님
		isprime = false
	} else { //add number
		j := 3 // change start value
		for j <= int(math.Sqrt(float64(n))) {
			if n%j == 0 {
				isprime = false
				break
			}
			fmt.Printf("%d ", j) //check j loop
			j = j + 2
		}
	}

	if isprime {
		fmt.Printf("%d is prime numbers.", n)
	} else {
		fmt.Printf("%d is not prime numbers.", n)
	}
}
