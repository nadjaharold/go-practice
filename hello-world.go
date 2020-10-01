package main

import "fmt"

func main() {
	fmt.Println("hello world")

	fmt.Println("go" + "lang")

	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	// int型の変数nを定義・代入
	var n int
	n = 5
	fmt.Printf("n=%d\n", n)

	// int型の変数x,y,zを定義・代入
	// var x, y, z int
	// x, y, z = 1, 3, 5

	// int型の変数x,yとstring型のnameを定義
	// var (
	// 	a, b int
	// 	name string
	// )

	// i := 1
	// b1 := true
	// f := 3.14
	// s := "abc"

	// var i1 = 1

	// var (
	// 	b2 = true
	// 	f1 = 3.14
	// 	s1 = "abc"
	// )

	// a := [5]int{1, 2, 3, 4, 5} // a == "[1, 2, 3, 4, 5]"
	// b := [5]int{1, 2, 3}       // b == "[1, 2, 3, 0, 0]"
	// c := [5]int{}              // c == "[0, 0, 0, 0, 0]"

	// var d [5]int // d == "[0, 0, 0, 0, 0]"

	// e := [...]int{1, 2, 3}       // e == "[1, 2, 3]"
	f := [...]int{1, 2, 3, 4, 5} // e == "[1, 2, 3, 4, 5]"

	fmt.Println(f[0])
	f[1] = 10
	fmt.Println(f[1])

	var x interface{}
	x = 1
	// x = 3.14
	// x = "hello"
	x = [...]float64{1.2, 2.1, 3.5}
	fmt.Println(x)
	i, isString := x.([3]float64)
	fmt.Println(i)
	fmt.Println(isString)

	if x, y := 1, 2; x < y {
		fmt.Println("y is greater than x.", y, x)
	}
	for i := 0; i < 10; i++ {
		fmt.Println(i)

	}
	/*
	   go文は並列処理を司る特別な機能である。
	   Goはスレッドよりも小さい処理単位であるゴルーチン(goroutine）が並行して動作するように実装されている。
	    go文はこのゴルーチンを新たに生成して並行して処理される新しい処理の流れをランタイムに追加するための機能である。
	*/
	// go sub() // start goroutine
	// for i := 0; i < 10; i++ {
	// 	fmt.Println("main loop")
	// }
	s2 := make([]int, 3)
	s2[0] = 5
	fmt.Println(s2) // => "[5, 0, 0]"

	s3 := []int{1, 2, 3, 4, 5}
	s4 := s3[0:3]       // => "[1, 2, 3]"
	s5 := s3[len(s3)-2] // => "[4, 5]"
	fmt.Println(s4)
	fmt.Println(s5)
	s3 = append(s3, 5, 4, 6, 7, 8)
	fmt.Println(s3)

	a := []int{10, 2, 30, 4, 5, 62, 7, 8, 9, 19}

	z1 := a[2:7] // == "[3, 4]"
	fmt.Println(z1)

	fmt.Println(len(z1))
	fmt.Println(cap(z1))

	// パッケージに定義された定数、変数、関数などが他のパッケージから参照可能であるかは、識別子の1文字目が大文字であるかどうかで決定される。
}
func sub() {
	for i := 0; i < 100; i++ {
		fmt.Println("sub loop")
	}
}

// go run hello-world.go
// go build hello-world.go
// ./hello-world
