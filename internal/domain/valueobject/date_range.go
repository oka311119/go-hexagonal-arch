package valueobject

import (
	"errors"
	"time"
)

type DateRange struct {
	Date1, Date2 time.Time
}

func NewDateRange(date1, date2 time.Time) (*DateRange, error) {
	if date1.After(date2) {
		return nil, errors.New("date1 cannot be later than date2")
	}
	return &DateRange{Date1: date1, Date2: date2}, nil
}

func (d *DateRange) GetDate1() time.Time {
	return d.Date1
}

func (d *DateRange) GetDate2() time.Time {
	return d.Date2
}

func (d *DateRange) SetDate1(date time.Time) (*DateRange, error) {
	return NewDateRange(date, d.Date1)
}

func (d *DateRange) SetDate2(date time.Time) (*DateRange, error) {
	return NewDateRange(d.Date2, date)
}
