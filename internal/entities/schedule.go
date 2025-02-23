package entities

import "fmt"

type Schedule struct {
	LunchDuration int
	lunchStart    *timestamp
}

func NewSchedule(lunchDuration int) *Schedule {
	return &Schedule{
		LunchDuration: lunchDuration,
		lunchStart:    newExactTimestamp(0, 4, 0),
	}
}

func (s *Schedule) Update(lunchDuration int) (*Schedule, error) {
	if s == nil {
		return nil, fmt.Errorf("nil Schedule")
	}
	s.LunchDuration = lunchDuration
	return s, nil
}

func (s *Schedule) GetLunchTime(t *timestamp) string {
	if s.LunchDuration == 0 || t == nil || t.getWeekDay() == "Sunday" {
		return "no lunch"
	}
	return fmt.Sprintf("%s-%s", s.lunchStart.getTime(), s.lunchStart.addMinutes(s.LunchDuration).getTime())
}

func (s *Schedule) GetWorkTime(t *timestamp) string {
	if t == nil || t.getWeekDay() == "Sunday" {
		return "no work"
	}
	if t.getWeekDay() == "Saturday" {
		return "10:00-16:00"
	}
	return "10:00-18:00"
}
