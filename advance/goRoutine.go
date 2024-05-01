package advance

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

func Start() {
	//basic() // this will not be executed properly because of function end before go routine can be completed
	//
	//fmt.Println(strings.Repeat("##", 20))
	//
	//basicWithWaitGroup()
	//
	//fmt.Println(strings.Repeat("##", 20))
	//
	//channels()
	//
	//fmt.Println(strings.Repeat("##", 20))
	//
	//timeout()
	//
	//fmt.Println(strings.Repeat("##", 20))
	//
	//channelClosing()
	//
	//fmt.Println(strings.Repeat("##", 20))
	//
	//timerExample()
	//
	//fmt.Println(strings.Repeat("##", 20))
	//
	//ticker()
	//
	//fmt.Println(strings.Repeat("##", 20))
	//
	//checkConcurrent()
	fmt.Println(strings.Repeat("##", 20))

	go func() {

		select {}

		fmt.Println("hello world")
	}()

	for {

		time.Sleep(time.Second * 100)
	}

	//select {} --> fatal error: all goroutines are asleep - deadlock!

	//var str string = "hello world"
	//fmt.Println(str)
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

func channels() {
	ch := make(chan string, 2)

	defer close(ch)

	defer func() { fmt.Println("closing channel") }()

	wg := &sync.WaitGroup{}

	wg.Add(1)

	go func(c chan<- string, wg *sync.WaitGroup) {
		for i := 0; i < 10; i++ {
			c <- fmt.Sprintf("hello world, message %d", i)
			time.Sleep(2 * time.Second)
		}
		wg.Done()
	}(ch, wg)

	go func(c <-chan string) {
		for v := range c { // equivalent v:= <- c
			fmt.Println(v, " go routine 1")
		}
	}(ch)

	go func(c <-chan string) {
		for v := range c {
			fmt.Println(v, " go routine 2")
		}
	}(ch)

	wg.Wait()

	fmt.Println("program complete")

	fmt.Println(strings.Repeat("--", 20))

	queue := make(chan string, 2)

	queue <- "one"

	queue <- "two"

	close(queue)

	for elem := range queue { // it is possible to read from channel even after closing
		// if there is value in channel buffer
		fmt.Println(elem)
	}

}

func timeout() {

	c1 := make(chan string, 1)

	go func() {

		time.Sleep(2 * time.Second)

		c1 <- "result 1"

	}()

	select { // this will print timeout

	case res := <-c1:

		fmt.Println(res)

	case <-time.After(1 * time.Second):

		fmt.Println("timeout 1")

	}

	// If we allow a longer timeout of 3s, then the receive
	// from `c2` will succeed and we'll print the result.

	c2 := make(chan string, 1)

	go func() {

		time.Sleep(2 * time.Second)

		c2 <- "result 2"

	}()

	select { // this will print result 2

	case res := <-c2:

		fmt.Println(res)

	case <-time.After(3 * time.Second):

		fmt.Println("timeout 2")

	}
}

func channelClosing() {
	jobs := make(chan int, 5)

	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done // program will wait until all  job has been received

	_, ok := <-jobs
	fmt.Println("received more jobs:", ok) // received more jobs: false

}

func timerExample() {

	timer1 := time.NewTimer(2 * time.Second) // this will return channel

	<-timer1.C // message received on channel of timer is fired

	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)

	go func() {

		<-timer2.C

		fmt.Println("Timer 2 fired")

	}()

	stop2 := timer2.Stop()

	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(2 * time.Second)

}

func ticker() {

	ticker := time.NewTicker(500 * time.Millisecond) // This will also return channel

	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(2000 * time.Millisecond)

	//ticker.Stop()

	done <- true

	fmt.Println("Ticker stopped")

}

type test struct {
	value int
}

func checkConcurrent() {
	var wg sync.WaitGroup
	wg.Add(1)
	map1 := map[string]int{
		"key": 1,
	}
	_ = map1

	a := 0

	var st test
	st.value = 0

	go func() {
		for i := 0; i < 50000; i++ {
			//map1["key"]++ --> concurrent modification error
			a++
			st.value++
		}
		wg.Done()
	}()

	for i := 0; i < 50000; i++ {
		//map1["key"]--		--> concurrent modification error
		a--
		st.value--
	}
	wg.Wait()
	fmt.Println(a)
}
