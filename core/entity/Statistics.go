package entity

import "github.com/qauzy/chocolate/xtime"

func (this *Statistics) SetDate(Date xtime.Xtime) (result *Statistics) {
	this.Date = Date
	return this
}
func (this *Statistics) GetDate() (Date xtime.Xtime) {
	return this.Date
}
func (this *Statistics) SetI(I int) (result *Statistics) {
	this.I = I
	return this
}
func (this *Statistics) GetI() (I int) {
	return this.I
}
func NewStatistics(date xtime.Xtime, i int) (ret *Statistics) {
	ret = new(Statistics)
	ret.Date = date
	ret.I = i
	return
}

type Statistics struct {
	Date xtime.Xtime `gorm:"column:date" json:"date"`
	I    int         `gorm:"column:i" json:"i"`
}
