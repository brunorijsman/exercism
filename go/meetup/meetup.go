// Go exercise meetup.
package meetup

import "time"

// Different scheduling methods.
type WeekSchedule int

const (
	First  WeekSchedule = iota // First xxxday of the month.
	Second                     // Second xxxday of the month.
	Third                      // Third xxxday of the month.
	Fourth                     // Fourth xxxday of the moth.
	Last                       // Last xxxday of the month.
	Teenth                     //
)

// Compute the day-of-the-month for the given meeting schedule.
func Day(weekSchedule WeekSchedule, weekday time.Weekday, month time.Month, year int) int {
	switch weekSchedule {
	case First:
		return nthWeekDay(1, weekday, month, year)
	case Second:
		return nthWeekDay(2, weekday, month, year)
	case Third:
		return nthWeekDay(3, weekday, month, year)
	case Fourth:
		return nthWeekDay(4, weekday, month, year)
	case Last:
		return lastWeekDay(weekday, month, year)
	case Teenth:
		return teenthDay(weekday, month, year)
	}
	panic("Unreachable code")
}

// Determine day in range [13..19] in month of year that is the given weekday.
func teenthDay(weekday time.Weekday, month time.Month, year int) int {
	for day := 13; day <= 19; day++ {
		date := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
		if date.Weekday() == weekday {
			return day
		}
	}
	panic("Unreachable code")
}

// Determine nth (in range [1..4]) weekday of month in year.
func nthWeekDay(n int, weekday time.Weekday, month time.Month, year int) int {
	count := 0
	for day := 1; day <= 31; day++ {
		date := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
		if date.Weekday() == weekday {
			count++
			if count == n {
				return day
			}
		}
	}
	panic("Unreachable code")
}

// Determine last weekday of month in year.
func lastWeekDay(weekday time.Weekday, month time.Month, year int) int {
	for day := daysInMonth(month, year); day >= 1; day-- {
		date := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
		if date.Weekday() == weekday {
			return day
		}
	}
	panic("Unreachable code")
}

func daysInMonth(month time.Month, year int) int {
	switch month {
	case time.January, time.March, time.May, time.July, time.August, time.October, time.December:
		return 31
	case time.April, time.June, time.September, time.November:
		return 30
	case time.February:
		if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
			return 29
		}
		return 28
	default:
		panic("Unknown month")
	}

}
