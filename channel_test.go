package golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Sigit Priadi"
		fmt.Println("Selesai mengirim data ke Channel")
	}()

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)

	// fmt.Println(<-channel)
}
