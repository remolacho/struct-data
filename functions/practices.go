package functions

import (
	"fmt"
	"struct-data/arrays/arrayInteger"
)

func ArrayMethods() {
	numbers := arrayInteger.New()
	numbers.Rand(20)
	numbers.Add(-5)
	numbers.Add(9)

	//fmt.Printf("%v Size \n", numbers.Size())
	//fmt.Printf("%v Min \n", numbers.Min())
	//fmt.Printf("%v Max \n", numbers.Max())
	//fmt.Printf("%v Avg \n", numbers.Avg())
	//fmt.Printf("%v Sum \n", numbers.Sum())
	//fmt.Printf("%v FindAll \n", numbers.FindAll(9))
	//fmt.Printf("%v Find \n", numbers.Find(9))
	//fmt.Printf("%v Detect \n", numbers.Detect(9))
	//fmt.Printf("%v Sort Asc \n", numbers.SortAsc())
	//fmt.Printf("%v Sort Desc \n", numbers.SortDesc())
	//fmt.Printf("%v Sort Selection \n", numbers.SortSelection())
	//fmt.Printf("%v Sort Insertion \n", numbers.SortByInsertion())
	//fmt.Printf("%v Sort SortByDoubleRun \n", numbers.SortByDoubleRun())
	fmt.Printf("%v DeleteInCopy element \n", numbers.DeleteInCopy(7))
	fmt.Printf("%v Origin Values \n", numbers.Values())
	fmt.Printf("%v Delete element \n", numbers.Delete(7))
	fmt.Printf("%v Origin Values \n", numbers.Values())

}

func Palindrome(str string) {
	var strReverse string
	chars := []rune(str)

	for i := len(chars) - 1; i >= 0; i-- {
		strReverse += string(chars[i])
	}

	if str == strReverse {
		println(str + " " + "es palindrome")
	} else {
		println(str + " " + "no es palindrome")
	}
}

func Palindrome2(str string) {
	isPal := true
	chars := []rune(str)

	for i := 0; i <= (len(chars)-1)/2 && isPal; i++ {
		if string(chars[i]) != string(chars[len(chars)-i-1]) {
			isPal = false
		}
	}

	if isPal {
		println(str + " " + "es palindrome")
	} else {
		println(str + " " + "no es palindrome")
	}
}

func PerfectNumber(num int) {
	if num%2 != 0 {
		fmt.Printf("%d no es un numero perfecto \n", num)
		return
	}

	top := num / 2
	quantity := 0

	for i := 1; i <= top; i++ {
		if num%i == 0 {
			quantity += i
		}
	}

	if quantity == num {
		fmt.Printf("%d es un numero perfecto \n", num)
	} else {
		fmt.Printf("%d no es un numero perfecto \n", num)
	}
}

// restar un numero sin restar
func SubtNumber(a int, b int) {
	var tmp int
	value := 0
	convertion := 1

	if b > a {
		tmp, b, convertion = b, a, -1
		a = tmp
	}

	fmt.Printf("%d - %d", a, b)

	for i := 0; i < a; i++ {
		if a > b {
			value += 1
		}
		b += 1
	}

	fmt.Printf(" = %d \n", value*convertion)
}

func NumberGuay(a int) {
	value := 0
	isGuay := false

	for i := 1; i < a && isGuay == false; i++ {
		value += i

		if value == a {
			isGuay = true
		} else if value > a {
			break
		}
	}

	if isGuay {
		fmt.Printf("%d es un numero guay \n", a)
	} else {
		fmt.Printf("%d no es un numero guay \n", a)
	}
}

func OpenAndClose(str string, strOpen string) {
	chars := []rune(str)
	var countOpen int
	var countClose int

	for i := 0; i <= len(chars)-1; i++ {
		char := string(chars[i])

		if char == strOpen {
			countOpen++
		} else {
			countClose++
		}
	}

	if countOpen == countClose {
		fmt.Printf("Esta homolagada \n")
	} else {
		fmt.Printf("No esta homolagada \n")
	}
}
