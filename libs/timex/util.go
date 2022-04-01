package timex

import (
	"bitrade/libs/timex/LocalDate"
	"time"
)

func GetFirstDateOfWeek(current *LocalDate.LocalDate) *LocalDate.LocalDate {
	now := time.Time(*current)
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}

	weekStartDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	weekMonday := LocalDate.LocalDate(weekStartDate)
	return &weekMonday
}

func GetDayIndexOfWeek(current *LocalDate.LocalDate) int {
	now := time.Time(*current)
	w := now.Weekday()
	return int(w)
}
