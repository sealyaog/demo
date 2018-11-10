package main
import "fmt"

var gi = 1  //first run here

func init() { // second run here
    fmt.Printf("in init gi:%v\n",gi);
    gi = 2  
}


func main() { //last here
    fmt.Printf("in main gi:%v\n",gi);
}
