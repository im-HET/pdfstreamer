package main

import (
	"bytes"
	"fmt"
	"io/fs"
	"io/ioutil"
	"math/rand"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func execFfmpeg(fileName string, outFileName string) {
	//cmd := exec.Command("ffmpeg", "-r", "1/5", "-i", "c:\\Users\\sviridovgm\\Desktop\\a-%02d.png", "-r", "30",  "out.mp4")
	//cmd := exec.Command("ffmpeg", "-r", "1/"+showtime, "-i", fileName+"-%01d.png", "-r", "30", "-vcodec", "mpeg4", "-vf", "scale=1280:1024", outFileName)
	//cmd := exec.Command("./ffmpeg/ffmpeg", "-r", "1/"+showtime, "-i", fileName+"-%01d.png", "-r", "25", "-vcodec", "h264", "-vf", "scale=1366:768,format=yuv420p", "-aspect", "16:9", outFileName)

	//получаем количество изображений в папке
	numPdfFiles := 0
	files, err := ioutil.ReadDir(tmpDir)
	if err != nil {
		fmt.Println("Ошибка просмотра исходной папки ", err)
	}
	for _, file := range files {
		if !file.IsDir() {
			if strings.HasSuffix(file.Name(), ".png") {
				numPdfFiles += 1
			}
		}
	}
	//если количество изображений больше 10 то немного меняем синтаксис каманды ffmpeg
	ffmpegParamFileName := "-%01d.png"
	if numPdfFiles >= 10 {
		ffmpegParamFileName = "-%02d.png"
	}
	//создаем видео из набора картинок
	cmd := exec.Command(filepath.Join(pathProgramm, "ffmpeg", "ffmpeg.exe"), "-r", "1/"+showtime, "-i", fileName+ffmpegParamFileName, "-r", "25", "-vcodec", "h264", "-vf", "scale=1366:768,format=yuv420p", "-aspect", "16:9", filepath.Join(tmpDir, "tmp.mp4"))
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		fmt.Println("Ошибка ", err)
	}

	if addSempls == "yes" { //если параметр установлен в "yes" то к видео добавляем семпл
		var sempls []fs.FileInfo              //создаем список семплов
		files, err = ioutil.ReadDir(semplDir) //В папке с семплами ищем все семплы ogg
		if err != nil {
			fmt.Println("Ошибка просмотра исходной папки ", err)
		}
		for _, file := range files {
			if !file.IsDir() {
				if strings.HasSuffix(file.Name(), ".ogg") {
					sempls = append(sempls, file)
				}
			}
		}
		sempl := sempls[rand.Intn(len(sempls))] //выбираем из полученного списка семплов случайный rand()
		//Добавляем к видео звуковую дорожку из семпла
		cmd = exec.Command(filepath.Join(pathProgramm, "ffmpeg", "ffmpeg.exe"), "-i", filepath.Join(tmpDir, "tmp.mp4"), "-stream_loop", strconv.Itoa(numPdfFiles), "-i", filepath.Join(semplDir, sempl.Name()), "-c:v", "copy", "-c:a", "aac", "-map", "0:v:0", "-map", "1:a:0", "-shortest", outFileName)
		cmd.Stdout = &out
		err = cmd.Run()
		if err != nil {
			fmt.Println("Ошибка ", err)
		}
	} else { //если параметр addSempls установлен "no"
		copyFile(filepath.Join(tmpDir, "tmp.mp4"), outFileName) //то просто копируем темп файл в папку с видео
	}
	delFile(filepath.Join(tmpDir, "tmp.mp4")) //удаляем temp файл видео
	fmt.Println("Конвертировали видео", outFileName)
}
