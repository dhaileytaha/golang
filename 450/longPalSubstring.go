package main

type Pointer struct {
	left  int
	right int
}

func (p *Pointer) sizeOfPal() int {
	return p.right - p.left
}

func LongestPalindromicSubstring(str string) string {
	// Write your code here.

	currentPal := Pointer{0, 1}
	for i := 1; i < len(str); i++ {
		odd := getLongestPal(str, i-1, i+1)
		even := getLongestPal(str, i-1, i)
		longest := even
		if odd.sizeOfPal() > even.sizeOfPal() {
			currentPal = odd
		}
		if longest.sizeOfPal() > currentPal.sizeOfPal() {
			currentPal = longest
		}
	}
	return str[currentPal.left:currentPal.right]
}

func getLongestPal(str string, left int, right int) Pointer {
	for left >= 0 && right < len(str) {
		if str[left] != str[right] {
			break
		}
		left++
		right--
	}
	return Pointer{left + 1, right}
}
