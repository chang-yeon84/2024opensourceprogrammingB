package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	fmt.Printf("오늘은 %d 년 %d 월 %d 일\n", now.Year(), int(now.Month()), now.Day())
	fmt.Printf("지금은 %d 시 %d 분 %d 초\n", now.Hour(), now.Minute(), now.Second())

}
