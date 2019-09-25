package breaker

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

type wrapperE1 struct {
	name string
}

func (w *wrapperE1) CommandFunc() {
	var log = logrus.New()
	log.Formatter = new(logrus.JSONFormatter)
	time.Sleep(1 * time.Millisecond)
	log.WithFields(logrus.Fields{"Executed": w.name}).Info("task was executed")
	fmt.Println("Executing ", w.name)
}
func (w *wrapperE1) DefaultFunc() {
	log.Formatter = new(logrus.JSONFormatter)
	log.WithFields(logrus.Fields{"Defaulted": w.name}).Info("task was defaulted")
	//fmt.Println("Defaulting command.....")
}
func (w *wrapperE1) CleanupFunc() {
	log.WithFields(logrus.Fields{"Cleaned": w.name}).Info("task was cleaned")
	//fmt.Println("Canceling command.....")
}

type wrapperE2 struct {
	name string
	exec bool
}

func (w *wrapperE2) CommandFunc() {
	time.Sleep(1000 * time.Millisecond)
	//fmt.Println("Executing ", w.name)
	log.Formatter = new(logrus.JSONFormatter)
	log.WithFields(logrus.Fields{"Executed": w.name}).Info("task was executed")
	w.exec = true
}
func (w *wrapperE2) DefaultFunc() {
	//fmt.Println("Defaulting ", w.name)
	log.Formatter = new(logrus.JSONFormatter)
	log.WithFields(logrus.Fields{"Defaulted": w.name}).Info("task was defaulted")
}
func (w *wrapperE2) CleanupFunc() {
	log.Formatter = new(logrus.JSONFormatter)
	log.WithFields(logrus.Fields{"Cleaned": w.name}).Info("task was cleaned")
	//fmt.Println("Cleaning ", w.name)
}

type wrapperE3 struct {
}

func (w *wrapperE3) CommandFunc() {
}
func (w *wrapperE3) DefaultFunc() {
}
func (w *wrapperE3) CleanupFunc() {
}
func (w *wrapperE3) timeout() time.Duration {
	return time.Millisecond
}
