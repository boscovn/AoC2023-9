package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getNextValue(values []int) int {
	var diffs []int
	allZeroes := values[len(values)-1] == 0
	for i := 0; i < len(values)-1; i++ {
		diff := values[i+1] - values[i]
		allZeroes = allZeroes && values[i] == 0
		diffs = append(diffs, diff)
	}
	if allZeroes {
		return 0
	}
	return values[0] - getNextValue(diffs)
}

func convertList(list []string) []int {
	var intList []int
	for _, v := range list {
		num, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		intList = append(intList, num)
	}
	return intList
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	for scanner.Scan() {
		text := scanner.Text()
		s := strings.Split(text, " ")
		sum += getNextValue(convertList(s))

	}

	fmt.Println(sum)
}
