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

func Test_Set_0(t *testing.T) {
  if _, err := Set(2014, time.Local); err != nil {
    t.Error("Should return time ", err)
  }
}

func Test_Set_1(t *testing.T) {
  // Short test
  if testing.Short() {
    t.Skip("Skip second Set test in short mode.")
  }

  if _, err := Set(2014, time.Local); err != nil {
    t.Error("Should return time ", err)
  }
}

func Test_Day_0(t *testing.T) {

  // Short test
  if testing.Short() {
    pt, _ := time.Parse(test_layout, "20140420")
    et, _ := Set(2014, time.UTC)
    if try := et.Day(); pt != try {
      t.Error("error")
    }
  }

  // Loop known dates
  for year, easterDate := range test_easterDates {
    knowEasterDate := fmt.Sprintf("%d%s", year, easterDate)
    parsed, parseError := time.Parse(test_layout, knowEasterDate)

    if parseError != nil {
      t.Error("Error")
    }

    // Ask for easter date on year year
    date, _ := Set(year, time.UTC)
    if checkDate := date.Day(); checkDate != parsed {
      t.Error("Error")
    }
  }
}

func Test_Const_0(t *testing.T) {
  localSexagesima := time.Date(2014, 02, 23, 0, 0, 0, 0, time.Local)
  easterDate, _ := Set(2014, time.Local)
  if single := easterDate.Day(); single.AddDate(0,0, Sexagesima) != localSexagesima {
    t.Error("Error")
  }

}

func Test_Day_1(t *testing.T) {
  testing, _ := Set(2014, time.Local)
  if arr := testing.Get(Sexagesima, Monday, GoodFriday); arr == nil {
    t.Error("Error")
  }
}

func Test_List_0(t *testing.T) {
  testing := List()
  if testing == nil {
    t.Error("error")
  }
}

