package golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

// goroutines
func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Ups")
	time.Sleep(1 * time.Second)
}

// goroutines with looping
func DisplayNumber(number int) {
	fmt.Println("Number", number)
}

func TestCreateGoroutineLoop(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}
	
	time.Sleep(1 * time.Second)
}

// channel
func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Ryan Adhitama Putra"
		fmt.Println("Sending data to channel done")
	}()

	data := <- channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

// Channel as Parameter
func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Ryan Adhitama"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <- channel
	fmt.Println(data)
}

// Channel in and out
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Ryan Adhitama Putra"
}

func OnlyOut(channel <-chan string) {
	data := <- channel
	fmt.Println(data)
}

func TestChannelInOut(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}