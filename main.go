package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите выражение для строкового калькулятора (например, \"abc\" + \"def\", \"abc\" - \"b\", \"abc\" * 3, \"abc\" / 2):")
	expression, _ := reader.ReadString('\n')
	expression = strings.TrimSpace(expression)

	parts := strings.Split(expression, " ")

	if len(parts) < 3 {
		fmt.Println("Ошибка! Неверный формат выражения. Попробуйте еще раз.")
		return
	}

	var op1, operator, op2 string

	if len(parts) > 3 {
		op1 = strings.Trim(parts[0], "\"") + " " + strings.Trim(parts[1], "\"")
		operator = parts[2]
		op2 = strings.Trim(parts[3], "\"")
	} else {
		op1 = strings.Trim(parts[0], "\"")
		operator = parts[1]
		op2 = strings.Trim(parts[2], "\"")
	}

	switch operator {
	case "+":
		result := op1 + op2
		if len(result) > 40 {
			result = result[:40] + "..."
		}
		fmt.Printf("Результат: %q", result)
	case "-":
		result := strings.Replace(op1, op2, "", -1)
		if len(result) > 40 {
			result = result[:40] + "..."
		}
		fmt.Printf("Результат: %q", result)
	case "*":
		n, _ := strconv.Atoi(op2)
		result := strings.Repeat(op1, n)
		if len(result) > 40 {
			result = result[:40] + "..."
		}
		fmt.Printf("Результат: %q", result)
	case "/":
		n, _ := strconv.Atoi(op2)
		result := op1[:len(op1)/n]
		if len(result) > 40 {
			result = result[:40] + "..."
		}
		fmt.Printf("Результат: %q", result)
	default:
		fmt.Println("Ошибка! Неправильный оператор. Попробуйте еще раз.")
	}
}
