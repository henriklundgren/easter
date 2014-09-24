/*

Package Easter calculates and returns the Easterday date.

Introduction

From Wikipedia:
Easter is a moveable feast, meaning it is not fixed in relation to the civil calendar.
The First Council of Nicaea (325) established the date of Easter as
the first Sunday after the full moon (the Paschal Full Moon) following the March equinox.

Credit

The Paschal algorithm utilized in the code is based solely on the work of Ronald Mallen.

Usage

var year int = 2014
var location *time.Location = time.Local

easter, err := easter.Set(year, location)
if err != nil {
  panic(err)
}

// Return the easterday as type time.Time
easterday := easter.Day()

// Return a collection of the specified feastdays.
easterCollection := easter.Get(easter.Day)

*/

package easter
