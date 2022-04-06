package main

import "fmt"

type Person interface {
	introduceMyself()
}

type shira struct{
	name string
}

type kishi struct{
	name string
}

type wasa struct{
	name string
}

func (s shira) introduceMyself() {
	fmt.Println(s.name)
}

func (k kishi) introduceMyself() {
	fmt.Println(k.name)
}

func (w wasa) introduceMyself() {
	fmt.Println(w.name)
}

func main (){
	list := []Person{wasa{name:"おおさわ"}, shira{name:"しらかわ"}, kishi{name:"きし"}}

	for _, v := range list {
		v.introduceMyself()
	}

}