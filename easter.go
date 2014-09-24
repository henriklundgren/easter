package easter

import (
  "fmt"
  "time"
  "errors"
)

const (
  error_int string = "The given year fall below the established international calendar."
  layout string = "2006-1-2"
)

// Exported const
const (
  Septuagesima  = -64
  Sexagesima    = -56
  Quinquagesima = -50
  ShroveTuesday = -47
  AshWednesday  = -46
  Invocabit     = -42
  ShroveMonday  = -41
  Reminiscere   = -35
  Oculi	        = -28
  Laetare       = -21
  Judica        = -16
  PalmSunday    = -9
  GoodFriday    = -2
  Day           = 0
  Monday        = +1
  HolyThursday  = +39
  Pentecost     = +49
  WhitMonday    = +50
  CorpusChristi = +60
)

var nameMap = map[int]string{
  -64: "Septuagesima",
  -56: "Sexagesima",
  -50: "Quinquagesima",
  -47: "ShroveTuesday",
  -46: "AshWednesday",
  -42: "Invocabit",
  -41: "ShroveMonday",
  -35: "Reminiscere",
  -28: "Oculi",
  -21: "Laetare",
  -16: "Judica",
   -9: "PalmSunday",
   -2: "GoodFriday",
    0: "Day",
   +1: "Monday",
  +39: "HolyThursday",
  +49: "Pentecost",
  +50: "WhitMonday",
  +60: "CorpusChristi",
}

type Easter struct {
  Date time.Time
  Name string
}

type EasterCollection []Easter

// Set/Initialize
func Set(year int, location *time.Location) (Easter, error) {
  if year < 1753 {
    return Easter{time.Now(), ""}, errors.New(error_int)
  }

  // Calculate Paschal
  easterday := paschalAlgorithm(year)

  // Make time string
  timeString := fmt.Sprintf("%[1]d-%[2]d-%[3]d",
    easterday["year"],
    easterday["month"],
    easterday["day"],
  )

  parsed, _ := time.ParseInLocation(layout, timeString, location)
  return Easter{parsed, "Easterday"}, nil
}

func List() (arr []string) {
  for _, name := range nameMap {
    arr = append(arr, name)
  }
  return
}

func (d Easter) Day() time.Time {
  return d.Date
}

func (d Easter) Get(arg ...int) (arr EasterCollection) {
  for _, value := range arg {
    key := nameMap[value]
    arr = append(arr, Easter{d.Date.AddDate(0,0, value), key})
  }
  return
}

