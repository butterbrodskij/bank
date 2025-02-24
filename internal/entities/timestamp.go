package entities

import "fmt"

type timestamp struct {
	day  int
	hour int
	min  int
}

func newTimestamp() *timestamp {
	return &timestamp{
		day:  1,
		hour: 0,
		min:  0,
	}
}

func newExactTimestamp(day, hour, min int) *timestamp {
	return &timestamp{
		day:  day,
		hour: hour,
		min:  min,
	}
}

func (t *timestamp) normalize() *timestamp {
	if t == nil {
		return nil
	}
	if t.min >= 60 {
		t.hour += t.min / 60
		t.min %= 60
	}
	if t.hour >= 8 {
		t.day += t.hour / 8
		t.hour %= 8
	}
	return t
}

func (t *timestamp) isTheEndOfDay() bool {
	if t == nil {
		return false
	}
	weekDay := t.getWeekDay()
	switch weekDay {
	case "Saturday":
		return t.hour == 6 && t.min == 0
	case "Sunday":
		return true
	default:
		return t.hour == 8 && t.min == 0
	}
}

func (t *timestamp) isTheEndOfSimulation() bool {
	if t == nil {
		return false
	}
	return t.day >= 28
}

func (t *timestamp) nextDay() *timestamp {
	if t == nil {
		return newTimestamp()
	}
	return newExactTimestamp(t.day+1, 0, 0)
}

func (t *timestamp) toTheEndOfDay() *timestamp {
	if t == nil {
		newT := newTimestamp()
		newT.hour = 8
		return newT
	}
	if t.getWeekDay() == "Saturday" {
		t.hour = 6
		t.min = 0
		return t
	}
	if t.getWeekDay() == "Sunday" {
		t.hour = 0
		t.min = 0
		return t
	}
	t.hour = 8
	t.min = 0
	return t
}

func (t *timestamp) overTime() bool {
	if t == nil {
		return false
	}
	weekDay := t.getWeekDay()
	switch weekDay {
	case "Sunday":
		return true
	case "Saturday":
		return t.hour > 6 || t.hour == 6 && t.min > 0
	default:
		return false
	}
}

func (t *timestamp) difference(nt timestamp) int {
	return (t.day-nt.day)*24*60 + (t.hour-nt.hour)*60 + t.min - nt.min
}

func (t *timestamp) addValidMinutes(min int) (*timestamp, int) {
	if t == nil {
		return newTimestamp(), 0
	}
	oldT := *t
	t.min += min
	t.normalize()
	if oldT.day < t.day || t.overTime() {
		newT := oldT
		return newT.toTheEndOfDay(), newT.difference(oldT)
	}
	return t, min
}

func (t *timestamp) addMinutes(min int) *timestamp {
	if t == nil {
		return newTimestamp()
	}
	t.min += min
	return t.normalize()
}

func (t *timestamp) getTime() string {
	return fmt.Sprintf("%02d:%02d", t.hour+10, t.min)
}

func (t *timestamp) getWeekDay() string {
	switch t.day % 7 {
	case 0:
		return "Sunday"
	case 1:
		return "Monday"
	case 2:
		return "Tuesday"
	case 3:
		return "Wednesday"
	case 4:
		return "Thursday"
	case 5:
		return "Friday"
	case 6:
		return "Saturday"
	}
	return ""
}

func (t *timestamp) getDayInfo() string {
	return fmt.Sprintf("day %d: %s", t.day, t.getWeekDay())
}
