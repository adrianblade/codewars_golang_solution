package src

import (
    "strings"
    "strconv"
)

func HighAndLow(in string) string {
  arr := ToIntArray(in)
  min, max := MinMax(arr)
  return strconv.Itoa(max) + " " + strconv.Itoa(min)
}

func ToIntArray(A string) []int {
	strs := strings.Split(A, " ")
	ary := make([]int, len(strs))
	for i := range ary {
		ary[i], _ = strconv.Atoi(strs[i])
	}  
	return ary
}

func MinMax(array []int) (int, int) {
    var max int = array[0]
    var min int = array[0]
    for _, value := range array {
        if max < value {
            max = value
        }
        if min > value {
            min = value
        }
    }
    return min, max
}