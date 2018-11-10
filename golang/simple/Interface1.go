package main
import "fmt"

type S struct {i int}

func (p *S) Get() int {return p.i}
func (p *S) Put(v int) {p.i = v}

type R struct {i int}

func (p *R) Get() int {return p.i}
func (p *R) Put(v int) {p.i = v}


type I interface{
    Get() int
    Put(int)
}

func f(p I){
    fmt.Println(p.Get())
    p.Put(1);
}

func mytype(p I) {
    switch p .(type) {
        case *S: 
            fmt.Println("*S")    
            break
        case *R:
            fmt.Println("*R")
            break
        /* can not pass compile here, must be *S or *R
        case S:
            fmt.Println("S")
            break
        case R:
            fmt.Println("R")
            break
        */
        default:
            fmt.Println("default")
    }
}

func RunTime(any interface{}) int{
    return any.(I).Get() //yaog: ".(I)" is type assert
}

func main(){
    var s S;
    f(&s);

    var r R;
    //mytype(s);
    mytype(&s);
    //mytype(r);
    mytype(&r);

    RunTime(r);
}
