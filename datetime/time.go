// Package datetime 提供日期时间相关的功能，并提供了一个秒级精度的本地时间及相关变量，当仅需要秒级精度的时间时，可以直接从这里获取而不用调用time.Now()
// 该文件维护了一个本地时间，每秒更新一次，同时更新相关的一些常用变量
package datetime

import (
	"sync"
	"time"
)

var (
	// 这两个值一旦设置就不会改变
	zone string
	offsetSec int

	// 每秒都变化的
	now time.Time
	unixTime int64
	datetimeNow string
	secondNow int

	// 每分钟变化的
	minuteNow int

	// 每小时变化的
	hourNow int

	// 每天变化的
	dayNow int
	dateNow string
	firstTimestampToday int64    // 今天的第一秒的时间戳
	lastTimestampToday int64    // 今天的最后一秒的时间戳
	dayOfUnixEpochNow int     // 获取今天是自1970-1-1以来的第几天
	weekday time.Weekday
	weekdayISO int    // ISO weekday，星期日设为7，范围[1, 7]

	// 每月变化的
	monthNow time.Month

	// 每年变化的
	yearNow int

	rwLock sync.RWMutex
)

func init() {
	initData()

	go func() {
		// 此时now不会发生变化，所以不加锁了
		time.Sleep(
			time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second() + 1, 0, time.Local).
			Sub(now))

		setData()

		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		for {
			select {
			case <- ticker.C:
				setData()
			}
		}
	}()
}

func initData() {
	// 初始化时不会并发

	zone, offsetSec = now.Zone()

	now := time.Now()
	unixTime = now.Unix()
	secondNow = now.Second()
	minuteNow = now.Minute()
	hourNow = now.Hour()
	yearNow, monthNow, dayNow = now.Date()
	dateNow = now.Format(TimeFormatDate)
	datetimeNow = now.Format(TimeFormatDatetime)

	firstTimestampToday = now.Unix() - (now.Unix() + int64(offsetSec)) % SecOneDay
	lastTimestampToday = firstTimestampToday + SecOneDay - 1
	dayOfUnixEpochNow = DayOfUnixEpoch(unixTime)

	weekday = now.Weekday()
	weekdayISO = ToWeekdayISO(weekday)
}

func setData() {
	rwLock.Lock()
	defer rwLock.Unlock()

	now = time.Now()
	unixTime = now.Unix()
	datetimeNow = now.Format(TimeFormatDatetime)

	oldSec := secondNow
	secondNow = now.Second()

	if secondNow >= oldSec {
		return
	}

	// minute 发生变化时
	oldMin := minuteNow
	minuteNow = (minuteNow + 1) % 60

	if minuteNow >= oldMin {
		return
	}

	// hour 发生变化时
	oldHour := hourNow
	hourNow = (hourNow + 1) % 24

	if hourNow >= oldHour {
		return
	}

	// day 发生变化时
	oldDay := dayNow
	dayNow = now.Day()
	dateNow = now.Format(TimeFormatDate)
	firstTimestampToday += SecOneDay
	lastTimestampToday += SecOneDay
	dayOfUnixEpochNow++
	weekday = (weekday + 1) % 7
	weekdayISO = weekdayISO + 1
	if weekdayISO == 8 {
		weekdayISO = 1
	}

	if dayNow >= oldDay {
		return
	}

	// month 发生变化时
	oldMonth := monthNow
	monthNow = monthNow + 1
	if monthNow == 13 {
		monthNow = 1
	}

	if monthNow >= oldMonth {
		return
	}

	// year 发生变化时
	yearNow++
}

func Zone() string {
	return zone
}

func OffsetSec() int {
	return offsetSec
}

func Now() time.Time {
	rwLock.RLock()
	defer rwLock.RUnlock()

	return now
}

func UnixTime() int64 {
	rwLock.RLock()
	defer rwLock.RUnlock()

	return unixTime
}

func HourNow() int {
	rwLock.RLock()
	defer rwLock.RUnlock()

	return hourNow
}

func MinuteNow() int {
	rwLock.RLock()
	defer rwLock.RUnlock()

	return minuteNow
}

func SecondNow() int {
	rwLock.RLock()
	defer rwLock.RUnlock()

	return secondNow
}

func DayNow() int {
	rwLock.RLock()
	defer rwLock.RUnlock()

	return dayNow
}

func MonthNow() time.Month {
	rwLock.RLock()
	defer rwLock.RUnlock()

	return monthNow
}

func YearNow() int {
	rwLock.RLock()
	defer rwLock.RUnlock()

	return yearNow
}

func DateNow() string {
	rwLock.RLock()
	defer rwLock.RUnlock()

	return dateNow
}

func DatetimeNow() string {
	rwLock.RLock()
	defer rwLock.RUnlock()

	return datetimeNow
}

func FirstTimeStampToday() int64 {
	rwLock.RLock()
	defer rwLock.RUnlock()

	return firstTimestampToday
}

func LastTimeStampToday() int64 {
	rwLock.RLock()
	defer rwLock.RUnlock()

	return lastTimestampToday
}

func DayOfUnixEpochNow() int {
	rwLock.RLock()
	defer rwLock.RUnlock()

	return dayOfUnixEpochNow
}

func Weekday() time.Weekday {
	rwLock.RLock()
	defer rwLock.RUnlock()

	return weekday
}

func WeekdayISO() int {
	rwLock.RLock()
	defer rwLock.RUnlock()

	return weekdayISO
}