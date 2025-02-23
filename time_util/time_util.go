package time_util

import "time"

func PreMonthFirst(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month()-1, 1, 0, 0, 0, 0, t.Location())
}

func PreMonthLast(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 0, 0, 0, 0, 0, t.Location())
}

func MonthFirst(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

func MonthLast(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month()+1, 0, 0, 0, 0, 0, t.Location())
}

func NextMonthFirst(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month()+1, 1, 0, 0, 0, 0, t.Location())
}

func NextMonthLast(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month()+2, 0, 0, 0, 0, 0, t.Location())
}

func DateFirst(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func WeekFirst(t time.Time) time.Time {
	return t.AddDate(0, 0, -int(t.Weekday()))
}

func WeekLast(t time.Time) time.Time {
	return t.AddDate(0, 0, 7-int(t.Weekday()))
}
