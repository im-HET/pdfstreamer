package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

func streamVideo(wg *sync.WaitGroup) {
	for {
		files, err := ioutil.ReadDir(videoDir)
		if err != nil {
			fmt.Println("Ошибка просмотра исходной папки ", err)
		}
		for _, file := range files {
			if !file.IsDir() {
				if strings.HasSuffix(file.Name(), ".mp4") {
					//ffmpeg.exe -re -r 24 -i c:\ffmpeg\temp\%FileName% -f mpegts udp://192.168.0.14:1001?pkt_size=1316
					//cmd := exec.Command("ffmpeg", "-re", "-r", "24", "-i", videoDir+file.Name(), "-f", "mpegts", "udp://235.10.10.1:1234?pkt_size=1316")
					if IsExistFileWithoutEx(pdfDir, file.Name()) {
						fmt.Println("Стримим файл ", filepath.Join(videoDir, file.Name()))
						//cmd := exec.Command("ffmpeg", "-re", "-i", videoDir+file.Name(), "-c:v", "copy", "-f", "rtp_mpegts", "rtp://235.10.10.1:1001?pkt_size=1316")
						cmd := exec.Command(filepath.Join(pathProgramm, "vlc", "vlc.exe"), "-I", "dummy", "-vvv", filepath.Join(videoDir, file.Name()), "--sout=#rtp{dst="+dstip+",port="+dstport+",mux=ts,sap,name=TV}", "--no-sout-all", "--sout-keep", "vlc://quit")
						var out bytes.Buffer
						cmd.Stdout = &out
						err = cmd.Run()
						if err != nil {
							fmt.Println("Ошибка стриминга", err)
						}
					} else {
						delFile(filepath.Join(videoDir, file.Name()))
					}
				}
			}
		}
	}
	wg.Done()
}

func IsExistFileWithoutEx(path string, filename string) bool {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println("Ошибка просмотра исходной папки ", err)
	}
	for _, file := range files {
		if !file.IsDir() {
			if file.Name()[:len(file.Name())-4] == filename[:len(filename)-4] {
				return true
			}
		}
	}
	return false
}
