package LocalDateTime

import (
	"bitrade/libs/timex/LocalDate"
	"database/sql/driver"
	"errors"
	"time"
)

type LocalDateTime time.Time

const (
	Nanosecond  time.Duration = 1
	Microsecond               = 1000 * Nanosecond
	Millisecond               = 1000 * Microsecond
	Second                    = 1000 * Millisecond
	Minute                    = 60 * Second
	Hour                      = 60 * Minute
)
const (
	dateTimeFormart = "2006-01-02 15:04:05"
)

func (t *LocalDateTime) UnmarshalJSON(data []byte) (err error) {
	if data == nil || len(data) == 2 {
		return
	}
	now, err := time.ParseInLocation(`"`+dateTimeFormart+`"`, string(data), time.Local)
	*t = LocalDateTime(now)
	return
}

func (t LocalDateTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(dateTimeFormart)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, dateTimeFormart)
	b = append(b, '"')
	return b, nil
}

func (t LocalDateTime) Value() (driver.Value, error) {
	// MyTime 转换成 timex.Time 类型
	tTime := time.Time(t)
	return tTime.Format(dateTimeFormart), nil
}

func (t *LocalDateTime) Scan(v interface{}) error {
	switch vt := v.(type) {
	case string:
		// 字符串转成 timex.Time 类型
		tTime, _ := time.Parse(dateTimeFormart, vt)
		*t = LocalDateTime(tTime)
	case []byte:
		tTime, _ := time.Parse(dateTimeFormart, string(vt))
		*t = LocalDateTime(tTime)
	default:
		return errors.New("类型处理错误")
	}
	return nil
}
func (t LocalDateTime) ToLocalDate() *LocalDate.LocalDate {
	// MyTime 转换成 timex.Time 类型
	tTime := LocalDate.LocalDate(t)
	return &tTime
}

func (t *LocalDateTime) PlusDays(d int) (t2 *LocalDateTime) {
	// MyTime 转换成 timex.Time 类型
	time1 := time.Time(*t).Add(time.Duration(d) * time.Hour * 24)
	time2 := LocalDateTime(time1)
	return &time2
}
func (t *LocalDateTime) MinusDays(d int) (t2 *LocalDateTime) {
	// MyTime 转换成 timex.Time 类型
	time1 := time.Time(*t).Add(-time.Duration(d) * time.Hour * 24)
	time2 := LocalDateTime(time1)
	return &time2
}

func (t LocalDateTime) Format(fmt string) string {
	// MyTime 转换成 timex.Time 类型
	tTime := time.Time(t)
	return tTime.Format(fmt)
}

func Now() LocalDateTime {
	// MyTime 转换成 timex.Time 类型
	t := time.Now()
	return LocalDateTime(t)
}

func Of(ld *LocalDate.LocalDate, lt time.Time) *LocalDateTime {
	// MyTime 转换成 timex.Time 类型
	t := time.Date(time.Time(*ld).Year(), time.Time(*ld).Month(), time.Time(*ld).Day(), time.Time(lt).Hour(), time.Time(lt).Minute(), time.Time(lt).Second(), time.Time(lt).Nanosecond(), time.Local)
	ldnew := LocalDateTime(t)
	return &ldnew
}
