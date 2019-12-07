package breaker

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/sirupsen/logrus"
)

type wrapperE1 struct {
	state string
}

func (w *wrapperE1) CommandFunc() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Executing ", w.state)
}
func (w *wrapperE1) DefaultFunc() {
	log.Formatter = new(logrus.JSONFormatter)
	log.WithFields(logrus.Fields{"Defaulted": w.state}).Info("task was defaulted")
	//fmt.Println("Defaulting command.....")
}
func (w *wrapperE1) CleanupFunc() {
	log.WithFields(logrus.Fields{"Cleaned": w.state}).Info("task was cleaned")
	//fmt.Println("Canceling command.....")
}
func (w *wrapperE1) timeout() time.Duration {
	return 200 * time.Millisecond
}
func (w *wrapperE1) Name() string {
	return "task1"
}

type wrapperE2 struct {
	state string
}

func (w *wrapperE2) CommandFunc() {
	time.Sleep(100 * time.Millisecond)
	log.Formatter = new(logrus.JSONFormatter)
	log.WithFields(logrus.Fields{"Executed": w.state}).Info("task was executed")
}
func (w *wrapperE2) DefaultFunc() {
	log.Formatter = new(logrus.JSONFormatter)
	log.WithFields(logrus.Fields{"Defaulted": w.state}).Info("task was defaulted")
}
func (w *wrapperE2) CleanupFunc() {
	log.Formatter = new(logrus.JSONFormatter)
	log.WithFields(logrus.Fields{"Cleaned": w.state}).Info("task was cleaned")
}
func (w *wrapperE2) timeout() time.Duration {
	return 50 * time.Millisecond
}
func (w *wrapperE2) Name() string {
	return "task1"
}

type wrapperE3 struct {
	state string
	done  bool
}

func (w *wrapperE3) CommandFunc() {
	for i := 1; i < 100; i++ {
		time.Sleep(1 * time.Millisecond)
		if w.done {
			return
		}
	}
	log.Formatter = new(logrus.JSONFormatter)
	log.WithFields(logrus.Fields{"Executed": w.state}).Info("task was executed")
}
func (w *wrapperE3) DefaultFunc() {
	fmt.Println("Calling done")
	w.done = true
}
func (w *wrapperE3) CleanupFunc() {
}

func randMillis(i int64) time.Duration {
	r := rand.Int63n(i)
	return time.Duration(r) * time.Millisecond
}
func (w *wrapperE3) Name() string {
	return "task1"
}
