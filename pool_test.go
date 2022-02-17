package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	// var pool sync.Pool
	pool := sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}

	pool.Put("Sigit")
	pool.Put("Priadi")
	pool.Put("Joko")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	time.Sleep(11 * time.Second)
	fmt.Println("All goroutines finished")
}
