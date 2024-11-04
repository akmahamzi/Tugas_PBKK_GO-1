package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Fungsi untuk membalik kata
func reverseWord(word string) string {
	runes := []rune(word)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan kalimat: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	words := strings.Fields(input)

	if len(words) < 3 {
		fmt.Println("Error: Minimal 3 kata mazz.")
		return
	}

	for i, word := range words {
		words[i] = reverseWord(word)
	}

	result := strings.Join(words, " ")
	fmt.Println("Hasil:", result)
}
