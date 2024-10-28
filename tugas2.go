package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CountWords(s string) int {
	return len(strings.Fields(s))
}

func reverseString(str string) (result string) {
	for _ , i := range str {
		result = string(i) + result
	}
	return
}

func main() {

	input := bufio.NewReader(os.Stdin)

	for{
		fmt.Println("Masukkan Minimal 3 Kata: ")
		text, _ := input.ReadString('\n')
		JumlahKata := CountWords(text)

		if JumlahKata < 3 {
			fmt.Println("Kata yang dimasukkan kurang dari 3 (click enter to retry)")
			text, _ = input.ReadString('\n')
			continue
		} 

		if JumlahKata >= 3 {
			fmt.Println("\nJumlah Kata: ", JumlahKata)
			fmt.Println("Kata yang dimasukkan: ", text)
			fmt.Println("Reversed Strings: ", reverseString(text))
			break
		}
	}
}