package main 

import (
	"fmt"
	"math"
)

func main() {
	SieveOfEratosthenes(5)
}
func SieveOfEratosthenes(value int) {
    f := make([]bool, value)
    // fmt.Println(f)
    for i := 2; i <= int(math.Sqrt(float64(value))); i++ {
        if f[i] == false {
            for j := i * i; j < value; j += i {
                f[j] = true
            }
        }
    }
    for i := 2; i < value; i++ {
        if f[i] == false {
            fmt.Printf("%v ", i)
        }
    }
    fmt.Println("")
}


func IsPrime(value int) bool {
	
	for i:=2; i<= int(math.Floor(float64(value)/2)); i++ {
		if value%i==0{
			return false
		}
	}
	return value > 1
}

