package glog

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

var gLog *Log

type Log struct {
	filename string

	curfilename string
	curtm       time.Time
	curfile     *os.File
	logger      *log.Logger

	mutex sync.RWMutex
}

func Create(filename string) bool {

	if gLog != nil {
		return false
	}

	gLog = &Log{
		filename: filename,
	}

	if !gLog.init() {
		return false
	}

	return true
}

func (this *Log) init() bool {

	this.curtm = time.Now()
	this.curfilename = this.GetTmFileName()

	file := this.OpenFile(this.curfilename)
	if file == nil {
		return false
	}
	this.curfile = file
	this.logger = log.New(file, "", log.Lshortfile|log.LstdFlags)
	this.CreateLink(this.filename, this.curfilename)

	go gLog.TimeAction()

	return true
}

func (this *Log) OpenFile(filename string) *os.File {

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("log open file error ", err)
		return nil
	}

	return file
}

// create link
// like server.log -> server.log_20200707-10
func (this *Log) CreateLink(source string, target string) {

	linkname, _ := os.Readlink(source)
	if len(linkname) != 0 {

		if linkname == target {
			return
		} else if linkname != target {
			// remove link
			os.Remove(source)
		}
	}

	// create link
	os.Symlink(target, source)
}

func (this *Log) GetTmFileName() string {

	tm := this.curtm
	y := uint32(tm.Year())
	m := uint32(tm.Month())
	d := uint32(tm.Day())
	h := uint32(tm.Hour())
	mm := uint32(tm.Minute())
	newname := fmt.Sprintf("%s_%d%02d%02d-%02d%02d", this.filename, y, m, d, h, mm)
	return newname
}

func (this *Log) TimeAction() {

	for {
		time.Sleep(1 * time.Second)
		now := time.Now()
		last := this.curtm

		if now.Year() == last.Year() &&
			now.Month() == last.Month() &&
			now.Day() == last.Day() &&
			now.Hour() == last.Hour() &&
			now.Minute() == last.Minute() {
			continue
		}

		// file swap
		this.curtm = now
		this.curfilename = this.GetTmFileName()

		file := this.OpenFile(this.curfilename)
		if file == nil {
			continue
		}

		if this.curfile != nil {
			this.curfile.Close()
		}

		this.curfile = file
		this.logger.SetOutput(file)

		this.CreateLink(this.filename, this.curfilename)
	}
}

func Info(args ...interface{}) {

	gLog.logger.Println(args...)
}
