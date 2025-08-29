package solutions

import "fmt"

func wrapLines(words []string, length int) []string {
	var returnArray []string
	initialIndex := 0
	i := 0
	currentWord := ""
	for i < len(words) {
		word := words[i]
		if len(currentWord) == 0 {
			currentWord = word
			initialIndex += len(word)
		} else if len(currentWord) > 0 && initialIndex+len(word) < length {
			currentWord = fmt.Sprintf("%v-%v", currentWord, word)
			initialIndex += len(word) + 1
		} else {
			if len(currentWord) > 0 {
				returnArray = append(returnArray, currentWord)
			}
			initialIndex = 0
			currentWord = word
			initialIndex += len(word)
		}
		i++
	}
	if len(currentWord) > 0 {
		returnArray = append(returnArray, currentWord)
	}
	return returnArray
}

func WrapLines() {
	var words1 []string = []string{"The", "day", "began", "as", "still", "as", "the", "night", "abruptly", "lighted", "with", "brilliant", "flame"}
	fmt.Print(wrapLines(words1, 13))

}
