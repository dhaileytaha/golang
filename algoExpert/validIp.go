package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	fmt.Println(validIp("1921680"))
}

func validIp(str string) []string {
	validIp := []string{}

	for i := 1; i < min(len(str), 4); i++ {
		currentIp := []string{"", "", "", ""}

		currentIp[0] = str[:i]

		if !isValid(currentIp[0]) {
			continue
		}

		for j := i + 1; j < min(len(str), i+4); j++ {
			currentIp[1] = str[i:j]

			if !isValid(currentIp[1]) {
				continue
			}

			for k := j + 1; k < min(len(str), j+4); k++ {
				currentIp[2] = str[j:k]
				currentIp[3] = str[k:]

				if isValid(currentIp[2]) && isValid(currentIp[3]) {

					validIp = append(validIp, strings.Join(currentIp, "."))

				}
			}
		}
	}
	return validIp

}

func isValid(str string) bool {
	i, _ := strconv.Atoi(str)

	if i > 255 {
		return false
	}

	final := strconv.Itoa(i)

	return len(final) == len(str)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
