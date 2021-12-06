package lib

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadInputAsStr(value int) []string {
	data, err := ioutil.ReadFile("q" + strconv.Itoa(value) + ".txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return nil
	}

	str_values := strings.Split(string(data), "\r\n")
	return str_values
}

func ReadInputAsInt(value int) []int {
	str_values := ReadInputAsStr(value)
	ary := make([]int, len(str_values))
	for i := range ary {
		ary[i], _ = strconv.Atoi(str_values[i])
	}
	return ary
}

func ReadCSVAsInt(value int) []int {
	str_values := ReadInputAsStr(value)
	str_values = strings.Split(str_values[0], ",")
	ary := make([]int, len(str_values))
	var err error
	for i := range str_values {
		ary[i], err = strconv.Atoi(str_values[i])
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println(ary)
	return ary
}
