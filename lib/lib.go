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
