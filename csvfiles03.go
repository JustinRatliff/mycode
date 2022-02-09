/* RZFeeser | Alta3 Research
   Reading CSV file - ReadAll() - With custom delimiters*/


package main

import (
    "encoding/csv"
    "fmt"
    "log"
    "os"
)

func main() {

    f, err := os.Open("origin.csv")

    if err != nil {

        log.Fatal(err)
    }

    // set the separator and the comment character
    // so that the package knows how to parse the file
    r := csv.NewReader(f)
    r.Comma = ';'            // delimiter
    r.Comment = '#'          // comment at the top of the file

    records, err := r.ReadAll()

    if err != nil {
        log.Fatal(err)
    }

    fmt.Print(records)
}    

