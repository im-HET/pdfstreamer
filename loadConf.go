package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func loadConf() {
	var err error
	pathProgramm, err = os.Executable()
	if err != nil {
		fmt.Println("Не удается определить рабочую папку программы")
		os.Exit(1)
	}
	pathProgramm = filepath.Dir(pathProgramm)
	pdfDir = filepath.Join(pathProgramm, "pdf")
	videoDir = filepath.Join(pathProgramm, "video")
	tmpDir = filepath.Join(pathProgramm, "tmp")
	semplDir = filepath.Join(pathProgramm, "sempls")
	readFile, err := os.Open(filepath.Join(pathProgramm, "pdfstreamer.conf"))
	if err != nil {
		fmt.Println("Ошибка открытия конфигурации")
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	//создаем хэш таблицу с параметром и значением
	params := map[string]string{}
	for fileScanner.Scan() {
		s := strings.Split(fileScanner.Text(), "=")
		params[s[0]] = s[1]
	}
	readFile.Close()

	//в хэш таблице ищем нужные нам параметры, если есть то присваиваем
	param, isExistKey := params["sourceDir"]
	if isExistKey {
		sourceDir = param
		fmt.Println("Исходная папка ", sourceDir)
	}
	param, isExistKey = params["showtime"]
	if isExistKey {
		showtime = param
		fmt.Println("Длинна показа страницы ", showtime)
	}
	param, isExistKey = params["addSempls"]
	if isExistKey {
		addSempls = param
		fmt.Println("Добавлять звуковую дорожку ", addSempls)
	}
	param, isExistKey = params["semplDir"]
	if isExistKey {
		semplDir = param
		fmt.Println("Папка семплов ", semplDir)
	}
	param, isExistKey = params["dstip"]
	if isExistKey {
		dstip = param
		fmt.Println("Адрес получателя ", dstip)
	}
	param, isExistKey = params["dstport"]
	if isExistKey {
		dstport = param
		fmt.Println("Адрес получателя ", dstport)
	}
	fmt.Println("Папка программы ", pathProgramm)
}
