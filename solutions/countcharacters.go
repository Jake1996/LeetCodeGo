package solutions

import "fmt"

func countCharacters(words []string, chars string) int {
	arr := make([]int, 26)
	bs := []byte(chars)
	for _, b := range bs {
		arr[b-'a']++
	}
	sum := 0
	for _, word := range words {
		if canForm(word, arr) {
			sum += len(word)
		}
	}
	return sum
}

func canForm(word string, arr []int) bool {
	newArr := make([]int, 26)
	ba := []byte(word)
	for _, b := range ba {
		newArr[b-'a']++
	}
	for i := range 26 {
		if newArr[i] > arr[i] {
			return false
		}
	}
	return true
}

func CountCharacters() {
	fmt.Print(countCharacters([]string{}, ""))
}
