package main  
  
import (  
    "fmt"  
)  

var s = "world"

type Callback func()

func test( f Callback) {
	f();
}

  
func main() {  
    var s = "hello";  
  
    f := func() { 
		fmt.Println("var:", s)
    }  
	test(f);
	return ;
} 
