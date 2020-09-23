package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

var c bool //var là 1 chuỗi các biến
var i1, j int = 1, 2

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func add(x, y int) int { //function add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x*0.1
}

func main() {
	fmt.Println("Hello everybody. Today I feel so good!")
	fmt.Println("Random number is: ", rand.Intn(10000000))
	fmt.Println("Pi=", math.Pi)
	fmt.Println("x + y = ", add(10, 20))
	a, b := swap("Hello ", "World ") //gán biến
	fmt.Println(a, b)
	fmt.Println(split(17))
	var i int
	fmt.Println(i, c)
	var c1, python, java = true, false, "no!" //cách 1
	fmt.Println(i1, j, c1, python, java)
	c11, python1, java1 := true, false, "no!"
	fmt.Println(c11, python1, java1)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	const Pi = 3.14
	fmt.Println("Pi = ", Pi)
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
