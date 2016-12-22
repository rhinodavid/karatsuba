package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(multiply(os.Args[1], os.Args[2]))
}

func multiply(x string, y string) string {
	if len(x) == 1 || len(y) == 1 {
		xi, _ := strconv.Atoi(x)
		yi, _ := strconv.Atoi(y)
		prod := strconv.Itoa(xi * yi)
		return prod
	}

	if len(x)%2 == 1 {
		x = "0" + x
	}
	if len(y)%2 == 1 {
		y = "0" + y
	}

	x, y = matchLengths(x, y)

	xm := len(x) / 2
	ym := len(y) / 2

	a := x[:xm]
	b := x[xm:]
	c := y[:ym]
	d := y[ym:]
	ac := multiply(a, c)
	bd := multiply(b, d)
	three := add(multiply(a, d), multiply(b, c))
	result := add(add(pad(ac, len(x)), pad(three, len(x)/2)), bd)
	return removeLeadingZeros(result)

}

func matchLengths(x string, y string) (string, string) {
	for len(x) > len(y) {
		y = "0" + y
	}
	for len(y) > len(x) {
		x = "0" + x
	}
	return x, y
}

func add(x string, y string) string {
	x, y = matchLengths(x, y)
	carry := 0
	result := ""
	for i := len(x) - 1; i >= 0; i-- {
		a, _ := strconv.Atoi(x[i : i+1])
		b, _ := strconv.Atoi(y[i : i+1])
		res := a + b + carry
		carry = 0
		if res >= 10 {
			carry = 1
			res = res - 10
		}
		resStr := strconv.Itoa(res)
		result = resStr + result
	}
	if carry == 1 {
		result = "1" + result
	}
	return result
}

func pad(x string, n int) string {
	result := x
	for i := 0; i < n; i++ {
		result = result + "0"
	}
	return result
}

func removeLeadingZeros(str string) string {
	if str == "0" {
		return str
	}
	for str[0:1] == "0" && len(str) > 1 {
		str = str[1:]
	}
	return str
}
