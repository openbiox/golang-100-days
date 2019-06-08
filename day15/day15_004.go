package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	seconds := []int{1, 2, 3, 4}
	for i, s := range seconds {
		// counter plus 1
		wg.Add(1)
		go func(i, s int) {
			// counter minus 1
			defer wg.Done()
			fmt.Printf("goroutine%d ending\n", i)
		}(i, s)
	}

	// waiting
	wg.Wait()
	fmt.Println("all goroutine finished")
}
