package main 

import "fmt"

func main(){
	generate(1,10)
}
func generate(min int, max int){
	a := make([]int, max+min-1)

	for i, v := range a {
		v =  min+i
		fmt.Println(i, v)
	}
}