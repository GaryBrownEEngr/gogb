package gogb

import (
	"errors"
	"time"
)

func Time(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) time.Time {
	t := time.Date(year, month, day, hour, min, sec, nsec, loc)
	return t
}

func TimePtr(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) *time.Time {
	t := Time(year, month, day, hour, min, sec, nsec, loc)
	return &t
}

func DateUtc(year int, month time.Month, day int) time.Time {
	t := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	return t
}

func DateUtcPtr(year int, month time.Month, day int) *time.Time {
	t := DateUtc(year, month, day)
	return &t
}

func GetAge(dateOfBirth, compareDate time.Time) (int, error) {
	// Set both times to UTC
	dateOfBirth = dateOfBirth.UTC()
	compareDate = compareDate.UTC()

	// Only use year, month, day
	aYear, aMonth, aDay := dateOfBirth.Date()
	dateOfBirth = time.Date(aYear, aMonth, aDay, 0, 0, 0, 0, time.UTC)
	bYear, bMonth, bDay := compareDate.Date()
	compareDate = time.Date(bYear, bMonth, bDay, 0, 0, 0, 0, time.UTC)

	if compareDate.Before(dateOfBirth) {
		return 0, errors.New("invalid, negative age")
	}

	age := bYear - aYear

	// Check if a full year hasn't passed yet for the last year
	anniversary := dateOfBirth.AddDate(age, 0, 0)
	if anniversary.After(compareDate) {
		age--
	}

	return age, nil
}
