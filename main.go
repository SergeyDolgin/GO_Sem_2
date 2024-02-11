package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	fmt.Println("Введите предложение")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Println("Исходное предложение:", text)
	counts := make(map[rune]int)
	var totalChar int
	for _, char := range text {
		char = unicode.ToLower(char)
		if unicode.IsLetter(char) {
			counts[char]++
			totalChar++
		}
	}
	fmt.Println("Общее количество символов в исходном предложении", totalChar)
	for char, count := range counts {
		var percentage float32
		// c := float32(count)
		if count > 0 {
			percentage = float32(count) / float32(totalChar) * 100
			fmt.Printf("Символ %c появляется %d  раз, в процентном соотношении %0.1f\n", char, count, percentage)
		}
	}
}
