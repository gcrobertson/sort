package main

import "fmt"

func main1() {
	c := make(chan int)
	go func() { // another thread
		// this code blocks in the go-routine
		c <- 42
	}()
	// the main routine blocks until it can take the value off
	fmt.Println(<-c)
}

func main2() {
	// buffered channel. channel can store this many before blocking.
	c := make(chan int, 1) // type: "chan int"
	c <- 42
	fmt.Println(<-c)
}

func main3() {
	// directional channels
	c := make(chan<- int, 2) // send channel
	d := make(<-chan int, 2) // receive channel
	fmt.Printf("%T %T", c, d)
}

func main4() {
	// general to specific channel is ok, specific to general is not
	c := make(chan int)
	cr := make(<-chan int)
	fmt.Printf("%T %T", cr, (chan<- int)(c))
}

func main5() {
	c := make(chan int)
	go sender(c)
	receiver(c) // receiver blocks. if go routine here then good chance `main` will exit before send/receive.
	fmt.Println("exiting main...")
}

// send
func sender(c chan<- int) {
	c <- 42
}

// receive
func receiver(c <-chan int) {
	fmt.Println(<-c)
}

func main6() {
	c := make(chan int)

	// without `go` there is deadlock because all goroutines are asleep
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c) // we must close c or the receiver will never know when to stop ranging.
	}()

	// receive
	for v := range c { // will range over a channel until it is closed.
		fmt.Println(v)
	}

	fmt.Println("exiting main...")
}

// when you are relaxed, the brain works better.
// it's not a big deal. remember what go is about: it's about "ease" of programming.
// i want to code idiomatic go.
// - channels block
// - buffered channels hold so many values until they also block
// the range loop to receive will continue to hang out until the channel is closed
func main7() {

	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// send
	go send(even, odd, quit)

	// receive
	receive(even, odd, quit)
}

func send(e, o, q chan<- int) {

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	// close(e) these add more 0s to the channel, not sure why...
	// close(o)
	q <- 0
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("from even channel: ", v)
		case v := <-o:
			fmt.Println("from odd channel: ", v)
		case _ = <-q:
			fmt.Println("quitting...")
			// close(q) // cannot close a receive only channel
			return
		}
	}
}

func main() {

	// comma, ok idiom for channels
	// https://golang.org/doc/effective_go.html

	c := make(chan int)

	go func() {
		c <- 42
	}()

	v, ok := <-c

	fmt.Println(v, ok) // 42, ok

	close(c)

	v, ok = <-c

	fmt.Println(v, ok) // 0, false

	// 22
	// 23

	// comma ok idiom for channels:
	// - when the channel is open, ok is true.
	// - when the channel is closed, ok is false.

	// https://www.udemy.com/course/learn-how-to-code/learn/lecture/11922356#overview

	// fan in pattern
}
