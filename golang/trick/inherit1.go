package main 
import "fmt"

type Base struct {
	Name string
	Age int
}

type Child struct {
	Base 
	//Age int
}

func main() {
	c := new(Child)
	c.Name = "hello" 
	c.Age = 20 
	//c.Base.Age = 30


	fmt.Println(c.Name)  // hello
	fmt.Println(c.Age)  // 20
	//fmt.Println(c.Base.Age) 
}
