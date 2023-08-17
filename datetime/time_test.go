package datetime

import (
	"fmt"
	"os/exec"
	"time"
)

func settime(t time.Time) {
	hms := t.Format("15:04:05") // hour-minute-second
	err := exec.Command("cmd", "/c", "time", hms).Run()
	if err != nil {
		fmt.Println("set hour failed,err: \n", err)
	}
	time.Sleep(time.Second - time.Duration(time.Now().Nanosecond()) + 960*time.Millisecond)
}

func setdatetime(t time.Time) {
	hms := t.Format("15:04:05") // hour-minute-second
	err := exec.Command("cmd", "/c", "time", hms).Run()
	if err != nil {
		fmt.Println("set hour failed,err: \n", err)
	}
	date := t.Format("2006-01-02")
	err = exec.Command("cmd", "/c", "date", date).Run()
	if err != nil {
		fmt.Println("set day failed,err: \n", err)
	}
	time.Sleep(time.Second - time.Duration(time.Now().Nanosecond()) + 960*time.Millisecond)
}
