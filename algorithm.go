/*
 * Algorithm translated from https://www.assa.org.au/edm by Ronald Mallen.
 */

/*
 * EASTER DATE CALCULATION FOR YEARS 1583 TO 4099

 * y is a 4 digit year 1583 to 4099
 * d returns the day of the month of Easter
 * m returns the month of Easter

 * Easter Sunday is the Sunday following the Paschal Full Moon
 * (PFM) date for the year

 * This algorithm is an arithmetic interpretation of the 3 step
 * Easter Dating Method developed by Ron Mallen 1985, as a vast
 * improvement on the method described in the Common Prayer Book

 * Because this algorithm is a direct translation of the
 * official tables, it can be easily proved to be 100% correct

 * This algorithm derives values by sequential inter-dependent
 * calculations, so ... DO NOT MODIFY THE ORDER OF CALCULATIONS!

 * The \ operator may be unfamiliar - it means integer division
 * for example, 30 \ 7 = 4 (the remainder is ignored)

 * All variables are integer data types

 * It's free!  Please do not modify code or comments!
 */

package easter

func paschalAlgorithm(year int) (obj map[string]int) {
  obj = make(map[string]int)

  // Integer division
  century := year / 100

  // Divide
  remain19 := year - (19 * (year / 19))

  // Calculate PFM
  c1 := (century - 15) / 2
  temp := c1 + 202 - 11 * int(remain19)

  temp = temp - (30 * (temp / 30))

  tA := int(temp + 21)

  if temp == 29 || (temp == 28 && remain19 > 10) {
    tA = tA - 1
  }

  // Find next sunday
  tA_temp := tA - 19
  tB := tA_temp - (7 * (tA_temp / 7))

  tC_temp := 40 - century
  tC := tC_temp - (4 * (tC_temp / 4))

  if tC == 3 || tC > 1 {
    tC = tC + 1
  }

  temp = year - (100 * (year / 100))

  tD_temp := temp + temp / 4
  tD := tD_temp - (7 * (tD_temp / 7))

  tE_temp := int(20) - int(tB) - int(tC) - int(tD)
  tE := tE_temp - (7 * (tE_temp / 7)) + 1

  d := int(tA) + int(tE)

  // Return date
  if d > 31 {
    d = d - 31
    m := 4

    obj["month"] = int(m)
    obj["day"] = int(d)
    obj["year"] = int(year)
  } else {
    m := 3

    obj["month"] = int(m)
    obj["day"] = int(d)
    obj["year"] = int(year)
  }
  return obj
}

