package datetime

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os/exec"
	"testing"
	"time"
)

func settime(t time.Time) {
	hms := t.Format("15:04:05") // hour-minute-second
	err := exec.Command("cmd", "/c", "time", hms).Run()
	if err != nil {
		fmt.Println("set hour failed,err: \n", err)
	}
	time.Sleep(time.Second - time.Duration(time.Now().Nanosecond()) + 960 * time.Millisecond)
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
	time.Sleep(time.Second - time.Duration(time.Now().Nanosecond()) + 960 * time.Millisecond)
}

func TestSetData(t *testing.T) {
	// 测试每秒变化和分钟变化
	settime(time.Date(2021, 9, 7, 22, 49, 55, 0, time.Local))

	ticker := time.NewTicker(time.Second)
	cnt := 0
	for {
		select {
		case <- ticker.C:
			tnow := time.Now()
			assert.Equal(t, UnixNow(), tnow.Unix())
			assert.Equal(t, DatetimeNow(), tnow.Format(TimeFormatDatetime))
			assert.Equal(t, SecondNow(), tnow.Second())
			assert.Equal(t, MinuteNow(), tnow.Minute())
			fmt.Println(tnow.Format(TimeFormatDatetime))
		}

		if cnt > 8 {
			ticker.Stop()
			break
		}
		cnt++
	}

	// 测试小时变化
	settime(time.Date(2021, 9, 7, 22, 59, 55, 0, time.Local))

	ticker = time.NewTicker(time.Second)
	cnt = 0
	for {
		select {
		case <- ticker.C:
			tnow := time.Now()
			assert.Equal(t, UnixNow(), tnow.Unix())
			assert.Equal(t, DatetimeNow(), tnow.Format(TimeFormatDatetime))
			assert.Equal(t, SecondNow(), tnow.Second())
			assert.Equal(t, MinuteNow(), tnow.Minute())
			assert.Equal(t, HourNow(), tnow.Hour())
			fmt.Println(tnow.Format(TimeFormatDatetime))
		}

		if cnt > 8 {
			ticker.Stop()
			break
		}
		cnt++
	}

	// 测试天变化
	setdatetime(time.Date(2021, 9, 7, 23, 59, 55, 0, time.Local))
	ticker = time.NewTicker(time.Second)
	cnt = 0
	for {
		select {
		case <- ticker.C:
			tnow := time.Now()
			assert.Equal(t, UnixNow(), tnow.Unix())
			assert.Equal(t, DatetimeNow(), tnow.Format(TimeFormatDatetime))
			assert.Equal(t, SecondNow(), tnow.Second())
			assert.Equal(t, MinuteNow(), tnow.Minute())
			assert.Equal(t, HourNow(), tnow.Hour())
			assert.Equal(t, DayNow(), tnow.Day())
			assert.Equal(t, DayOfUnixEpochNow(), DayOfUnixEpoch(tnow.Unix()))
			assert.Equal(t, Weekday(), tnow.Weekday())
			assert.Equal(t, WeekdayISO(), ToWeekdayISO(tnow.Weekday()))
			fmt.Println(tnow.Format(TimeFormatDatetime))
		}

		if cnt > 8 {
			ticker.Stop()
			break
		}
		cnt++
	}

	// 测试月变化
	setdatetime(time.Date(2021, 9, 30, 23, 59, 55, 0, time.Local))
	ticker = time.NewTicker(time.Second)
	cnt = 0
	for {
		select {
		case <- ticker.C:
			tnow := time.Now()
			assert.Equal(t, UnixNow(), tnow.Unix())
			assert.Equal(t, DatetimeNow(), tnow.Format(TimeFormatDatetime))
			assert.Equal(t, SecondNow(), tnow.Second())
			assert.Equal(t, MinuteNow(), tnow.Minute())
			assert.Equal(t, HourNow(), tnow.Hour())
			assert.Equal(t, DayNow(), tnow.Day())
			assert.Equal(t, DayOfUnixEpochNow(), DayOfUnixEpoch(tnow.Unix()))
			assert.Equal(t, Weekday(), tnow.Weekday())
			assert.Equal(t, WeekdayISO(), ToWeekdayISO(tnow.Weekday()))
			assert.Equal(t, MonthNow(), tnow.Month())
			fmt.Println(tnow.Format(TimeFormatDatetime))
		}

		if cnt > 8 {
			ticker.Stop()
			break
		}
		cnt++
	}

	// 测试年变化
	setdatetime(time.Date(2021, 12, 31, 23, 59, 55, 0, time.Local))
	ticker = time.NewTicker(time.Second)
	cnt = 0
	for {
		select {
		case <- ticker.C:
			tnow := time.Now()
			assert.Equal(t, UnixNow(), tnow.Unix())
			assert.Equal(t, DatetimeNow(), tnow.Format(TimeFormatDatetime))
			assert.Equal(t, SecondNow(), tnow.Second())
			assert.Equal(t, MinuteNow(), tnow.Minute())
			assert.Equal(t, HourNow(), tnow.Hour())
			assert.Equal(t, DayNow(), tnow.Day())
			assert.Equal(t, DayOfUnixEpochNow(), DayOfUnixEpoch(tnow.Unix()))
			assert.Equal(t, Weekday(), tnow.Weekday())
			assert.Equal(t, WeekdayISO(), ToWeekdayISO(tnow.Weekday()))
			assert.Equal(t, MonthNow(), tnow.Month())
			assert.Equal(t, YearNow(), tnow.Year())
			fmt.Println(tnow.Format(TimeFormatDatetime))
		}

		if cnt > 8 {
			ticker.Stop()
			break
		}
		cnt++
	}
}
