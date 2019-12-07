package breaker

import (
	"time"

	"github.com/sirupsen/logrus"
)

type wrapper struct {
	state string
	exec  bool
}

func (w *wrapper) CommandFunc() {
	var log = logrus.New()
	log.Formatter = new(logrus.JSONFormatter)
	time.Sleep(1 * time.Millisecond)
	log.WithFields(logrus.Fields{"Executed": w.state}).Info("task was executed")
	//fmt.Println("Executing ", w.name)
	w.exec = true
}
func (w *wrapper) DefaultFunc() {
	log.Formatter = new(logrus.JSONFormatter)
	log.WithFields(logrus.Fields{"Defaulted": w.state}).Info("task was defaulted")
	//fmt.Println("Defaulting command.....")
}
func (w *wrapper) CleanupFunc() {
	log.WithFields(logrus.Fields{"Cleaned": w.state}).Info("task was cleaned")
	//fmt.Println("Canceling command.....")
}

func (w *wrapper) Name() string {
	return "task1"
}

type wrapper2 struct {
	name string
	exec bool
}

func (w *wrapper2) CommandFunc() {
	time.Sleep(1000 * time.Millisecond)
	//fmt.Println("Executing ", w.name)
	log.Formatter = new(logrus.JSONFormatter)
	log.WithFields(logrus.Fields{"Executed": w.name}).Info("task was executed")
	w.exec = true
}
func (w *wrapper2) DefaultFunc() {
	//fmt.Println("Defaulting ", w.name)
	log.Formatter = new(logrus.JSONFormatter)
	log.WithFields(logrus.Fields{"Defaulted": w.name}).Info("task was defaulted")
}
func (w *wrapper2) CleanupFunc() {
	log.Formatter = new(logrus.JSONFormatter)
	log.WithFields(logrus.Fields{"Cleaned": w.name}).Info("task was cleaned")
	//fmt.Println("Cleaning ", w.name)
}

func (w *wrapper2) Name() string {
	return "task1"
}

type wrapper3 struct {
}

func (w *wrapper3) CommandFunc() {
}
func (w *wrapper3) DefaultFunc() {
}
func (w *wrapper3) CleanupFunc() {
}
func (w *wrapper3) timeout() time.Duration {
	return time.Millisecond
}
func (w *wrapper3) Name() string {
	return "task1"
}
