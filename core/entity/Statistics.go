package entity

import "time"

func (this *Statistics) SetDate(date time.Time) (result *Statistics) {
	this.Date = date
	return this
}
func (this *Statistics) GetDate() (date time.Time) {
	return this.Date
}
func (this *Statistics) SetI(i int64) (result *Statistics) {
	this.I = i
	return this
}
func (this *Statistics) GetI() (i int64) {
	return this.I
}

type Statistics struct {
	Date time.Time
	I    int64
}
