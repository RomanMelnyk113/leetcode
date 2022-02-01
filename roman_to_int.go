package main

var rti = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func isDeduct(v string, n string) bool {
	switch {
	case v == "I" && (n == "V" || n == "X"):
		return true
	case v == "X" && (n == "L" || n == "C"):
		return true
	case v == "C" && (n == "D" || n == "M"):
		return true
	}
	return false
}

func romanToInt(s string) int {
	sum := 0
	l := len(s)
	var n string
	for i := 0; i < l; i++ {
		v := string(s[i])
		if i+1 != l {
			n = string(s[i+1])
		} else {
			n = ""
		}
		if isDeduct(v, n) {
			sum -= rti[v]
		} else {
			sum += rti[v]
		}
	}
	return sum
}
