package main 

import (
    //"./even" use the even in the same dir
    "even"
    "fmt"
)

func main() {
    i := 5 
    fmt.Printf("Is %d even? %v\n", i , even.Even(i))
}
