package main

import "fmt"

func main() {
	var map1 map[int]string   // create map with null
	fmt.Printf("%#v\n", map1) // map[int]string(nil)

	var map2 = make(map[int]string) // method-2 to create map
	fmt.Printf("%#v\n", map2)       // map[int]string{}

	map3 := map[int]string{}  // method-3 to create map
	fmt.Printf("%#v\n", map3) // map[int]string{}

	//map1[1] = "shivam"  --> this will give an error : [panic: assignment to entry in nil map]

	map2[1] = "shivam"
	map2[3] = "yash"
	map2[2] = "raj"

	fmt.Println(map2)

	for key, value := range map2 {
		fmt.Printf("key=%d, value=%s\n", key, value)
		if key == 2 {
			map2[2] = "vaibhav"
		}
	}

	fmt.Println(map2) // map[1:shivam 2:vaibhav 3:yash]

	map4 := map2 // this will point to same memory location

	map4[2] = "yash" // change in map4 also changes map2

	fmt.Println(map2) // map[1:shivam 2:yash 3:yash]

	fmt.Println(map4) // map[1:shivam 2:yash 3:yash]

	map5 := map[string]float64{
		"A": 201.4,
		"B": 201.4,
		//"B": 600.4,  --> will give error : having same key
	}

	fmt.Println(map5)

	fmt.Println(map5["C"]) // 0

	v, ok := map5["C"] // print else statement

	//v is the key's corresponding value
	// ok is bool type value which is true if the key exists and false otherwise

	if ok {
		fmt.Println("The RON Balance is: ", v)
	} else {
		fmt.Println("The RON key doesn't exist in the map!")
	}

	map6 := map[string][]int{}
	_ = map6

	map7 := map[string]map[string]float64{}

	map7["test1"] = map5

	fmt.Println(map7)

}
