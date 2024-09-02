package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"
)

func execMagick(pathPDFfile string, fileName string) {
	//cmd := exec.Command("magick", "convert", "-density", "150", "-alpha", "remove", filePath, fileName)
	//cmd := exec.Command("magick", "convert", "-density", "109", "-alpha", "remove", "c:\\Users\\sviridovgm\\Desktop\\picture.pdf", "a.png")
	//cmd := exec.Command("magick", "convert -density 109 -alpha remove c:\\Users\\sviridovgm\\Desktop\\picture.pdf a.png")

	//Выполняем команду конвертирования с помошью popler
	cmd := exec.Command(filepath.Join(pathProgramm, "poppler", "pdftoppm.exe"), "-png", pathPDFfile, fileName)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("Ошибка ", err)
	}
	fmt.Println("Конвертировали в png")
}
