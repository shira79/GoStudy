package main

import (
	"fmt"
	"math"
)

// 構造体の型だけじゃなくても、任意の型にもメソッドを宣言できる
type MyFloat float64

// 絶対値を返却するメソッド
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
