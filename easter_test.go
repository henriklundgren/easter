package easter

import (
  "testing"
  "time"
  "fmt"
)

var test_layout string = "20060102"

var test_easterDates = map[int]string{
  1996: "0407",
  1997: "0330",
  1998: "0412",
  1999: "0404",
  2000: "0423",
  2001: "0415",
  2002: "0331",
  2003: "0420",
  2004: "0411",
  2005: "0327",
  2006: "0416",
  2007: "0408",
  2008: "0323",
  2009: "0412",
  2010: "0404",
  2011: "0424",
  2012: "0408",
  2013: "0331",
  2014: "0420",
}

var test_goodFridays = map[int]string{
  2012: "0406",
  2013: "0329",
  2014: "0418",
}

var test_corpusChristi = map[int]string{
  2013: "0530",
  2014: "0619",
  2015: "0604",
  2016: "0526",
}

func Test_Set_0(t *testing.T) {
  if _, err := Set(2014, "", ""); err != nil {
    t.Error("Should return time ", err)
  }
}

func Test_Set_1(t *testing.T) {
  if _, err := Set(2014, "en_US", "Europe/Berlin"); err != nil {
    t.Error("Should return time ", err)
  }
}

func Test_Day_0(t *testing.T) {

  // Loop known dates
  for year, easterDate := range test_easterDates {
    knowEasterDate := fmt.Sprintf("%d%s", year, easterDate)
    parsed, parseError := time.Parse(test_layout, knowEasterDate)

    if parseError != nil {
      t.Error("Error")
    }

    // Ask for easter date on year year
    date, _ := Set(year, "", "")
    if checkDate := date.Day(); checkDate != parsed {
      t.Error("Error")
    }
  }
}

// Check if my function matches known goodfriday dates.
func Test_GoodFriday_0(t *testing.T) {
  for year, fridays := range test_goodFridays {
    knownDatesString := fmt.Sprintf("%d%s", year, fridays)
    knownDatesParsed, _ := time.Parse(test_layout, knownDatesString)

    date, _ := Set(year, "", "")
    if goodfriday := date.GoodFriday(); goodfriday != knownDatesParsed {
      t.Error("Error")
    }
  }
}

// Check that my function doesnt match random dates.
func Test_GoodFriday_1(t *testing.T) {
  randomDate, _ := time.Parse(test_layout, "20141224")
  date, _ := Set(2014, "", "")
  if goodfriday := date.GoodFriday(); goodfriday == randomDate {
    t.Error("Error")
  }
}

func Test_CorpusChristi_0(t *testing.T) {
  for year, corpus := range test_corpusChristi {
    knownDatesString := fmt.Sprintf("%d%s", year, corpus)
    knownDatesParsed, _ := time.Parse(test_layout, knownDatesString)

    date, _ := Set(year, "", "")
    if current := date.CorpusChristi(); current != knownDatesParsed {
      t.Error("Error")
    }
  }
}
