package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Проверяем, что передано два аргумента
	if len(os.Args) < 3 {
		fmt.Println("Использование: go run multiply.go <число1> <число2>")
		return
	}

	// Читаем аргументы командной строки и конвертируем в числа
	num1, err1 := strconv.Atoi(os.Args[1])
	num2, err2 := strconv.Atoi(os.Args[2])

	if err1 != nil || err2 != nil {
		fmt.Println("Ошибка: введите два целых числа")
		return
	}

	// Вычисляем результат
	result := num1 * num2

	// Записываем результат в файл
	dirPath := "/multiply_result"
	// Создаем все необходимые директории
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		fmt.Println("Ошибка при создании директорий:", err)
		return
	}

	file, err := os.Create(dirPath + "/result.txt")
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("%d\n", result))
	if err != nil {
		fmt.Println("Ошибка при записи в файл:", err)
		return
	}

	fmt.Println("Результат записан в result.txt")
}

