package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	words []string
)

type InputDataStruct struct {
	yellowLetters []string
	blackLetters  []string
	greenLetters  []string
}

func ReadWordsFromFile() ([]string, error) {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
		return words, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	return words, err
}

func LettersExist(word string, letters []string) bool {
	for _, v := range letters {
		if !strings.Contains(word, v) {
			return false
		}
	}
	return true
}

func LettersNotExist(word string, letters []string) bool {
	for _, v := range letters {
		if strings.Contains(word, v) {
			return false
		}
	}
	return true
}

func IsLetterPosition(word string, letters []string) bool {

	wordTmp := strings.Split(word, "")
	for k, v := range letters {
		if !(v < "a" || v > "z") {
			if wordTmp[k] != letters[k] {
				return false
			}
		}
	}

	return true
}

func Calculate(letters InputDataStruct) (words []string) {
	wordsTmp, err := ReadWordsFromFile()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, word := range wordsTmp {
		str := []rune(word)
		if len(str) == 5 {
			if IsLetterPosition(word, letters.greenLetters) &&
				LettersExist(word, letters.yellowLetters) &&
				LettersNotExist(word, letters.blackLetters) {
				words = append(words, word)
			}
		}

	}

	return
}

func main() {

	//	enter youur leters
	i := InputDataStruct{
		yellowLetters: []string{"a"},
		blackLetters:  []string{"k", "r", "t", "y"},
		greenLetters:  []string{"-", "-", "b", "-", "-"}}

	ret := Calculate(i)
	fmt.Println(ret)

}
