package daily

import (
  "os"
  "encoding/csv"
  "time"
  // "bufio"
  // "strings"
  // "strconv"
  // "fmt"
)

type Daily struct {
  StationID string
  Year int
  Month int
  Element string
  Value1 int
  MFlag1 string
  QFlag1 string
  SFlag1 string
  Value2 int
  MFlag2 string
  QFlag2 string
  SFlag2 string

}

type Parser struct {
  file * os.File
  reader * csv.Reader
}

func ( d * Parser ) OpenFile( path string ) {

  var err error
  d.file, err = os.Open(path)

  if err != nil {
    panic(err)
  }

  d.reader = csv.NewReader( d.file )

}

func ( d * Parser ) CloseFile(){
 
  d.file.Close()

}

func ( d * Parser ) Read() ( Daily, error ) {
  var daily Daily

  record, err := d.reader.Read() 

  if err != nil {
    return daily, err
  }

  daily.StationID = record[0]
  
  return daily, nil
}

