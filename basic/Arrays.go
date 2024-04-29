package main

import (
	"fmt"
	"math"
	"slices"
	"strings"
)

// IMP
// 1. cannot assign value nil in array
func main() {

	slice()

	//arrays()

}

func arrays() {
	var arr = [5]int{1, 2, 3, 4, 5} // any data-type on right side will work

	var arr2 [5]int = [5]int{1, 2, 3, 4, 5} // data-type that is left side should be on right side

	fmt.Printf("%T, %v\n", arr, arr2)

	fmt.Printf("%T, %#v\n", arr2, arr2) // %#v --> also print type of variable along with data

	arr3 := [...]float64{1, 2, 3, 4.345, 5} // declaring array without knowing it size

	fmt.Printf("%T, %v\n", arr3, arr3)

	arr4 := [5]string{"x", "y", "z"}

	fmt.Printf("%T, %v\n", arr4, arr4) // rest will start with zero value

	fmt.Println("length of arr4: ", len(arr4))

	// iterating over an array (2-ways)
	for i, v := range arr4 { // range is a language keyword used for iteration
		fmt.Println("index:", i, "value: ", v)

	}

	fmt.Println(strings.Repeat("--", 20))

	// iterating over an array (C/C++, Java Style)
	for i := 0; i < len(arr4); i++ {
		fmt.Println("index:", i, "value: ", arr4[i])
	}

	fmt.Println(strings.Repeat("--", 20))

	// declaring a multi-dimensional arrays (array of arrays or matrix)
	balances := [2][3]int{
		[3]int{5, 6, 7}, //[3]int is optional
		{8, 9, 10},      //if we do not write this line then it will by default make array of size 3  with default value
		// {8,9}		 //if array size != 3 then it will automatically add extra element with zero value
	}

	for _, arr := range balances {
		for _, value := range arr {
			fmt.Printf("%d ", value)
		}
		fmt.Println("")
	}

	fmt.Println(strings.Repeat("--", 20))

	//  = operator creates a copy of an array by value
	// the arrays are not connected and are saved in different memory locations
	m := [3]int{1, 2, 3}
	n := m //n is a copy of m (by value)

	fmt.Println("n is equal to m: ", n == m) // => true

	m[0] = -1 //only m is modified

	fmt.Println("n is equal to m: ", n == m) // => false

	grades := [3]int{ //the keyed elements can be in any order
		1: 10,
		0: 5,
		2: 7,
	}
	fmt.Println(grades) // -> [5 10 7]

	accounts := [3]int{
		2: 50,
	}
	fmt.Println(accounts) //[0 0 50]

	names := [...]string{
		4: "Dan",
	}
	fmt.Println(len(names)) // -> 5

	// un un-keyed element gets its index from the last keyed element
	cities := [...]string{
		5:        "Paris",
		"London", // this is at index 6, takes next index of its above defined index
		1:        "NYC",
	}
	fmt.Printf("%#v\n", cities) // -> [7]string{"", "NYC", "", "", "", "Paris", "London"}

	weekend := [7]bool{5: true, 6: true} // use case of indexed array declaration
	fmt.Println(weekend)                 // => [false false false false false true true]

	////////////////////////////////////////
	// EXERCISE
	/*package main

	import "fmt"

	func main() {
		myArray := [3]float64{1.2, 5.6}

		myArray[2] = 6

		a := 10
		myArray[0] = a

		myArray[3] = 10.10

		fmt.Println(myArray)

	}
	*/
	////////////////////////////////////////

	myArray := [3]float64{1.2, 5.6}

	myArray[2] = 6

	fmt.Println(myArray)

}

func slice() {
	s1 := []int{1, 2, 3} //method-1
	_ = s1

	s := make([]string, 3, 6) // method-2 (here 6 is the capacity of slice)

	fmt.Printf("%#v %v\n", s, cap(s))

	s = append(s, "hello")

	fmt.Println(s)

	//////////////////////////////////////////
	//advance slice

	numbers := []int{2, 3}

	// append() returns a new slice after appending a value to its end
	numbers = append(numbers, 10)
	fmt.Println(numbers) //-> [2 3 10]

	// appending more elements at once
	numbers = append(numbers, 20, 30, 40)
	fmt.Println(numbers) //-> [2 3 10 20 30 40]

	// appending all elements of a slice to another slice
	n := []int{100, 200, 300}
	numbers = append(numbers[:5], n...) // ... is the ellipsis operator  // this will overwrite value 40 with 100
	fmt.Println(numbers)                // -> [2 3 10 20 30 100 200 300]

	//** Slice's Length and Capacity **//

	nums := []int{1}
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums)) // Length: 1, Capacity: 1

	nums = append(nums, 2)
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums)) // Length: 2, Capacity: 2

	nums = append(nums, 3)
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums)) // Length: 3, Capacity: 4
	// the capacity of the new backing array is now larger than the length
	// to avoid creating a new backing array when the next append() is called.

	nums = append(nums, 4, 5)
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums)) // Length: 5, Capacity: 8

	// copy() function copies elements into a destination slice from a source slice and returns the number of elements copied.
	// if the slices don't have the same no of elements, it copies the minimum of length of the slices
	src := []int{10, 20, 30}
	dst := make([]int, len(src))
	nn := copy(dst, src)
	dst[1] = 50

	fmt.Printf("%p, --------- %v --------- %#v\n", src, cap(src), src)
	src1 := slices.Delete(src, 0, 2) // delete
	fmt.Printf("%p, --------- %v --------- %#v\n", src1, cap(src1), src1)
	fmt.Println(src, dst, nn) // => [10 20 30] [10 50 30] 3
	src[0] = 40
	fmt.Printf("%p, --------- %v --------- %#v\n", src, cap(src), src)
	fmt.Printf("%p\n", &src1[0])
	fmt.Printf("%p\n", &src[0])
	fmt.Printf("%p, --------- %v --------- %#v\n", &src1, cap(src1), *(&src1))
	test := []int{}
	fmt.Println(test == nil) // false

	newdst := slices.Clone(dst)

	newdst[0] = 0

	fmt.Println(dst)

	var test2 []int
	fmt.Println(test2 == nil) // true

	// cannot use nil as int value in array or slice literal
	//nullValue := [2]int{1, nil}
	//fmt.Println(nullValue)

	//test3 := make([]int, 2147483647+100000000)
	//fmt.Println(test3)

	test3 := [math.MaxInt32 + 1000000000]int32{}
	_ = test3

}
