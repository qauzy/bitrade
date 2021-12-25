package dateutils

import "time"

type SimpleDateFormat struct {
	sf string
}

func (this *SimpleDateFormat) Format(date time.Time) (result string) {
	return date.Format(this.sf)
}

func (this *SimpleDateFormat) Parse(date string) (result time.Time, err error) {
	return time.ParseInLocation(this.sf, date, time.Local)
}

func NewSimpleDateFormat(sf string) (result *SimpleDateFormat) {
	result = &SimpleDateFormat{sf: sf}
	return
}
