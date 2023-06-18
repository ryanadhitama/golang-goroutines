package golang_goroutines

import (
	"fmt"
	"testing"
	"time"
	"sync"
)

func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex
	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x = x + 1
				fmt.Println(x)
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Counter =", x)
}