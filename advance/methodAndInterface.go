package advance

import (
	"fmt"
	"strings"
	"time"
)

// declaring a new type
type names []string

// declaring a method (function receiver)
func (n names) print() {
	// n is called method's receiver
	// n is the actual copy of the names we're working with and is available in the function.
	// n is like this or self from OOP.
	// any variable of type names can call this function on itself like variable_name.print()

	// iterating and printing all names
	for i, name := range n {
		fmt.Println(i, name)
	}
}

func Method() {

	// Go doesn't have classes, but you can define methods on defined types.
	// a type may have a method set associated with it which enhances the type with extra behaviour.

	const day = 24 * time.Hour
	fmt.Printf("%T\n", day)

	// calling a method on time.Duration type
	// Seconds() is a method aka a receiver function.
	seconds := day.Seconds()

	// Seconds() returns the duration as a floating point number of seconds.
	fmt.Printf("%T\n", seconds) //its type is float64

	fmt.Println("Seconds in a day:", seconds) // Seconds in a day: 86400

	fmt.Println(strings.Repeat("-", 20))

	// declaring a value of type names
	friends := names{"Dan", "Marry"}
	// calling the print() method on friends variable
	friends.print()

	fmt.Println(strings.Repeat("-", 20))

	// method-2
	names.print(friends) // not recommended

	fmt.Println(strings.Repeat("-", 20))

}

//////////////////////////////////
// METHOD USING POINTER
/////////////////////////////////

// declaring a new struct type
type car struct {
	brand string
	price int
}

// declaring a method for car type
// it changes the value it works on
func (c car) changeCar1(newBrand string, newPrice int) {
	c.brand = newBrand
	c.price = newPrice
	// the changes are not seen to the outside world (pass by value)
}

// declaring a method with a pointer receiver
func (c *car) changeCar2(newBrand string, newPrice int) {
	(*c).brand = newBrand
	(*c).price = newPrice
	// the changes are seen the outside world

}

// method declarations are not permitted on named types that are themselves pointer types
type distance *int

// ERROR ->  invalid receiver type *distance (distance is a pointer type)
// func (d *distance) f() {
//  fmt.Println("say something")
// }

func MethodPointer() {

	// declaring a car value
	myCar := car{brand: "Audi", price: 40000}

	// calling the method with a value receiver
	myCar.changeCar1("Opel", 21000)

	// no change due to the same pass by value mechanism  !!!
	fmt.Println(myCar) // => {Audi 40000}

	fmt.Println(strings.Repeat("-", 20))

	// calling the method with a pointer receiver. It changes the value!
	(&myCar).changeCar2("Seat", 25000)
	fmt.Println(myCar) // -> {Seat 25000}

	// declaring a pointer to car
	var yourCar *car
	yourCar = &myCar      // it stores the address of myCar
	fmt.Println(*yourCar) // -> {Seat 25000}

	fmt.Println(strings.Repeat("-", 20))

	// calling the method on pointer type
	// valid ways to call the method:
	yourCar.changeCar2("VW", 30000)
	(*yourCar).changeCar2("VW", 30000)
	fmt.Println(*yourCar) // => {VW 30000}

	// in the above example both myCar and yourCar variables have been modified.
	fmt.Println(myCar) // => {VW 30000}

	fmt.Println(strings.Repeat("-", 20))
}
