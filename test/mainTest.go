package main

import (
	"fmt"
	"lol/entity"
	"reflect"
)

type A struct {
	name string
	age  int
}

func (a A) IsEmpty() bool {
	return reflect.DeepEqual(a, A{})
}

func main() {
	var aa entity.Match
	if aa == (entity.Match{}) {
		fmt.Println("empty")
	}
	fmt.Println(aa)

}
