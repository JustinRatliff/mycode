/* RZFeeser | Alta3 Research
   HTTP GET with io.Copy() */
   
package main

import (
    "io"
    "log"
    "net/http"
    "os"
)

func main() {

    resp, err := http.Get("http://webcode.me")

    if err != nil {
      log.Fatal(err)
    }

    defer resp.Body.Close()
    
    
    /* The io.Copy() function copies from source to destination until
       either EOF is reached on source or an error occurs. Returns the number of bytes
       copied and the first error encountered while copying, if any. */
    
    _, err = io.Copy(os.Stdout, resp.Body)

    if err != nil {
      log.Fatal(err)
    }
}

