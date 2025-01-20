package main

import (
	"fmt"
	"log"
)

func main() {
	// task 1
	// log.Println("Task 1")
	// log.Println(task1("aabbcc"))
	// log.Println(task1("aaabcccccaaa"))
	// log.Println(task1("abcd"))
	// log.Println(task1("baaabbbiiii"))
	// log.Println(task1(""))

	// // // task 2
	// log.Println("Task 2")
	// log.Println(task2("kakap", "kapak"))
	// log.Println(task2("sajen", "senja"))
	// log.Println(task2("kocak", "kacik"))
	// log.Println(task2("valas", "vaSal"))

	// // task 3
	// log.Println("Task 3")
	// log.Println(task3(10))
	// log.Println(task3(20))
	// log.Println(task3(1))
	// log.Println(task3(50))
	// log.Println(task3(100))

	// task 4
	log.Println("Task 4")
	log.Println(task4([]int{2, 7, 11, 15}, 9))
	log.Println(task4([]int{3, 2, 4}, 6))
	log.Println(task4([]int{3, 3}, 6))
	log.Println(task4([]int{7, 11, 2, 15}, 9))
}

type stringTask1 struct {
	Letter string
	Count  int
}

func task1(text string) string {
	mappedStrings := []*stringTask1{}
	for i := 0; i < len(text); i++ {
		letter := text[i : i+1]

		// get last entity from mapped strings
		lastMappedString := &stringTask1{}
		if len(mappedStrings) > 0 {
			lastMappedString = mappedStrings[len(mappedStrings)-1]
		}

		// check if last string letter is same with current letter
		if lastMappedString.Letter == letter {
			lastMappedString.Count++
		} else {
			mappedStrings = append(mappedStrings, &stringTask1{
				Letter: letter,
				Count:  1,
			})
		}
	}

	// get converted string from mapped strings
	convertedString := ""
	for _, mappedString := range mappedStrings {
		convertedString += fmt.Sprintf("%s%d", mappedString.Letter, mappedString.Count)
	}

	// check if result is longer than original text
	if len(convertedString) >= len(text) {
		return text
	}

	return convertedString
}

type stringTask2 struct {
	Letter string
	IsUsed bool
}

func task2(text1, text2 string) bool {
	// check if length of text1 and text2 is same
	if len(text1) != len(text2) {
		return false
	}

	splittedText1 := []*stringTask2{}

	// split text1 to array of stringTask2
	for i := 0; i < len(text1); i++ {
		splittedText1 = append(splittedText1, &stringTask2{
			Letter: text1[i : i+1],
			IsUsed: false,
		})
	}

	// loop text2 and check if letter is exist in text1
	for i := 0; i < len(text2); i++ {
		splittedText2 := text2[i : i+1]

		for _, letterText1 := range splittedText1 {
			if letterText1.Letter == splittedText2 && !letterText1.IsUsed {
				letterText1.IsUsed = true
				break
			}
		}
	}

	// check if all letter in text1 is used
	countUnusedLetter := 0
	for _, letterText1 := range splittedText1 {
		if !letterText1.IsUsed {
			countUnusedLetter++
		}
	}

	return countUnusedLetter == 0
}

func task3(n int) int {
	primeNumbers := []int{}
	for i := 1; i <= n; i++ {
		// check if number is prime number
		isPrime := true
		for j := 1; j <= i; j++ {
			if i%j == 0 && i != j && j != 1 {
				isPrime = false
				break
			}
		}

		// if it is prime number, add to prime numbers
		if isPrime && i != 1 {
			primeNumbers = append(primeNumbers, i)
		}
	}

	// calculate sum of prime numbers
	totalNumber := int(0)
	for _, primeNumber := range primeNumbers {
		totalNumber += primeNumber
	}

	return totalNumber
}

func task4(nums []int, target int) []int {
	// loop nums and check if sum of two number is equal to target
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i]+nums[j] == target && i != j {
				return []int{i, j}
			}
		}
	}

	// return -1, -1 if not found
	return []int{-1, -1}
}
