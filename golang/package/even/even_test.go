package even

import "testing"

//cmd: go test
func TestEven(t *testing.T) {
    if true != Even(2) {
        t.Log("2 should be even!")
        t.Fail()
    }
}

//cmd: test -bench=. even
func BenchmarkTestEven(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Even(i)
    }
}
