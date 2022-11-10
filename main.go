package main

import "fmt"

func scanNumber() (int, error) {
	var a int
	_, err := fmt.Scanf("%d", &a)
	if err != nil {
		return 0, fmt.Errorf("scanf error: %s", err.Error())
	}
	return a, nil
}

func main() {
	for {

		var what string

		fmt.Println("Выберите действие (+, -, *, /)")
		fmt.Scanf("%s\n", &what)
		if what == "exit" {
			return
		}

		fmt.Println("Введите первое значение: ")

		a, err := scanNumber()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println("Введите второе значение: ")
		b, err := scanNumber()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		var c int
		if what == "+" {
			c = a + b
			fmt.Println("Результат: ", c)
		} else if what == "-" {
			c = a - b
			fmt.Println("Результат: ", c)
		} else if what == "*" {
			c = a * b
			fmt.Println("Результат: ", c)
		} else if what == "/" {
			c = a / b
			fmt.Println("Результат: ", c)
		}
	}
}
