package main
import "fmt"
func main() {
    m := make(map[string]int)
    m["k1"] = 7
    m["k2"] = 13
    fmt.Println("map:", m)
    v1 := m["k1"]
    fmt.Println("v1: ", v1)
    fmt.Println("len:", len(m))

    delete(m, "k2")
    fmt.Println("map:", m)

    _, prs := m["k2"] 
    fmt.Println("prs:", prs) // false, m does not have key "k2"

    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)

	if v,ok := m["k3"] ; ok{
		fmt.Println(v);
	} else {
		fmt.Println("Key Not Found")
	}


	l := map[string]int{"k1": 1, "k2": 2, "k3": 3, "k4": 4, "k5": 5, "k6": 6}

	var randomKey string 
	for randomKey = range l {
		fmt.Println("rondomKey", randomKey);	
		break
	} 

    for randomKey = range l {
	    fmt.Println("rondomKey", randomKey);    
	    break
	}
}
