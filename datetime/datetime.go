// Package datetime 提供日期时间相关的功能，并提供了一个秒级精度的本地时间及相关变量，当仅需要秒级精度的时间时，可以直接从这里获取而不用调用time.Now()
// 该文件内提供了与日期时间相关的函数
package datetime

import (
	"fmt"
	"github.com/puresnr/go-cell/cast"
	"time"
)

const (
	SecOneDay  = 86400
)

const (
	InvalidAge = -1
)

const (
	TimeFormatDate   = "2006-01-02"
	TimeFormatDatetime = "2006-01-02 15:04:05"
	TimeFormatY_M = "2006_01"
)

// LawAge 返回周岁
// param: day - 待计算的日期，birthday: 生日，两者均是 TimeFormatDate 后的日期，比如"2011-11-11"
// return: 如果输入格式不正确，返回 InvalidAge, 否则返回周岁
func LawAge(day, birthday string) int {
	if len(day) != len(birthday) || len(birthday) != 10 {
		return InvalidAge
	}

	age := cast.Stoi(day[:4]) - cast.Stoi(birthday[:4])
	if day[5:] <= birthday[5:] {
		age -= 1
	}

	return age
}

// WeekFirstDate : 获取指定时间所在星期的星期一的日期, 形如 2006-01-01
func WeekFirstDate(t time.Time) string {
	offset := int(time.Monday - t.Weekday())
	if offset == 0 {
		return t.Format(TimeFormatDate)
	}
	if offset > 0 {
		return t.AddDate(0, 0, -6).Format(TimeFormatDate)
	}

	return t.AddDate(0, 0, offset).Format(TimeFormatDate)
}

// ToWeekdayISO 把 time.Weekday() 转换为 ISO weekday，星期日设为7，范围[1, 7]
func ToWeekdayISO(weekday time.Weekday) int {
	if weekday == time.Sunday {
		return 7
	}

	return int(weekday)
}

// FirstDayListOfMonth ：返回一个时间段内的所有月份第一天的列表(包括 start 和 end 所在月份)，形如2019-03-01
func FirstDayListOfMonth(sYear, sMonth, eYear, eMonth int) (ret []string) {
	sFirst, _ := time.ParseInLocation(TimeFormatDate, fmt.Sprintf("%4d-%02d-%02d", sYear, sMonth, 1), time.Local)

	eFirst, _ := time.ParseInLocation(TimeFormatDate, fmt.Sprintf("%4d-%02d-%02d", eYear, eMonth, 1), time.Local)

	for first := sFirst; !first.After(eFirst); first = first.AddDate(0, 1, 0) {
		ret = append(ret, first.Format(TimeFormatDate))
	}

	return
}

// SecPastOfDay ：获取指定的时间戳在其所在的那一天已经过去了多少秒 [0, 86399]
func SecPastOfDay(unixTime int64) int {
	return int((unixTime + int64(OffsetSec())) % SecOneDay)
}

// SecRestOfDay : 获取指定的时间戳在其所在的那一天还剩余多少秒 [1, 86400]
func SecRestOfDay(unixTime int64) int {
	return SecOneDay - SecPastOfDay(unixTime)
}

// DayOfUnixEpoch : 获取指定时间戳是自1970-1-1以来的第几天，当时间戳在1970-1-1之前会返回什么，尚未测试
func DayOfUnixEpoch(unixTime int64) int {
	return int((unixTime + int64(OffsetSec())) / SecOneDay)
}
