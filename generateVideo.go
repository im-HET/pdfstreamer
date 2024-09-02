package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

func generateVideo(wg *sync.WaitGroup) {
	for {
		<-time.After(1 * time.Minute) //с остановкой примерно в минуту просматриваем темп папку на наличее пдф файлов
		files, err := ioutil.ReadDir(tmpDir)
		if err != nil {
			fmt.Println("Ошибка просмотра исходной папки ", err)
		}
		for _, file := range files {
			if !file.IsDir() {
				if strings.HasSuffix(file.Name(), ".pdf") {
					namePDF := filepath.Join(tmpDir, file.Name())
					nameMP4 := filepath.Join(tmpDir, strings.ReplaceAll(file.Name(), ".pdf", ".mp4"))
					name := filepath.Join(tmpDir, strings.ReplaceAll(file.Name(), ".pdf", ""))

					execMagick(namePDF, name) //пдф конвертируем в картинки
					execFfmpeg(name, nameMP4) //из картинок создаем видео файл

					fmt.Println("Копируем получившееся видео")
					if copyFile(nameMP4, filepath.Join(videoDir, strings.ReplaceAll(file.Name(), ".pdf", ".mp4"))) { //копируем получившееся видео в видео папку
						delFile(namePDF) //если копирование успешно то удалем пдф файл из темп папки
					}
					//если копирование получившегося видео файла не удалось выше то при следующей итерации цикла генерация пдф файла выполнится заного
					delFile(nameMP4) //удаляем получившееся видео в любом случае
				}
				clearTemp() //удаляем картинки
			}
		}
	}
	wg.Done()
}

func clearTemp() { // удаляем все пнг файлы
	files, err := ioutil.ReadDir(tmpDir)
	if err != nil {
		fmt.Println("Ошибка просмотра исходной папки ", err)
	}
	for _, file := range files {
		if !file.IsDir() {
			if strings.HasSuffix(file.Name(), ".png") {
				delFile(filepath.Join(tmpDir, file.Name()))
			}
		}
	}
}
