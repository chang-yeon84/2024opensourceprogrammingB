package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	answer := rand.Intn(6) + 1 // 범위 1~6
	fmt.Println(answer)

	for guesses := 0; guesses < 3; guesses++ {

		fmt.Printf("%d번의 기회가 남았습니다. 숫자 입력 : ", 3-guesses)
		in := bufio.NewReader(os.Stdin)
		input, err := in.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(guess)

		if answer == guess {
			fmt.Println("정답")
			break
		} else if answer > guess {
			fmt.Println("입력하신 값은 정답보다 작은 수 입니다.")
		} else {
			fmt.Println("입력하신 값은 정답보다 큽니다.")

		}
	}
}
