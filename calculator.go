package main

import (
	"fmt"
	"strconv"
	"strings"
)

var input string

func main() {
	fmt.Println("Введите выражение")
	fmt.Scanln(&input)
	//fmt.Println(in

	// splitFunc вызывается для каждой
	// руны в строке. Если руна
	// равна любому символу в "*%,_"
	// input разделяется.
	splitFunc := func(r rune) bool {
		return strings.ContainsRune("+,-,/,*", r)
	}

	words := strings.FieldsFunc(input, splitFunc)
	if len(words) != 2 {
		fmt.Println("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).\n")
		return

	}
	var romeOne int
	var one int
	var err error
	switch words[0] {
	case "I":
		romeOne = 1
	case "II":
		romeOne = 2
	case "III":
		romeOne = 3
	case "IV":
		romeOne = 4
	case "V":
		romeOne = 5
	case "VI":
		romeOne = 6
	case "VII":
		romeOne = 7
	case "VIII":
		romeOne = 8
	case "IX":
		romeOne = 9
	case "X":
		romeOne = 10
	default:
		one, err = strconv.Atoi(words[0])
		if err != nil {
			panic(err)
		}
		if one >= 11 {
			fmt.Println("Вывод ошибки, так как формат математической операции не удовлетворяет заданию-принимаются целые числа от 1 до 10 включительно")
			return

		}
	}
	var romeTwo int
	var two int
	var err2 error
	switch words[1] {
	case "I":
		romeTwo = 1
	case "II":
		romeTwo = 2
	case "III":
		romeTwo = 3
	case "IV":
		romeTwo = 4
	case "V":
		romeTwo = 5
	case "VI":
		romeTwo = 6
	case "VII":
		romeTwo = 7
	case "VIII":
		romeTwo = 8
	case "IX":
		romeTwo = 9
	case "X":
		romeTwo = 10
	default:

		two, err2 = strconv.Atoi(words[1])
		if err2 != nil {
			panic(err2)
		}
		if two >= 11 {
			fmt.Println("Вывод ошибки, так как формат математической операции не удовлетворяет заданию-принимаются целые числа от 1 до 10 включительно")
			return
		}

	}
	//var summ int
	var summRome int
	if strings.Contains(input, "+") {
		if one >= 1 && one < 11 && two >= 1 && two < 11 {
			fmt.Println(one + two)
		} else if romeOne != 0 && romeTwo != 0 {
			summRome = romeOne + romeTwo
			if summRome < 1 {
				fmt.Println("Вывод ошибки, так как в римской системе нет отрицательных чисел")

			} else if summRome == 1 {
				fmt.Println("I")
			} else if summRome == 2 {
				fmt.Println("II")
			} else if summRome == 3 {
				fmt.Println("III")
			} else if summRome == 4 {
				fmt.Println("IV")
			} else if summRome == 5 {
				fmt.Println("V")
			} else if summRome == 6 {
				fmt.Println("VI")
			} else if summRome == 7 {
				fmt.Println("VII")
			} else if summRome == 8 {
				fmt.Println("VIII")
			} else if summRome == 9 {
				fmt.Println("IX")
			} else if summRome == 10 {
				fmt.Println("x")
			} else if summRome == 11 {
				fmt.Println("XI")
			} else if summRome == 12 {
				fmt.Println("XII")
			} else if summRome == 13 {
				fmt.Println("XIII")
			} else if summRome == 14 {
				fmt.Println("XIV")
			} else if summRome == 15 {
				fmt.Println("XV")
			} else if summRome == 16 {
				fmt.Println("XVI")
			} else if summRome == 17 {
				fmt.Println("XVII")
			} else if summRome == 18 {
				fmt.Println("XVIII")
			} else if summRome == 19 {
				fmt.Println("XIX")
			} else if summRome == 20 {
				fmt.Println("XX")
			}

		} else {
			fmt.Println("Вывод ошибки, так как используются одновременно разные системы счисления.")
		}
		//summ = one + two
	}
	if strings.Contains(input, "-") {
		if one >= 1 && one < 11 && two >= 1 && two < 11 {
			fmt.Println(one - two)
		} else if romeOne != 0 && romeTwo != 0 {
			summRome = romeOne - romeTwo
			if summRome < 1 {
				fmt.Println("Вывод ошибки, так как в римской системе нет отрицательных чисел")

			} else if summRome == 1 {
				fmt.Println("I")
			} else if summRome == 2 {
				fmt.Println("II")
			} else if summRome == 3 {
				fmt.Println("III")
			} else if summRome == 4 {
				fmt.Println("IV")
			} else if summRome == 5 {
				fmt.Println("V")
			} else if summRome == 6 {
				fmt.Println("VI")
			} else if summRome == 7 {
				fmt.Println("VII")
			} else if summRome == 8 {
				fmt.Println("VIII")
			} else if summRome == 9 {
				fmt.Println("IX")
			} else if summRome == 10 {
				fmt.Println("x")
			} else if summRome == 11 {
				fmt.Println("XI")
			} else if summRome == 12 {
				fmt.Println("XII")
			} else if summRome == 13 {
				fmt.Println("XIII")
			} else if summRome == 14 {
				fmt.Println("XIV")
			} else if summRome == 15 {
				fmt.Println("XV")
			} else if summRome == 16 {
				fmt.Println("XVI")
			} else if summRome == 17 {
				fmt.Println("XVII")
			} else if summRome == 18 {
				fmt.Println("XVIII")
			} else if summRome == 19 {
				fmt.Println("XIX")
			} else if summRome == 20 {
				fmt.Println("XX")
			}

		} else {
			fmt.Println("Вывод ошибки, так как используются одновременно разные системы счисления.")
		}
		//else if strings.Contains(input, "-") {
		//summ = one - two

	}
	if strings.Contains(input, "*") {
		if one >= 1 && one < 11 && two >= 1 && two < 11 {
			fmt.Println(one * two)
		} else if romeOne != 0 && romeTwo != 0 {
			summRome = romeOne * romeTwo
			if summRome < 1 {
				fmt.Println("Вывод ошибки, так как в римской системе нет отрицательных чисел")

			} else if summRome == 1 {
				fmt.Println("I")
			} else if summRome == 2 {
				fmt.Println("II")
			} else if summRome == 3 {
				fmt.Println("III")
			} else if summRome == 4 {
				fmt.Println("IV")
			} else if summRome == 5 {
				fmt.Println("V")
			} else if summRome == 6 {
				fmt.Println("VI")
			} else if summRome == 7 {
				fmt.Println("VII")
			} else if summRome == 8 {
				fmt.Println("VIII")
			} else if summRome == 9 {
				fmt.Println("IX")
			} else if summRome == 10 {
				fmt.Println("x")
			} else if summRome == 11 {
				fmt.Println("XI")
			} else if summRome == 12 {
				fmt.Println("XII")
			} else if summRome == 13 {
				fmt.Println("XIII")
			} else if summRome == 14 {
				fmt.Println("XIV")
			} else if summRome == 15 {
				fmt.Println("XV")
			} else if summRome == 16 {
				fmt.Println("XVI")
			} else if summRome == 17 {
				fmt.Println("XVII")
			} else if summRome == 18 {
				fmt.Println("XVIII")
			} else if summRome == 19 {
				fmt.Println("XIX")
			} else if summRome == 20 {
				fmt.Println("XX")
			} else if summRome == 21 {
				fmt.Println("XXI")

			} else if summRome == 24 {
				fmt.Println("XXIV")
			} else if summRome == 25 {
				fmt.Println("XXV")
			} else if summRome == 27 {
				fmt.Println("XXVII")
			} else if summRome == 28 {
				fmt.Println("XXVIII")
			} else if summRome == 30 {
				fmt.Println("XXX")
			} else if summRome == 32 {
				fmt.Println("XXXII")
			} else if summRome == 35 {
				fmt.Println("XXXV")
			} else if summRome == 36 {
				fmt.Println("XXXVI")
			} else if summRome == 40 {
				fmt.Println("XL")
			} else if summRome == 42 {
				fmt.Println("XLII")
			} else if summRome == 45 {
				fmt.Println("XLV")
			} else if summRome == 48 {
				fmt.Println("XLVIII")
			} else if summRome == 49 {
				fmt.Println("XLIX")
			} else if summRome == 50 {
				fmt.Println("L")
			} else if summRome == 54 {
				fmt.Println("LIV")
			} else if summRome == 56 {
				fmt.Println("LVI")
			} else if summRome == 60 {
				fmt.Println("LX")
			} else if summRome == 63 {
				fmt.Println("LXIII")
			} else if summRome == 64 {
				fmt.Println("LXIV")
			} else if summRome == 70 {
				fmt.Println("LXX")
			} else if summRome == 72 {
				fmt.Println("LXXII")
			} else if summRome == 80 {
				fmt.Println("LXXX")
			} else if summRome == 81 {
				fmt.Println("LXXXI")
			} else if summRome == 90 {
				fmt.Println("XC")
			} else if summRome == 100 {
				fmt.Println("C")
			}

		} else {
			fmt.Println("Вывод ошибки, так как используются одновременно разные системы счисления.")
		}

	}
	if strings.Contains(input, "/") {
		if one >= 1 && one < 11 && two >= 1 && two < 11 {
			fmt.Println(one / two)
		} else if romeOne != 0 && romeTwo != 0 {
			summRome = romeOne / romeTwo
			if summRome < 1 {
				fmt.Println("Вывод ошибки, так как в римской системе нет отрицательных чисел")

			} else if summRome == 1 {
				fmt.Println("I")
			} else if summRome == 2 {
				fmt.Println("II")
			} else if summRome == 3 {
				fmt.Println("III")
			} else if summRome == 4 {
				fmt.Println("IV")
			} else if summRome == 5 {
				fmt.Println("V")
			} else if summRome == 6 {
				fmt.Println("VI")
			} else if summRome == 7 {
				fmt.Println("VII")
			} else if summRome == 8 {
				fmt.Println("VIII")
			} else if summRome == 9 {
				fmt.Println("IX")
			} else if summRome == 10 {
				fmt.Println("x")
			} else if summRome == 11 {
				fmt.Println("XI")
			} else if summRome == 12 {
				fmt.Println("XII")
			} else if summRome == 13 {
				fmt.Println("XIII")
			} else if summRome == 14 {
				fmt.Println("XIV")
			} else if summRome == 15 {
				fmt.Println("XV")
			} else if summRome == 16 {
				fmt.Println("XVI")
			} else if summRome == 17 {
				fmt.Println("XVII")
			} else if summRome == 18 {
				fmt.Println("XVIII")
			} else if summRome == 19 {
				fmt.Println("XIX")
			} else if summRome == 20 {
				fmt.Println("XX")
			}

		} else {
			fmt.Println("Вывод ошибки, так как используются одновременно разные системы счисления.")
		}
		//summ = one + two
	}

	//	fmt.Println(romeOne, one, romeTwo, two, summ, summRome)

}
