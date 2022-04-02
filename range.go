package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("%d => %d\n", i, v)
	}
}

// for ループに利用する range は、スライスや、マップ( map )をひとつずつ反復処理するために使います。
// スライスをrangeで繰り返す場合、rangeは反復毎に2つの変数を返します。 
// 1つ目の変数はインデックス( index )で、2つ目はインデックスの場所の要素のコピーです。