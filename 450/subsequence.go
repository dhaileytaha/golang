package main

func main() {

}

func calcSubsequence(arr []int, seq []int) {
	arrId := 0
	seqId := 0
	for arrId < len(arr) && seqId < len(seq) {
		if arr[arrId] == seq[seqId] {
			seqId = +1
		} else {
			arrId = +1
		}
	}
	return seqIdx == len(sequence)
}
