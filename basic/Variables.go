package main

import (
	"fmt"
	"os"
	"strconv"
)

// IMP --> for every declared type variable there should be used in later code else it will give compile error
// you van use _ = variable_name  to remove error while not using variable
func main() {

	declaration()

	multiAssignment()

	typeAndZero()

	fmtApplication()

	constants()

	iotaPractical()

	dataType()

	numConvStr()

	goType()

	commandLineArguments()
}

func declaration() {
	var age int = 45

	var name string = "Jack"

	var isMarried bool = true

	fmt.Println("age: ", age)

	fmt.Println("name: ", name)

	fmt.Println("married status: ", isMarried)

	// method-2 automatic take variable type
	var age1 = 36

	var name1 = "Jack"

	fmt.Println("age1: ", age1)

	fmt.Println("name1: ", name1)

	//shortcut to declare variable
	//age1 :=35  --> this will not work this method can be used while declaring new variable
	age3 := 34

	name3 := "Jack"

	fmt.Println("age3: ", age3)

	fmt.Println("name3: ", name3)
}

// ///////////////////////////////////////////////////
// MULTIPLE ASSIGNMENT
// //////////////////////////////////////////////////
func multiAssignment() {

	// Method-1
	var (
		age4       int
		name4      string
		isMarried4 bool
	)

	age4, name4 = 23, "Shivam"

	fmt.Println("age4: ", age4)

	fmt.Println("name4: ", name4)

	fmt.Println("isMarried4: ", isMarried4) // this will take default value

	// Method-2 (this method will not take default value
	age5, name5, isMarried5 := 23, "shivam5", false

	fmt.Println("age5: ", age5)

	fmt.Println("name5: ", name5)

	_ = isMarried5 // muting variable

	//we can also perform expression while assigning variable
	sum := 5 - 3.2

	fmt.Println(sum)

	//swapping of two variable
	num1, num2 := 3, 4

	fmt.Println(num1, num2)

	num2, num1 = num1, num2

	fmt.Println(num1, num2)
}

// ////////////////////////////////////////////////
// Types and default value
// ////////////////////////////////////////////////
func typeAndZero() {

	// you must provide a type for each variable you declare or Go should be able to infer it
	var a int = 10

	var b = 15.5 // type inference (deduction)

	c := "Gopher" // short declaration, type inference

	_, _, _ = a, b, c // Blank Identifier (_) to get rid of unused variable error

	// Go is a Statically and Strong Typed Programming Language
	// a = 3.14 -> error. A variable cannot change it's type
	// a = b    -> error. It's not allowed to assign a type to another type  *ANS--> a = int(b)

	//** DEFAULT VALUES **//

	// An uninitialized variable or empty variable  will get the so called ZERO VALUE
	// The zero-value mechanism of Go ensures that a variable always holds a well defined value of its type
	var value int // initialized with 0

	var price float64 // initialized with 0.0

	var name6 string // initialized with empty string -> ""

	var done bool // initialized with false

	fmt.Println(value, price, name6, done) // -> 0 0.0 ""  false
}

// //////////////////////////////////////////
// FMT PACKAGE
// //////////////////////////////////////////
func fmtApplication() {
	//** fmt.Printf() **//

	// fmt.Printf() prints out to stdout according to a format specifier called verb.
	// It doesn't add a newline (\n)

	// VERBS:
	// %d -> decimal
	// %f -> float
	// %s -> string
	// %q -> double-quoted string
	// %v -> value (any)
	// %#v -> a Go-syntax representation of the value
	// %T -> value Type
	// %t -> bool (true or false)
	// %p -> pointer (address in base 16, with leading 0x)
	// %c -> char (rune) represented by the corresponding Unicode code point

	a, b, c := 10, 15.5, "Gophers"

	grades := []int{10, 20, 30}

	fmt.Printf("a is %d, b is %f, c is %s \n", a, b, c) // => a is 10, b is 15.500000, c is Gophers

	fmt.Printf("%q\n", c) // => "Gophers"

	fmt.Printf("%v\n", grades) // => [10 20 30]

	fmt.Printf("%#v\n", grades) // => b is of type float64 and grades is of type []int

	fmt.Printf("b is of type %T and grades is of type %T\n", b, grades)
	// => b is of type float64 and grades is of type []int

	fmt.Printf("The address of a: %p\n", &a) // => The address of a: 0xc000016128

	fmt.Printf("%c and %c\n", 100, 51011) // =>  d and 읃  (runes for code points 101 and 51011)

	const pi float64 = 3.14159265359

	fmt.Printf("pi is %.4f\n", pi) // => formatting with 4 decimal points

	// %b -> base 2
	// %x -> base 16
	fmt.Printf("255 in base 2 is %b\n", 255) //  => 255 in base 2 is 11111111

	fmt.Printf("101 in base 16 is %x\n", 101) // => 101 in base 16 is 65

	// fmt.Sprintf() returns a string. Uses the same verbs as fmt.Printf()
	s := fmt.Sprintf("a is %d, b is %f, c is %s \n", a, b, c)

	fmt.Println(s) // => a is 10, b is 15.500000, c is Gophers
}

// ///////////////////////////////
// CONSTANTS
// ///////////////////////////////
func constants() {
	// Constants need to be initialized when declared
	// it is "OK" if we do not use it after it's decleration
	const days int = 7 // typed constant

	const pi = 3.14 // untyped constant

	// There are ONLY boolean constants, rune constants, integer constants,
	// floating-point constants, complex constants, and string constants.

	// Declaring multiple (grouped) constants
	const (
		a         = 5   // untyped constant
		b float64 = 0.1 // typed constant
	)

	const n, m int = 4, 5

	const (
		min1 = -500
		max1 //gets its type and value form the previous constant. It's 500
		max2 //in a grouped constants, a constant repeats the previous one -> 500
	)

	// CONSTANTS RULES

	// 1. You cannot change a constant
	const temp int = 100
	// temp = 50 //compile-time error

	// 2. You cannot initiate a constant at runtime (constants belong to compile-time)
	// const power = math.Pow(2, 3) //error, functions calls belong to runtime

	// 3. You cannot use a variable to initialize a constant
	t := 5
	// error, variables belong to runtime and you cannot initialize a const to runtime values
	// const tc = t

	// 4. You can use a function like len() to initialize a const if it has as argument
	// a constant string literal (not a variable)
	// a string literal is an untyped constant

	const l1 = len("Hello") //ok

	str := "Hello"
	// const l2 = len(str) //error, str is a variable and belongs to runtime

	_, _ = t, str

	// UNTYPED CONSTANTS
	const x = 5

	const y float64 = 1.1

	const z float64 = x // --> const z float64 = float64(x)

	var v1 int = 5

	var v2 float64 = 1.1

	fmt.Println(x * y)
	// => 5.5, No Error because x is untyped and gets its type when its used first time (float64).

	// fmt.Println(v1 * v2)
	// => Error: invalid operation: v1 * v2 (mismatched types int and float64)
	_, _ = v1, v2
}

// ///////////////////////////////
// IOTA
// ///////////////////////////////
func iotaPractical() {
	// iota is number generator for constants which starts from zero
	// and is incremented by 1 automatically.

	const (
		c1 = iota
		c2 = iota
		c3 = iota
	)
	fmt.Println(c1, c2, c3) // => 0 1 2

	const (
		North = iota //by default 0
		East         //omitting type and value means, repeating its type and value so East = iota = 1 (it increments by 1 automatically)
		South        // -> 2
		West         // -> 3
	)

	// Initializing the constants using a step:
	const (
		c11 = iota * 2 // -> 0
		c22            // -> 2
		c33            // -> 4
	)
}

// ///////////////////////////////
// Go Data Types
// ///////////////////////////////
func dataType() {
	//** NUMERIC TYPES **//

	// uint8      the set of all unsigned  8-bit integers (0 to 255)
	// uint16      the set of all unsigned 16-bit integers (0 to 65535)
	// uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
	// uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

	// int8        the set of all signed  8-bit integers (-128 to 127)
	// int16       the set of all signed 16-bit integers (-32768 to 32767)
	// int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
	// int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

	// uint     either 32 or 64 bits (based on your system)
	// int      same size as uint

	// float32     the set of all IEEE-754 32-bit floating-point numbers
	// float64     the set of all IEEE-754 64-bit floating-point numbers
	// complex64   the set of all complex numbers with float32 real and imaginary parts
	// complex128  the set of all complex numbers with float64 real and imaginary parts

	// byte        alias for uint8
	// rune        alias for int32 (used to store character)

	//int type
	var i1 int8 = -128 //min value

	fmt.Printf("%T\n", i1) // => int8

	var i2 uint16 = 65535 //max value

	fmt.Printf("%T\n", i2) // => int16

	var i3 int64 = -324_567_345 // underscores are used to write large numbers for a better readability

	fmt.Printf("%T\n", i3) // => int64

	fmt.Printf("i3 is %d\n", i3) // => i3 is -324567345 (underscores are ignored)

	//float64 type
	var f1, f2, f3 float64 = 1.1, -.2, 5. // trailing and leading zeros can be ignored
	//1.1, -0.2, 5.0
	fmt.Printf("%T %T %T\n", f1, f2, f3)

	//rune type
	var r rune = 'f'

	fmt.Printf("%T\n", r) // => int32 (rune is an alias to int32)

	fmt.Printf("%x\n", r) // => 66,  the hexadecimal ascii code for 'f'

	fmt.Printf("%c\n", r) // => f

	//bool type
	var b bool = true

	fmt.Printf("%T\n", b) // => bool

	//string type
	var s string = "Hello Go!"

	fmt.Printf("%T\n", s) // => string

	//array type
	var numbers = [4]int{4, 5, -9, 100}

	fmt.Printf("%T\n", numbers) // =>  [4]int

	//slice type
	var cities = []string{"London", "Bucharest", "Tokyo", "New York"}

	fmt.Printf("%T\n", cities) // => []string

	//map type
	balances := map[string]float64{
		"USD": 233.11,
		"EUR": 555.11,
	}

	fmt.Printf("%T\n", balances) // => map[string]float64

	//struct type
	type Person struct {
		name string
		age  int
	}
	var you Person

	fmt.Printf("%T\n", you) // => main.Person

	//pointer type
	var x int = 2

	ptr := &x // pointer to int

	fmt.Printf("ptr is of type %T with value %v\n", ptr, ptr) // => ptr is of type *int with value 0xc000016168

	//function type
	fmt.Printf("%T\n", declaration) // => func()
}

// ///////////////////////////////
// CONVERTING NUMBER TO STRING AND STRING TO NUMBER
// ///////////////////////////////
func numConvStr() {

	var x = 3 //int type

	var y = 3.2 //float type

	// x = x * y //compile error ->  mismatched types

	x = x * int(y) // converting float64 to int

	fmt.Println(x) // => 9

	y = float64(x) * y //converting int to float64

	fmt.Println(y) // => 28.8

	x = int(float64(x) * y)

	fmt.Println(x) // => 259

	//In Go types with different names are different types.

	var a int = 5 // same size as int64 or int32 (platform specific)

	var b int64 = 2 // int and int64 are not the same type

	// a = b // error: cannot use b (type int64) as type int in assignment

	a = int(b) // converting int64 to int (explicit conversion required)

	// preventing unused variable error
	_ = a

	//** CONVERTING NUMBERS TO STRINGS AND STRINGS TO NUMBERS **//

	s := string(99) // int to rune (Unicode code point)

	fmt.Println(s) // => 99, the ascii code for symbol c

	fmt.Println(string(34234)) // => 34234 is the unicode code point for 薺

	// we cannot convert a float to a string similar to an int to a string
	// s1 := string(65.1) // error

	// converting float to string
	var myStr = fmt.Sprintf("%f", 5.12)

	fmt.Println(myStr) // => 5.120000

	// converting int to string
	var myStr1 = fmt.Sprintf("%d", 34234)

	fmt.Println(myStr1) // => 34234

	// converting string to float
	var result, err = strconv.ParseFloat("3.142", 64)

	if err == nil {
		fmt.Printf("Type: %T, Value: %v\n", result, result) // => Type: float64, Value: 3.142
	} else {
		fmt.Println("Cannot convert to float64!")
	}

	// Atoi(string to int) and Itoa(int to string).
	i, err := strconv.Atoi("-50")

	s = strconv.Itoa(20)

	fmt.Printf("i Type is %T, i value is %v\n", i, i) // => i Type is int, i value is -50

	fmt.Printf("s Type is %T, s value is %q\n", s, s) // => s Type is string, s value is "20"
}

// ///////////////////////////////
// TYPES
// ///////////////////////////////
func goType() {
	// new type speed (underlying type uint)
	type speed uint

	// s1, s2 of type speed
	var s1 speed = 10

	var s2 speed = 20

	// performing operations with the new types
	fmt.Println(s2 - s1) // -> 10

	// uint and speed are different types (they have different names)
	var x uint

	// x = s1  //error different types

	// correct
	x = uint(s1)

	_ = x

	// correct
	var s3 speed = speed(x)

	fmt.Printf("%T\n", s3)
}

// ///////////////////////////////
// Command Line Arguments
// ///////////////////////////////
func commandLineArguments() {
	fmt.Println("os.Args:", os.Args) // os.Args is slice of strings ([]string)

	// accessing command line arguments using indexes
	fmt.Println("Path:", os.Args[0])

	//fmt.Println("1st Argument:", os.Args[1])
	//
	//fmt.Println("2nd Argument:", os.Args[2])

	fmt.Println("No. of items inside os.Args:", len(os.Args))
}
