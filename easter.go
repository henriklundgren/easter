package easter

import (
  "fmt"
  "time"
  "errors"
  "github.com/goodsign/monday"
)

const (
  error_int string = "The given year fall below the established international calendar."
  error_national string = "Given year not applicable on region."
  layout string = "2006-1-2"
)

type Easter struct {
  Date time.Time
}

// Set/Initialize
// Use convinience methods below to return Easter dates in time.Time format.
func Set(year int, locale string, location string) (Easter, error) {
  // International exception[@see doc ref:1]
  if year < 1753 {
    return Easter{time.Now()}, errors.New(error_int)
  }
  // National exceptions
  if locale == "sv_SE" && year < 1844 {
    return Easter{time.Now()}, errors.New(error_national)
  }

  // Calculate Paschal
  easterday := paschalAlgorithm(year)

  // Make timestring
  timeString := fmt.Sprintf("%[1]d-%[2]d-%[3]d",
    easterday["year"],
    easterday["month"],
    easterday["day"],
  )

  // Return time object
  myEaster := Easter{}
  if loc, locErr := time.LoadLocation(location); locErr == nil {
    lang, _ := getLocale(locale)
    parsed, _ := monday.ParseInLocation(layout, timeString, loc, lang)
    myEaster = Easter{parsed}
  } else {
    parsed, _ := time.Parse(layout, timeString)
    myEaster = Easter{parsed}
  }

  return myEaster, nil
}

func (d Easter) Day() time.Time {
  return d.Date
}

func (d Easter) Septuagesima() time.Time {
  return d.Date.AddDate(0, 0, -64)
}

func (d Easter) Sexagesima() time.Time {
  return d.Date.AddDate(0, 0, -56)
}

// Fastlagssöndag
func (d Easter) Quinquagesima() time.Time {
  return d.Date.AddDate(0, 0, -50)
}

// Fettisdagen
func (d Easter) ShroveTuesday() time.Time {
  return d.Date.AddDate(0, 0, -47)
}

func (d Easter) AshWednesday() time.Time {
  return d.Date.AddDate(0, 0, -46)
}

// 1st Sunday of Lent
// Maybe only Lutherans that name the Sundays?
func (d Easter) Invocabit() time.Time {
  return d.Date.AddDate(0, 0, -42)
}

// 2nd Sunday of Lent
func (d Easter) Reminiscere() time.Time {
  return d.Date.AddDate(0, 0, -35)
}

// 3rd Sunday of Lent
func (d Easter) Oculi() time.Time {
  return d.Date.AddDate(0, 0, -28)
}

// 4th Sunday of Lent
// Midfastosöndagen
func (d Easter) Laetare() time.Time {
  return d.Date.AddDate(0, 0, -21)
}

// 5th Sunday of Lent
func (d Easter) Judica() time.Time {
  return d.Date.AddDate(0, 0, -16)
}

// 6th Sunday of Lent
func (d Easter) PalmSunday() time.Time {
  return d.Date.AddDate(0, 0, -9)
}

// Långfredag
func (d Easter) GoodFriday() time.Time {
  return d.Date.AddDate(0, 0, -2)
}

func (d Easter) Monday() time.Time {
  return d.Date.AddDate(0, 0, 1)
}

func (d Easter) HolyThursday() time.Time {
  return d.Date.AddDate(0, 0, 39)
}

func (d Easter) Pentecost() time.Time {
  return d.Date.AddDate(0, 0, 49)
}
