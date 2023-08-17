// Package datetime 提供日期时间相关的功能
package datetime

import (
	"github.com/puresnr/go-cell/cast"
	"time"
)

const (
	SecOneDay = 86400
)

const (
	TimeFormatDate     = "2006-01-02"
	TimeFormatDatetime = "2006-01-02 15:04:05"
	TimeFormatY_M      = "2006_01"
)

const (
	InvalidAge = -1
)

// LawAge 返回周岁
// param: day: 待计算的日期，birthday: 生日. 两者格式要一致, 且符合 4位年+分隔符+2位月+分隔符+2位日 的格式, 比如 2011/01/11
// return: 如果输入格式不正确，返回 InvalidAge, 否则返回周岁
func LawAge(day, birthday string) int {
	if len(day) != len(birthday) || len(birthday) != 10 {
		return InvalidAge
	}

	age := cast.Stoi(day[:4]) - cast.Stoi(birthday[:4])
	if age <= 0 {
		return 0
	}

	if day[7] != birthday[7] {
		return InvalidAge
	}

	if day[5:] > birthday[5:] {
		return age
	}

	return age - 1
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

// todo: FirstUnixTsOfDay 返回指定时间戳所在日期的第一秒的时间戳
//func FirstUnixTsOfDay(unixTs uint32) uint32 {
//
//}

// todo: LastUnixTimeOfDay 返回指定时间戳所在日期的最后一秒的时间戳
//func LastUnixTimeOfDay(unixTs uint32) uint32 {
//	return FirstUnixTsOfDay(unixTs) + SecOneDay - 1
//}
