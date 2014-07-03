package grid

import (
  "os"
  "bufio"
  "strings"
  "strconv"
)

type Anomaly struct {
  Month int
  Day int
  GridBoxID int
  Longitude float64 // lower left corner of grid box (degrees)
  Latitude float64 // lower left corner of grid box (degrees)
  TempAnomaly float64 // anomaly (whole degrees Celsius) 
}

func ParseFile( path string ) ( []Anomaly ){

  file, err := os.Open(path)

  if err != nil {
    panic(err)
  }
  
  defer file.Close()
  
  var days []Anomaly
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    day := ParseLine(scanner.Text())
    days = append(days, day)
  }

  return days
}

func ParseLine( line string ) ( Anomaly ) {
  scanner := bufio.NewScanner(strings.NewReader(line))
  scanner.Split(bufio.ScanWords)

  var grid Anomaly
  for i := 0; scanner.Scan(); i++ {
    text := scanner.Text()
    switch i {
      case 0: grid.Month       = StrToInt(text)
      case 1: grid.Day         = StrToInt(text)
      case 2: grid.GridBoxID   = StrToInt(text)
      case 3: grid.Longitude   = StrToFloat64(text)
      case 4: grid.Latitude    = StrToFloat64(text)
      case 5: grid.TempAnomaly = StrToFloat64(text)
      case 6: panic("Overflow in struct parsing")
    }
  }

  return grid
}

func StrToFloat64( s string ) float64 {
  f, err := strconv.ParseFloat(s, 32)
  if err != nil {
    panic(err)
  }
  return f
}

func StrToInt( s string ) int {
  i, err := strconv.Atoi(s)
  if err != nil {
    panic(err)
  }
  return i
}