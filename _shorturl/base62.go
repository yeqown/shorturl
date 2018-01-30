package shorturl

import (
	// "fmt"
	"math"
	"strings"
)

const CODE62 = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const LENGTH = 62

var EDOC = map[string]int{
	"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
	"a": 10, "b": 11, "c": 12, "d": 13, "e": 14, "f": 15, "g": 16, "h": 17, "i": 18, "j": 19, "k": 20, "l": 21, "m": 22, "n": 23, "o": 24, "p": 25, "q": 26, "r": 27, "s": 28, "t": 29, "u": 30, "v": 31, "w": 32, "x": 33, "y": 34, "z": 35,
	"A": 36, "B": 37, "C": 38, "D": 39, "E": 40, "F": 41, "G": 42, "H": 43, "I": 44, "J": 45, "K": 46, "L": 47, "M": 48, "N": 49, "O": 50, "P": 51, "Q": 52, "R": 53, "S": 54, "T": 55, "U": 56, "V": 57, "W": 58, "X": 59, "Y": 60, "Z": 61,
}

/**
 * 编码整数为base62字符串
 */
func Encode(number int64) string {
	if number == 0 {
		return "0"
	}
	result := make([]byte, 0)
	for number > 0 {
		round := number / LENGTH
		remain := number % LENGTH
		result = append(result, CODE62[remain])
		number = round
	}
	// 逆置
	for start, end := 0, len(result)-1; start < end; start, end = start+1, end-1 {
		result[start], result[end] = result[end], result[start]
	}
	return string(result)
}

/**
 * 解码字符串为整数
 */
func Decode(str string) int64 {
	str = strings.TrimSpace(str)
	var result int = 0
	bs := []byte(str)
	var l int = len(bs)
	for _, ch := range bs {
		l = l - 1
		w := int(math.Pow(LENGTH, float64(l)))
		v := EDOC[string(ch)]
		// fmt.Printf("`%s` weight %d, value: %d, l: %d\n", str, w, v, l)
		result += (v * w)
	}
	return int64(result)
}
