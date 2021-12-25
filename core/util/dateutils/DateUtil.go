package dateutils

import (
	"bitrade/libs/strutils"
	"errors"
	"github.com/shopspring/decimal"
	"net/http"
	"strconv"
	"time"
)

const (
	ERA           = 0
	YEAR          = 1
	MONTH         = 2
	WEEK_OF_YEAR  = 3
	WEEK_OF_MONTH = 4
	DATE          = 5
	DAY_OF_MONTH  = 5
	DAY_OF_YEAR   = 6
	DAY_OF_WEEK   = 7
)

var (
	YYYY_MM_DD_MM_HH_SS = NewSimpleDateFormat("2006-01-02 15:04:05")
	HHMMSS              = NewSimpleDateFormat("15:04:05")
	YYYY_MM_DD          = NewSimpleDateFormat("2006-01-02")
	YYYYMMDDMMHHSSSSS   = NewSimpleDateFormat("20060102150405.99999")
	YYYYMMDDHHMMSS      = NewSimpleDateFormat("20060102150405")
	YYYYMMDD            = NewSimpleDateFormat("20060102")
)

func DateToString(date time.Time) (result string) {
	return YYYY_MM_DD_MM_HH_SS.Format(date)
}
func DateToStringDate(date time.Time) (result string) {
	return YYYY_MM_DD.Format(date)
}
func ToYYYYMMDDMMHHSSSSS(date time.Time) (result string) {
	return YYYYMMDDMMHHSSSSS.Format(date)
}

/**
 * 开始时间 结束时间 是否合法  --> 判断是否开始时间小于今天并且开始时间小于结束时间
 *
 * @param startDate
 * @param endDate
 */
func ValidateDate(startDate time.Time, endDate time.Time) (err error) {
	var currentDate = GetCurrentDate()
	var compare = Compare(startDate, currentDate)
	var compare2 = Compare(startDate, endDate)
	if compare == -1 {
		err = errors.New("startDate cannot be less than currentDate")
		return
	}

	if compare2 == -1 {
		err = errors.New("startDate must be less than endDate")
		return
	}
	return
}

func ValidateEndDate(endDate time.Time) (err error) {
	var currentDate = GetCurrentDate()
	var compare int = Compare(currentDate, endDate)
	if compare == -1 {
		err = errors.New("currentDate must be less than endDate")
		return
	}
	return
}

/**
 * @param date1
 * @param date2
 * @return 1 大于 -1 小于 0 相等
 */
func Compare(date1 time.Time, date2 time.Time) (result int) {
	{
		if date1.UnixMilli() > date2.UnixMilli() {
			return 1
		} else {
			if date1.UnixMilli() < date2.UnixMilli() {
				return -1
			} else {
				return 0
			}
		}
	}
	return 0
}

/**
 * 获取当时日期时间串 格式 yyyy-MM-dd HH:mm:ss
 *
 * @return
 */
func GetDateTime() (result string) {
	return YYYY_MM_DD_MM_HH_SS.Format(time.Now())
}

func StringToDate(dateString string) (result time.Time, err error) {
	return YYYY_MM_DD_MM_HH_SS.Parse(dateString)
}

/**
 * 获取当时日期串 格式 yyyy-MM-dd
 *
 * @return
 */
func GetDate() (result string) {
	return YYYY_MM_DD.Format(time.Now())
}

/**
 * 获取当时日期串 格式 yyyy-MM-dd
 *
 * @return
 */
func GetNowDateYMD() (result string) {
	return YYYYMMDD.Format(time.Now())
}

/**
 * 获取当时日期串 格式 yyyy-MM-dd
 *
 * @return
 */
func GetDateYMD(date time.Time) (result string) {
	return YYYYMMDD.Format(date)
}

func StrToDate(dateString string) (result time.Time, err error) {

	return YYYY_MM_DD_MM_HH_SS.Parse(dateString)
}

func StrToYYMMDDDate(dateString string) (result time.Time, err error) {

	return YYYY_MM_DD.Parse(dateString)
}

func DiffDays(startDate time.Time, endDate time.Time) (result int64) {
	var days int64 = 0
	var start int64 = startDate.UnixMilli()
	var end int64 = endDate.UnixMilli()
	days = (end - start) / 86400000
	return days
}

/**
 * 获取当时日期串 格式 yyyy-MM-dd
 *
 * @return
 */func DateAddMonth(date time.Time, month int) (result time.Time) {
	return Add(date, MONTH, month)
}

/**
 * 获取当时日期串 格式 yyyy-MM-dd
 *
 * @return
 */
func DateAddDay(date time.Time, day int) (result time.Time) {
	return Add(date, DAY_OF_YEAR, day)
}

/**
 * 获取当时日期串 格式 yyyy-MM-dd
 *
 * @return
 */
func DateAddYear(date time.Time, year int) (result time.Time) {
	return Add(date, YEAR, year)
}

func DateStrAddDay(dateString string, day int) (result string) {
	var date, _ = StrToYYMMDDDate(dateString)
	var calendar = date.AddDate(0, 0, day)
	return YYYY_MM_DD.Format(calendar)
}

func DateNowAddDay(day int) (result string) {
	var now = time.Now()
	var calendar = now.AddDate(0, 0, day)
	return YYYY_MM_DD.Format(calendar)
}

func DateNowAddMonth(month int) (result string) {
	var now = time.Now()
	var calendar = now.AddDate(0, month, 0)
	return YYYY_MM_DD.Format(calendar)
}

func RemainDateToString(startDate time.Time, endDate time.Time) (result string) {
	var buff = strutils.NewStringBuilder()
	var times int64 = endDate.UnixMilli() - startDate.UnixMilli()
	if times < -1 {
		buff.Append("过期")
	} else {
		var temp int64 = 86400000
		var d = times / temp
		times %= temp
		temp /= 24
		var m = times / temp
		times %= temp
		temp /= 60
		var s = times / temp
		buff.Append(strconv.FormatInt(d, 10))
		buff.Append("天")
		buff.Append(strconv.FormatInt(m, 10))
		buff.Append("小时")
		buff.Append(strconv.FormatInt(s, 10))
		buff.Append("分")
	}
	return buff.ToString()
}
func Add(date time.Time, oType int, value int) (result time.Time) {
	switch oType {
	case YEAR:
		return date.AddDate(value, 0, 0)
	case MONTH:
		return date.AddDate(0, value, 0)
	case DAY_OF_YEAR:
		return date.AddDate(0, 0, value)
	}

	return
}

func GetLinkUrl(flag bool, content string, id string) (result string) {
	if flag {
		content = "<a href='finance.do?id=" + id + "'>" + content + "</a>"
	}
	return content
}

func GetTimeStrCur(format string, date string) (result time.Time, err error) {
	var sf = NewSimpleDateFormat(format)
	return sf.Parse(date)
}

func GetCurrentDate() (result time.Time) {
	return time.Now()
}

func GetFormatTime(format *SimpleDateFormat, date time.Time) (result string) {
	return format.Format(date)
}

/**
 * 获取时间戳
 *
 * @return
 */
func GetTimeMillis() (result int64) {
	return time.Now().UnixMilli()
}

/**
 * 获取时间戳
 *
 * @return
 */
func GetWeekDay(date time.Time) (result string) {

	var dayOfWeek = date.Weekday().String()
	switch dayOfWeek {
	case "Monday":
		return "周日"
	case "Tuesday":
		return "周一"
	case "Wednesday":
		return "周二"
	case "Thursday":
		return "周三"
	case "Friday":
		return "周四"
	case "Saturday":
		return "周五"
	case "Sunday":
		return "周六"
	}
	return ""
}

func ToGMTString(date time.Time) (result string) {
	var df = NewSimpleDateFormat(http.TimeFormat)
	return df.Format(date)
}

/**
 * 得到当前时间与某个时间的差的分钟数
 *
 * @param date
 * @return
 */
func DiffMinute(date time.Time) (result decimal.Decimal) {
	f := decimal.NewFromInt(time.Now().UnixMilli() - date.UnixMilli())
	f2 := decimal.NewFromInt(60000)
	f.Div(f2)
	return f.Div(f2)
}

/**
 * 获取过去第几天的日期
 *
 * @param past
 * @return
 */
func GetPastDate(past int) (result string) {
	var now = time.Now()
	var calendar = time.Date(now.Year(), now.Month(), now.Day()-past, 0, 0, 0, 0, time.Local)

	return YYYY_MM_DD.Format(calendar)
}

/**
 * 获取未来 第 past 天的日期
 *
 * @param past
 * @return
 */
func GetFetureDate(past int) (result string) {
	var now = time.Now()
	var calendar = time.Date(now.Year(), now.Month(), now.Day()+past, 0, 0, 0, 0, time.Local)

	return YYYY_MM_DD.Format(calendar)
}

/**
 * 获取未来 第 past 天的日期
 *
 * @param past
 * @return
 */
func GetDatePart(date time.Time, part int) (result time.Time) {
	var calendar = time.Date(date.Year(), date.Month(), date.Day()+part, 0, 0, 0, 0, time.Local)
	return calendar
}

/**
 * 获取当前时间到凌晨多少s
 *
 * @return
 */
func CalculateCurrentTime2SecondDaySec() (result int64) {
	var now = time.Now()
	var currentTime = now.UnixMilli()
	var calendar = time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, time.Local)
	var endTime int64 = calendar.UnixMilli()
	return (endTime - currentTime) / 1000
}

type DateUtil struct {
}
