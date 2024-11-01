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

	//counts := 0
	var isprime bool = true //가독성이 좋고 메모리 이득
	j := 2
	for j < n {
		if n%j == 0 {
			//counts = counts + 1
			isprime = false // 더하기 연산 제거
		}
		j++
	}
	//if counts == 0
	if isprime { // == 비교 연산 제거
		fmt.Printf("%d is prime numbers.", n)
	} else {
		fmt.Printf("%d is not prime numbers.", n)
	}
}
