package entity

func (this *Statistics) SetDate(date Date) {
	this.Date = date
}
func (this *Statistics) GetDate() (date Date) {
	return this.Date
}
func (this *Statistics) SetI(i int64) {
	this.I = i
}
func (this *Statistics) GetI() (i int64) {
	return this.I
}

type Statistics struct {
	Date Date
	I    int64
}
