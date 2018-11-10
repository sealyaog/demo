package main
import "fmt"

type Callables interface{
	Call() interface{};
}

type C1 struct {
}

func (c C1) Call() interface{} {
	fmt.Println("c1");
	return 0.5
}

type C2 struct {
}

func (c C2) Call() interface{} {
	fmt.Println("c2");
	return 1
}

func UntilFuncRet (callables []Callables) interface{} {
	//c := make(chan interface{}); //some goroutine will be hanged here
	c := make(chan interface{}, len(callables));
	for index := range callables {
		idx := index
		go func (i int) {
			fmt.Printf("before, %d\n",i)
			c <- callables[i].Call()
			fmt.Printf("after, %d\n",i)
		}(idx)
	}
	return <- c
}

func main() {
	var c1 C1;
	var c2 C2;
	calls := [...]Callables{c1,c2} //to be opted
	scalls := calls[0:2] 
	fmt.Println(UntilFuncRet(scalls));
}
