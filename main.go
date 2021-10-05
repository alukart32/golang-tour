package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"runtime"
	"strings"
)

var (
	paramOne bool   = false
	paramTwo bool   = false
	ToBe     bool   = false
	MaxInt   uint64 = 1<<64 - 1
)

func main() {

	fmt.Println("Part One")
	stepOne()
	stepTwo()
	stepThree(64)
	fmt.Printf("%d\n", stepFour(5, 6))

	o1, o2 := stepFive("hello", "world")
	fmt.Println(o1, o2)

	fmt.Println(stepSix(128))

	logMsg("Step 07: vars")
	var k int = 196
	fmt.Println(k, paramOne, paramTwo)

	logMsg("Step 07: variables with initializers")
	var a, b, c = 15, "test_1", true
	fmt.Println(a, b, c)

	// Outside a function := construct is not available
	logMsg("Step 08: variables with short initializers")
	a1, b1, c1 := 143, "test_2", 2.345
	fmt.Println(a1, b1, c1)

	// Basic types
	/*
		bool
		string

		int  int8  int16  int32  int64
		uint uint8 uint16 uint32 uint64 uintptr

		byte - alias for uint8

		rune - alias for int32 represents a Unicode code point

		float32 float64

		complex64 complex128
	*/
	logMsg("Step 09: basic types")
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)

	// The expression T(v) converts the value v to the type T
	logMsg("Step 10: type conversions")
	i := 42
	f := float64(i)
	u := uint(f)
	fmt.Printf("Type: %T Value: %v\n", u, u)

	/*
	 When the right hand side of the declaration is typed, the new variable is of that same type:

	 var i int
	 j := i // j is an int

	 But when the right hand side contains an untyped numeric constant,
	 the new variable may be an int, float64, or complex128 depending on the precision of the constant:

	 i := 42           // int
	 f := 3.142        // float64
	 g := 0.867 + 0.5i // complex128
	*/
	logMsg("Step 11: type inference")
	step_11 := 0.867 + 0.5i // change me!
	fmt.Printf("step_11 is of type %T\n", step_11)

	// Constants cannot be declared using the := syntax. Only =
	logMsg("Step 12: consts")
	const const1 = 25
	fmt.Printf("Const: %v\n", const1)

	fmt.Println("\nPart Two")

	for i := 0; i < 5; i++ {
		fmt.Printf("%v\n", i)
	}

	logMsg("If statement")
	fmt.Println(pow(3, 2, 10))

	fmt.Print("\nGo runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	// logMsg("Stacking defers")
	// fmt.Println("counting")
	// for i := 0; i < 10; i++ {
	// 	defer fmt.Println(i)
	// }
	// fmt.Println("done")

	fmt.Println("\nPart Three")
	logMsg("Pointers")
	var p *int = &i // p := &i
	fmt.Printf("Pointer p value: %v", *p)

	valOne, valTwo := 42, 2701

	valPoint := &valOne    // point to i
	fmt.Println(*valPoint) // read i through the pointer
	*valPoint = 21         // set i through the pointer
	fmt.Println(valOne)    // see the new value of i

	valPoint = &valTwo         // point to j
	*valPoint = *valPoint / 37 // divide j through the pointer
	fmt.Println(valTwo)        // see the new value of j

	logMsg("Struct")
	type Vertex struct {
		x int
		y int
	}

	fmt.Println(Vertex{10, 26})

	v := Vertex{44, 57}
	v.x = 90
	fmt.Println(v)

	vPoint := &v
	vPoint.y = 256 // (*vPoint).y = 256
	fmt.Println(*vPoint)

	v2 := Vertex{x: 1} // struct literals
	fmt.Println(v2)

	logMsg("Arrays")
	array := [6]int{1, 2, 3, 4, 5, 6} // var array[6]int
	fmt.Println(array)

	logMsg("Slices")
	slices := array[2:5]
	fmt.Println(slices)

	slices[0] = 100
	fmt.Println(array)

	slice1 := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(slice1)

	slice2 := []bool{true, false, true, true, false, true}
	fmt.Println(slice2)

	slice3 := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(slice3)

	logMsg("Slice with make")
	slice_a := make([]int, 5)
	printSlice("slice_a", slice_a)

	slice_b := make([]int, 0, 5)
	printSlice("slice_b", slice_b)

	slice_c := slice_b[:2]
	printSlice("slice_c", slice_c)

	slice_d := slice_c[2:5]
	printSlice("slice_d", slice_d)

	logMsg("Slices of slices")
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	logMsg("Appending to a slice")
	var s []int
	printSlice("void_slice", s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice("append_slice_0", s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice("append_slice_1", s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice("append_slice_arr", s)

	Pic(3, 2)

	logMsg("Map")
	// The zero value of a map is nil. A nil map has no keys, nor can keys be added
	map_vert := make(map[string]Vertex) // var map_vert := map[string]Verttex
	map_vert["key_1"] = Vertex{1, 2}
	fmt.Printf("\nMap value: %v\n", map_vert["key_1"])

	map_big := map[string]Vertex{
		"key_1": Vertex{1, 2},
		"key_2": Vertex{3, 4},
	}

	fmt.Printf("\nBig Map key: %v, value: %v\n", "key_2", map_big["key_2"])

	logMsg("Mutating Maps")
	map_big["key_3"] = Vertex{4, 5}
	map_big["key_zero"] = Vertex{0, 0}

	fmt.Println("Big Map: ", map_big)

	delete(map_big, "key_zero")

	fmt.Println("Big Map after delete by key 'key_zero': ", map_big)

	elem, ok := map_big["key_2"]
	fmt.Printf("Check Big Map by key: %s, elem: %v, ok: %v\n", "key_2", elem, ok)

	elem2, ok2 := map_big["key_no"]
	fmt.Printf("Check Big Map by key: %s, elem: %v, ok: %v\n", "key_no", elem2, ok2)
}

func logMsg(msg string) {
	fmt.Println()
	log.Printf("%v\n", msg)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func Pic(dx, dy int) [][]uint8 {
	matrix := [][]uint8{}

	for i := 0; i < dy; i++ {
		matrix = append(matrix, make([]uint8, dx))
	}

	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			matrix[i][j] = uint8(i * j)
		}
	}
	return matrix
}

func stepOne() {
	logMsg("Step 01: print string")
	fmt.Println("Hello, world!!!")
}

func stepTwo() {
	logMsg("Step 02: rand number")

	rand.Seed(1 ^ 32 - 1)
	fmt.Println("Rand number: ", rand.Intn(10))
}

func stepThree(v int) {
	logMsg("Step 03: math package")
	fmt.Printf("Sqrt of %d is %g\n", v, math.Sqrt(float64(v)))
}

func stepFour(x, y int) int {
	logMsg("Step 04: func with args")
	return x * y
}

func stepFive(o1, o2 string) (string, string) {
	logMsg("Step 05: func with multiple returns")
	return o2, o1
}

func stepSix(z int) (x, y int) {
	logMsg("Step 06: func with named return values")
	x = z*2 + z/2
	y = x/2 + 1
	return
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}
