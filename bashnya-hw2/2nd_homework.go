package main

import "fmt"

const MAX_VALUE = 12307

func main() {
	var input int
	var err error

	fmt.Scanf("%d", &input)

	if input > MAX_VALUE {
		err = fmt.Errorf("Wrong number. Please, input number < 12307")
	} else {
		for input < MAX_VALUE && err == nil {
			if input%13 == 0 && input%9 == 0 {
				err = fmt.Errorf("service error")
			} else {
				input++
			}

			switch {
			case input < 0:
				input = -input
			case input%7 == 0:
				input *= 39
			case input%9 == 0:
				input *= 13
				input++
				continue
			default:
				input = (input + 2) * 3
			}
		}
	}

	if err == nil {
		fmt.Printf("Результат вычислений: %d", input)
	} else {
		fmt.Println(err)
	}
}
