package main

import (
	"sync"
)

var pathProgramm = ""
var sourceDir = ".//"
var pdfDir = ".//pdf//"
var videoDir = ".//video//"
var tmpDir = ".//tmp//"
var showtime = "6"
var semplDir = ".//sempls//"
var addSempls = "no"
var dstip = "235.10.10.1"
var dstport = "1001"

func main() {
	loadConf()

	var wg sync.WaitGroup

	wg.Add(1)
	go compareDirPDF(&wg)

	wg.Add(1)
	go generateVideo(&wg)

	wg.Add(1)
	go streamVideo(&wg)

	wg.Wait()
}
