package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// ポインタレシーバーのメソッド
// ポインタレシーバーであれば、レシーバそのものにを更新できる
func (v *Vertex) Scale1(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// 変数レシーバーのメソッド
// 変数レシーバーは、関数の引数と同じふるまいで、レシーバのコピーを操作する
func (v Vertex) Scale2(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale1(10)
	// v.Scale2(10)
	fmt.Println(v.Abs())
}
