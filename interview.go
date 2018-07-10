package practice

import (
	"sort"
	"strings"
)

func SortString(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

// CommonChars returns a string of all the common characters in both provided strings.
func CommonChars(str1, str2 string) string {
	n1 := len(str1)
	n2 := len(str2)
	if n1 < 1 || n2 < 1 {
		return ""
	}
	sorted1 := SortString(str1)
	sorted2 := SortString(str2)

	response := ""
	for i1, i2 := 0, 0; i1 < n1 && i2 < n2; {
		if sorted1[i1] == sorted2[i2] {
			response = response + string(sorted1[i1])
			i1++
			i2++
		} else if sorted1[i1] > sorted2[i2] {
			i2++
		} else {
			i1++
		}
	}
	return response
}

// CommonChars2 returns a string of all the common characters in both provided strings.
// ("boo", "cat"): ""
// ("boo", "cot"): "o"
// ("book", "crabby"): "b"
// ("book", "kangaroo"): "ko"|"ok"
func CommonChars2(str1, str2 string) string {
	if len(str1) == 0 && len(str2) == 0 {
		return ""
	}
	str1Chars := make(map[rune]int, 0)
	for _, c := range str1 {
		str1Chars[c] = 0
	}
	for _, c := range str2 {
		if _, ok := str1Chars[c]; ok {
			str1Chars[c] += 1
		}
	}
	response := []rune{}
	for k, v := range str1Chars {
		if v > 0 {
			response = append(response, k)
		}
	}
	return string(response)
}

// Mapper to make the value to the keys of a map.
//If the input list is `1, 2` and the input function is `x * 2`, the return value is `2: 1, 4: 2`
func Mapper(list []int, fnc func(int) int) map[int]int {
	response := make(map[int]int)
	for _, v := range list {
		response[fnc(v)] = v
	}
	return response
}

// Mapper2 is a concurrent version of Mapper
func Mapper2(list []int, fnc func(int) int) map[int]int {
	n := len(list)
	if n == 0 {
		return make(map[int]int)
	}
	spread := make(chan [2]int)
	for _, v := range list {
		go func(v int) {
			spread <- [...]int{fnc(v), v}
		}(v)
	}

	response := make(map[int]int)
	for n > 0 {
		res := <-spread
		response[res[0]] = res[1]
		n--
	}
	return response
}
