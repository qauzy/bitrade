package LocalDate

import (
	"database/sql/driver"
	"errors"
	"time"
)

type LocalDate time.Time

const (
	dateFormart      = "2006-01-02"
	dateFormartShort = "20060102"
	DTF_YMD          = "2006-01-02"
	DTF_YM           = "2006_01"
)

func Now() *LocalDate {
	l := LocalDate(time.Now())
	return &l
}
func (t *LocalDate) UnmarshalJSON(data []byte) (err error) {
	if data == nil || len(data) == 2 {
		return
	}

	now, err := time.ParseInLocation(`"`+dateFormart+`"`, string(data), time.Local)
	if err != nil {
		now, err = time.ParseInLocation(`"`+dateFormartShort+`"`, string(data), time.Local)
	}
	*t = LocalDate(now)
	return
}

func (t LocalDate) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(dateFormart)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, dateFormart)
	b = append(b, '"')
	return b, nil
}

func (t LocalDate) Value() (driver.Value, error) {
	// MyTime 转换成 timex.Time 类型
	tTime := time.Time(t)
	return tTime.Format(dateFormart), nil
}

func (t *LocalDate) Scan(v interface{}) error {
	switch vt := v.(type) {
	case string:
		// 字符串转成 timex.Time 类型
		tTime, err := time.Parse(dateFormart, vt)
		if err != nil {
			tTime, _ = time.Parse(dateFormartShort, vt)
		}
		*t = LocalDate(tTime)
	case []byte:
		tTime, err := time.Parse(dateFormart, string(vt))
		if err != nil {
			tTime, _ = time.Parse(dateFormartShort, string(vt))
		}
		*t = LocalDate(tTime)
	default:
		return errors.New("类型处理错误")
	}
	return nil
}

func (t *LocalDate) IsBefore(t2 *LocalDate) bool {
	// MyTime 转换成 timex.Time 类型
	tTime := time.Time(*t)
	tTime2 := time.Time(*t2)
	return tTime.Before(tTime2)
}

func (t *LocalDate) IsAfter(t2 *LocalDate) bool {
	// MyTime 转换成 timex.Time 类型
	tTime := time.Time(*t)
	tTime2 := time.Time(*t2)
	return tTime.After(tTime2)
}

func (t *LocalDate) IsEqual(t2 *LocalDate) bool {
	// MyTime 转换成 timex.Time 类型
	tTime := time.Time(*t)
	tTime2 := time.Time(*t2)
	return tTime.Equal(tTime2)
}

func (t *LocalDate) MinusDays(d int) (t2 *LocalDate) {
	// MyTime 转换成 timex.Time 类型
	time1 := time.Time(*t).Add(-time.Duration(d) * time.Hour * 24)
	time2 := LocalDate(time1)
	return &time2
}
func (t *LocalDate) MinusMonths(m int) (t2 *LocalDate) {
	// MyTime 转换成 timex.Time 类型
	t0 := time.Time(*t)
	time1 := time.Date(t0.Year(), t0.Month()-1, t0.Day(), t0.Hour(), t0.Minute(), t0.Second(), t0.Nanosecond(), t0.Location())
	time2 := LocalDate(time1)
	return &time2
}

func (t *LocalDate) PlusDays(d int) (t2 *LocalDate) {
	// MyTime 转换成 timex.Time 类型
	time1 := time.Time(*t).Add(time.Duration(d) * time.Hour * 24)
	time2 := LocalDate(time1)
	return &time2
}
func (t *LocalDate) ToEpochDay() (days int) {
	// MyTime 转换成 timex.Time 类型
	d := time.Time(*t).Unix() / 86400

	return int(d)
}

func (t LocalDate) Format(fmt string) string {
	// MyTime 转换成 timex.Time 类型
	tTime := time.Time(t)
	return tTime.Format(fmt)
}
func (t LocalDate) String() string {
	// MyTime 转换成 timex.Time 类型
	return t.Format(dateFormart)
}
func Parse(str string, f string) *LocalDate {
	t, _ := time.ParseInLocation(f, str, time.Local)
	lt := LocalDate(t)
	return &lt
}

func ConvertStringToLocalDate(src string, f string) *LocalDate {
	t, _ := time.ParseInLocation(f, src, time.Local)
	lt := LocalDate(t)
	return &lt
}
