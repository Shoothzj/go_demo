package demo_leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

func strToIntegerArray(str string) []int {
	arraySplit := strings.Split(str, ",")
	fmt.Println(arraySplit)
	result := make([]int, len(arraySplit))
	for i := 0; i < len(result); i++ {
		result[i], _ = strconv.Atoi(processStr(arraySplit[i]))
	}
	return result
}

func strToTwoDimensionIntegerArray(str string) [][]int {
	arraySplit := strings.Split(str, "],")
	fmt.Println("array split is ", arraySplit)
	result := make([][]int, len(arraySplit))
	for i := 0; i < len(result); i++ {
		valueSplit := strings.Split(arraySplit[i], ",")
		fmt.Printf("value i %d split is %s\n", i, valueSplit)
		result[i] = make([]int, len(valueSplit))
		for j := 0; j < len(valueSplit); j++ {
			rawS := valueSplit[j]
			result[i][j], _ = strconv.Atoi(processStr(rawS))
		}
	}
	return result
}

func processStr(raw string) string {
	var result []byte
	for _, elem := range raw {
		if elem >= '0' && elem <= '9' {
			result = append(result, byte(elem))
		}
	}
	return string(result)
}
