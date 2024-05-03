package main

import (
	"Golang/advance"
)

func main() {

	//print("hello world")
	//advance.Structs()

	//advance.Method()

	//advance.MethodPointer()

	//advance.Error()

	advance.Start()

	//mapper := make(map[int]struct{})
	//prepare(mapper)
	//fmt.Println(fmt.Sprintf("value are %v", mapper))
	//fmt.Println(mapper())
}

//func mapper() int {
//	defer func() int {
//		if err := recover(); err != nil {
//			fmt.Println(err)
//			return 3
//		}
//		return 0
//
//	}()
//}

//func prepare(mapper map[int]struct{}) {
//	for i := 0; i <= 5; i++ {
//		defer func() { mapper[i] = struct{}{} }()
//	}
//}
