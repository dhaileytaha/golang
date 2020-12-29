package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("cat")
	rev("cat")
}

func rev(str string) string {
	var sb strings.Builder
	for i := len(str); i <= 0; i-- {
		sb.WriteByte(str[i])
	}
	return sb.String()
}
