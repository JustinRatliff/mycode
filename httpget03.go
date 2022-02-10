/* RZFeeser | Alta3 Research
   HTTP GET and write out to a file */


package main

import (
    "log"
    "net/http"
    "os"
)

func main() {

    r, err := http.Get("https://alta3.com")

    if err != nil {
      log.Fatal(err)
    }

    defer r.Body.Close()

    // create a new file 
    f, err := os.Create("alta3index.html")

    if err != nil {
      log.Fatal(err)
    }

    defer f.Close()

    // write the contents to the file with ReadFrom()
    _, err = f.ReadFrom(r.Body)

    if err != nil {
      log.Fatal(err)
    }
}

