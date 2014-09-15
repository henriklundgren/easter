/*
Package Easter calculates and returns the Easterday date.

Introduction

From Wikipedia:
Easter is a moveable feast, meaning it is not fixed in relation to the civil calendar.
The First Council of Nicaea (325) established the date of Easter as
the first Sunday after the full moon (the Paschal Full Moon) following the March equinox.

Credit

The paschal algorithm utilized in the code is based solely on the work of Ronald Mallen.

Usage

easter, _ := easter.Set(year int, locale string, location string) // Easter
easter.Day() // time.Time

*/

package easter
