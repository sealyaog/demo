package main
import "fmt"
func intSeq() func() int {
    i := 0
    return func() int {
        i += 1
        return i
    }
}
func main() {
    nextInt := intSeq() //the value of i will be remained in this instance
    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())
    newInts := intSeq()
    fmt.Println(newInts())
}
