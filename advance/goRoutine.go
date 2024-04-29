package advance

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

func Start() {
	//basic() // this will not be executed properly because of function end before go routine can be completed

	fmt.Println(strings.Repeat("--", 20))

	basicWithWaitGroup()
}

func basic() {
	fmt.Println("before starting of go routine")

	go func() {
		fmt.Println("starting of go routine")
		time.Sleep(2 * time.Second)
		fmt.Println("ending of go routine")

	}()

	fmt.Println("after starting of go routine")

	fmt.Println("ending of basic function")
}

func basicWithWaitGroup() {
	var wg sync.WaitGroup

	wg.Add(1)

	fmt.Println("before starting of go routine")

	go func(w *sync.WaitGroup) {
		fmt.Println("starting of go routine")
		time.Sleep(2 * time.Second)
		fmt.Println("ending of go routine")
		w.Done()
	}(&wg)

	wg.Wait()

	fmt.Println("after starting of go routine")

	fmt.Println("ending of basic function")
}
