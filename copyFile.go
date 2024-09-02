package main

import (
	"fmt"
	"io"
	"os"
)

// копируем файл
func copyFile(source string, dst string) bool {
	sourceFile, err := os.Open(source)
	if err != nil {
		fmt.Println("Ошибка открытия исходного файла ", err)
		return false
	}
	defer sourceFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		fmt.Println("Ошибка создания конечного файла ", err)
		return false
	}
	defer dstFile.Close()

	fmt.Println("Копируем файл ", sourceFile.Name())

	_, err = io.Copy(dstFile, sourceFile)
	if err != nil {
		fmt.Println("Ошибка копирования файла ", err)
		return false
	}
	return true
}

func delFile(filename string) bool {
	err := os.Remove(filename)
	if err != nil {
		fmt.Println("Ошибка удаления файла", err)
		return false
	}
	fmt.Println("Удаляем файл ", filename)
	return true
}
