package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

// Go's return values may be named. If so, they are treated as variables defined at the top of the function.
//
// These names should be used to document the meaning of the return values.
//
// A return statement without arguments returns the named return values. This is known as a "naked" return.
//
// Naked return statements should be used only in short functions, as with the example shown here. They can harm readability in longer functions.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// Declare multiple vars (type is inferred)
var c, python, java = true, false, "no!"

// Const
const Pi = 3.14

func main() {
	fmt.Println("Welcome to the playground!")

	fmt.Println("The time is", time.Now())
	fmt.Println("My favorite number is", rand.Intn(10))

	// Call a function
	fmt.Println(add(42, 13))

	// Return two strings swapped
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))

	var e int

	// Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
	//
	//Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
	k := 3
	fmt.Println(e, k, c, python, java)

	fmt.Printf("Type: %T Value: %v\n", python, python)

	// Type conversions
	//The expression T(v) converts the value v to the type T.
	// Unlike in C, in Go assignment between items of different type requires an explicit conversion.
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	// Type inference
	// When declaring a variable without specifying an explicit type (either by using the := syntax or var = expression syntax),
	// the variable's type is inferred from the value on the right hand side.
	// But when the right hand side contains an untyped numeric constant, the new variable may be an int, float64, or complex128 depending on the precision of the constant.
	i := 42           // int
	h := 3.142        // float64
	g := 0.867 + 0.5i // complex128
	fmt.Printf("Var of type %T is value %v\n", i, i)
	fmt.Printf("Var of type %T is value %v\n", h, h)
	fmt.Printf("Var of type %T is value %v\n", g, g)

	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

}

// Go's basic types are
//
//bool
//
//string
//
//int  int8  int16  int32  int64
//uint uint8 uint16 uint32 uint64 uintptr
//
//byte // alias for uint8
//
//rune // alias for int32
//     // represents a Unicode code point
//
//float32 float64
//
//complex64 complex128
