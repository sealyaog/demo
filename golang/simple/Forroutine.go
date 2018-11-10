package main
import "fmt"
import "sync"

func main() {
	var wg sync.WaitGroup;
	for index := 0 ; index < 10 ; index ++ {
		wg.Add(1)
		go func(int) {
			defer wg.Done()
			fmt.Printf("index: %d\n",index);
		}(index)
	}
	wg.Wait()
	return;
}
