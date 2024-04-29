package main

import (
	"01/oop1"
	"01/oop2"
)

type Num int

func (x Num) Equal(i int) bool {
	return int(x) == i
}

func (x Num) LessThan(i int) bool {
	return int(x) < i
}

func (x Num) BiggerThan(i int) bool {
	return int(x) > i
}

func (x *Num) Sum(i int) {
	*x = *x + Num(i)
}

func main() {
	//var f1 Num = 6
	//var f2 oop1.NumInterface1 = f1
	//var f3 oop2.NumInterface2 = f2

	var f1 Num = 6
	var f2 oop2.NumInterface2 = f1
	var f3 oop1.NumInterface1 = f2
	println(f3)

}
