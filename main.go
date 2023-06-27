package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var roman bool

var romanSlice = []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X",
	"XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII", "XIX", "XX",
	"XXI", "XXII", "XXIII", "XXIV", "XXV", "XXVI", "XXVII", "XXVIII", "XXIX", "XXX",
	"XXXI", "XXXII", "XXXIII", "XXXIV", "XXXV", "XXXVI", "XXXVII", "XXXVIII", "XXXIX", "XL",
	"XLI", "XLII", "XLIII", "XLIV", "XLV", "XLVI", "XLVII", "XLVIII", "XLIX", "L",
	"LI", "LII", "LIII", "LIV", "LV", "LVI", "LVII", "LVIII", "LIX", "LX",
	"LXI", "LXII", "LXIII", "LXIV", "LXV", "LXVI", "LXVII", "LXVIII", "LXIX", "LXX",
	"LXXI", "LXXII", "LXXIII", "LXXIV", "LXXV", "LXXVI", "LXXVII", "LXXVIII", "LXXIX", "LXXX",
	"LXXXI", "LXXXII", "LXXXIII", "LXXXIV", "LXXXV", "LXXXVI", "LXXXVII", "LXXXVIII", "LXXXIX", "XC",
	"XCI", "XCII", "XCIII", "XCIV", "XCV", "XCVI", "XCVII", "XCVIII", "XCIX", "C"}

func CheckSign(sign string) {
	if sign != "/" && sign != "+" && sign != "-" && sign != "*" {
		log.Fatal("Вывод ошибки, так как оператор должен быть: (+, -, /, *).")
	}
}

func IsRoman(num string) bool {
	for _, v := range romanSlice {
		if v == num {
			return true
		}
	}
	return false
}

func RomanToInt(num string) int {
	for i, v := range romanSlice {
		if v == num {
			return i + 1
		}
	}
	return 0
}

func Parcer() (int, string, int) {
	var a, b int
	var err error
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	input := sc.Text()
	in := strings.Split(input, " ")
	if len(in) != 3 {
		if len(in) == 1 {
			log.Fatal("Вывод ошибки, так как строка не является математической операцией.\nМежду операндами и оператором должен быть пробел: \"a + b\".")
		}
		log.Fatal("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}
	CheckSign(in[1])
	if IsRoman(in[0]) && IsRoman(in[2]) {
		roman = true
		a = RomanToInt(in[0])
		b = RomanToInt(in[2])
	} else if !IsRoman(in[0]) && !IsRoman(in[2]) {
		if a, err = strconv.Atoi(in[0]); err != nil {
			log.Fatal("Вывод ошибки, так как строка не является математической операцией.")
		}
		if b, err = strconv.Atoi(in[2]); err != nil {
			log.Fatal("Вывод ошибки, так как строка не является математической операцией.")
		}
	} else {
		log.Fatal("Вывод ошибки, так как используются одновременно разные системы счисления.")
	}
	return a, in[1], b
}

func Calc(a int, sign string, b int) int {
	switch sign {
	case "/":
		if b == 0 {
			log.Fatal("Нельзя делить на 0.")
		}
		return a / b
	case "*":
		return a * b
	case "+":
		return a + b
	case "-":
		return a - b
	default:
		return 0
	}
}
func main() {
	a, sign, b := Parcer()
	if (a > 10 || a < 1) || (b > 10 || b < 1) {
		log.Fatal("Вывод ошибки, так как числа должны быть от 1 до 10.")
	}
	res := Calc(a, sign, b)
	if roman {
		if res <= 0 {
			log.Fatal("Вывод ошибки, так как в римской системе нет отрицательных чисел.")
		}
		fmt.Println(romanSlice[res-1])
	} else {
		fmt.Println(res)
	}
}
