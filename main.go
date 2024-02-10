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
	b := len(text)
	fmt.Println("Общее количество символов в исходном предложении", b)
	counts := make(map[rune]int)
	for _, char := range text {
		if unicode.IsSpace(char) {
			continue
		}
		char = unicode.ToLower(char)

		counts[char]++
	}
	fmt.Println("Количество повторений символов:")
	for char, count := range counts {
		var a float32
		c := float32(count)
		if count > 0 {
			if char != ' ' {
				fmt.Printf("Символ %c появляется %d раз\n", char, count)
				a = c / b
				fmt.Printf("Символ %c в процентном соотношении %0.4f\n", char, a)
			} else {
				continue
			}
		}
	}
}

// func main() {
// 	sentence := "Hello, world!"
// 	letter := 'o'

// 	css

// 	count := strings.Count(strings.ToLower(sentence), string(letter))
// 	totalLetters := len(strings.ReplaceAll(strings.ToLower(sentence), " ", ""))
// 	percentage := float64(count) / float64(totalLetters) * 100

// 	fmt.Printf("The letter %c appears %d times in the sentence.\n", letter, count)
// 	fmt.Printf("The frequency of the letter %c is %.2f%%\n", letter, percentage)

// for _, l := range clean {
// 	if unicode.IsSpace(l) {
// 		wordLine++
// 	}
// }
