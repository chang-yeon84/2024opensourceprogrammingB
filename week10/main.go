package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
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
	} else {
		j := 2
		for j < n {
			if n%j == 0 {
				isprime = false
				break
			}
			fmt.Printf("%d ", j) //check j loop
			j++
		}
	}

	if isprime {
		fmt.Printf("%d is prime numbers.", n)
	} else {
		fmt.Printf("%d is not prime numbers.", n)
	}
}
