package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

func compareDirPDF(wg *sync.WaitGroup) { //функция сравнивает содержимое исходной папки с рабочей папкой программы
	for {
		<-time.After(1 * time.Minute)           //цикл выполняется примерно каждую 1 минуту
		files, err := ioutil.ReadDir(sourceDir) //получаем список файлов исходной папки
		if err != nil {
			fmt.Println("Ошибка просмотра исходной папки ", err)
		}
		for _, file := range files { //цикл по всем файлам
			if !file.IsDir() {
				if strings.HasSuffix(file.Name(), ".pdf") { //если файл пдф то начинаем его конвертацию
					sourceFileStat, err := os.Stat(filepath.Join(sourceDir, file.Name())) //получаем статы файла что бы сравнивать размер при совпадении имени
					if err != nil {
						fmt.Println("Неудалось прочитать атрибуты файла", filepath.Join(sourceDir, file.Name()))
					}

					dstFileStat, err := os.Stat(filepath.Join(pdfDir, file.Name())) //получаем статы файла в рабочей папке
					if err != nil {
						if os.IsNotExist(err) { //если такого файла нет то копируем его в рабочую папку
							copyFile(filepath.Join(sourceDir, file.Name()), filepath.Join(pdfDir, file.Name()))
							copyFile(filepath.Join(pdfDir, file.Name()), filepath.Join(tmpDir, file.Name()))
						}
					} else { //делаем тоже самое если оба файла есть но отличаются, в данном случае размерами
						if sourceFileStat.Size() != dstFileStat.Size() {
							copyFile(filepath.Join(sourceDir, file.Name()), filepath.Join(pdfDir, file.Name()))
							copyFile(filepath.Join(pdfDir, file.Name()), filepath.Join(tmpDir, file.Name()))
						}
					}

				}
				if strings.HasSuffix(file.Name(), ".mp4") { //если в исходной папке есть файл mp4 то копируем его в рабочую папку и папку с видео минуя темп
					copyFile(filepath.Join(sourceDir, file.Name()), filepath.Join(pdfDir, file.Name()))
					copyFile(filepath.Join(pdfDir, file.Name()), filepath.Join(videoDir, file.Name()))
				}
			}
		}
		files, err = ioutil.ReadDir(pdfDir) //теперь просматриваем уже рабочую папку и сравниваем с исходной
		if err != nil {
			fmt.Println("Ошибка просмотра ПДФ папки ", err)
		}
		for _, file := range files {
			if !file.IsDir() {
				_, err := os.Stat(filepath.Join(sourceDir, file.Name()))
				if err != nil {
					if os.IsNotExist(err) { //если файл не существует в исходной папке то удаляем его и из рабочей
						delFile(filepath.Join(pdfDir, file.Name()))
					}
				}
			}
		}
	}
	wg.Done()
}
