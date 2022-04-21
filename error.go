package main

import (
	"fmt"
	"math"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("時間 → %v,  \nメッセージ → %s", e.When, e.What)
}

func MySqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, &MyError{time.Now(), "負の値です"}
	}
	return math.Sqrt(f), nil

}

func main() {
	i, err := MySqrt(3);
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(i)
}
