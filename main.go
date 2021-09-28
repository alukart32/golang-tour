package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"runtime"
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
}

func logMsg(msg string) {
	fmt.Println()
	log.Printf("%v\n", msg)
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
