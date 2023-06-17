package golang_goroutines

import (
	"fmt"
	"testing"
	"time"
	"strconv"
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

// Buffered channel
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 2)
	defer close(channel)

	channel <- "Ryan"
	channel <- "Adhitama"

	fmt.Println(<- channel)
	fmt.Println(<- channel)

	time.Sleep(2 * time.Second)
}

// Range channel
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	} ()


	for data := range channel {
		fmt.Println("Menerima data", data)
	}

	fmt.Println("Selesai")
	
	time.Sleep(2 * time.Second)
}

// Select channel
func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0

	for {
		select {
		case data := <- channel1:
				fmt.Println("Data dari channel 1", data)
				counter++
		case data := <- channel2:
				fmt.Println("Data dari channel 2", data)
				counter++
		}

		if counter == 2 {
			break;
		}
	}

	
	
	time.Sleep(2 * time.Second)
}

// Default Select channel
func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0

	for {
		select {
		case data := <- channel1:
				fmt.Println("Data dari channel 1", data)
				counter++
		case data := <- channel2:
				fmt.Println("Data dari channel 2", data)
				counter++
		default:
				fmt.Println("Menunggu data")
		}

		if counter == 2 {
			break;
		}
	}
}