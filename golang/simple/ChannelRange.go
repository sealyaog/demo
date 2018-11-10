package main

import(
	"fmt"
	"sync"
)


func main() {
	toProcess := make(chan int , 1000)
 	for i:= 0; i < 1000; i++ {
		toProcess <- i
	}
	close(toProcess)

	wg := sync.WaitGroup{}
	wg.Add(100)

	var mux sync.Mutex;

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			for piece := range toProcess {
				mux.Lock()
				fmt.Println(piece);	
				mux.Unlock()
			}
		}()
	}
	wg.Wait()
}
