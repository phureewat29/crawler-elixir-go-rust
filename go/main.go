package main

import (
	"fmt"
	"sync"
)

func processMessages(messages []string, wg *sync.WaitGroup) {
	for _, msg := range messages {
		fmt.Printf("Processing message: %s\n", msg)
	}

	wg.Done()
}

func main() {
	messages := []string{"Message 1", "Message 2", "Message 3", "Message 4", "Message 5"}

	var wg sync.WaitGroup
	wg.Add(len(messages))

	for _, msg := range messages {
		go func(m string) {
			processMessages([]string{m}, &wg)
		}(msg)
	}

	wg.Wait()

	fmt.Println("All messages processed!")
}
