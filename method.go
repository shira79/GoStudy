package main

import (
	"fmt"
)

//構造体
type Vertex struct {
	X, Y float64
}

// Vertex構造体のメソッド
func (v Vertex) Sub(a float64) float64 {
	return (v.X - v.Y) * a
}

func (v Vertex) Add(a float64) float64 {
	return (v.X + v.Y) * a
}

func main() {
	v := Vertex{10, 4}
	fmt.Println(v.Sub(3))
	fmt.Println(v.Add(3))
}
