package main
import "fmt"
import "sync"

func main() {
	c := make (chan int , 100)
	for i:= 0; i < 100 ; i ++ {
		c <- i
	}
	close(c) //yaog: close here first

	var wg sync.WaitGroup;
	for i:=0; i < 10 ; i++ {
		wg.Add(1)
		go func () {
			defer wg.Done()
			for k:= range c {
				fmt.Printf("number: %d\n", k)
			}
		}()
	}
	wg.Wait()
	return 

}
