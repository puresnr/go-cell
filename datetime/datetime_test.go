package datetime

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestLawAge(t *testing.T) {
	assert.Equal(t, LawAge("2021-09-09","2020-01-01"), 1)
	assert.Equal(t, LawAge("2021-09-09","2020-09-09"), 0)
	assert.Equal(t, LawAge("2021-09-09","2020-09-10"), 0)
	assert.Equal(t, LawAge("2021-09-09","2021-09-08"), 0)
	assert.Equal(t, LawAge("2021-09-09","2021-09-09"), 0)
	assert.Equal(t, LawAge("2021-09-09","2021-09-10"), 0)
	assert.Equal(t, LawAge("2021-09-09","2021-1-01"), InvalidAge)
	assert.Equal(t, LawAge("2021-09-09",time.Now().Format(TimeFormatDatetime)), InvalidAge)
	assert.Equal(t, LawAge("2021-09-9","2021-1-01"), InvalidAge)
	assert.Equal(t, LawAge("2021-09-09 15:00:00",time.Now().Format(TimeFormatDatetime)), InvalidAge)
	assert.Equal(t, LawAge("2021-09-09","2022-09-08"), 0)
}

func TestToWeekdayISO(t *testing.T) {
	weekdays := []time.Weekday{time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday, time.Saturday, time.Sunday}
	for idx, w := range weekdays {
		assert.Equal(t, idx + 1, ToWeekdayISO(w))
	}
}