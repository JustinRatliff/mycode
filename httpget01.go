/* Alta3 Research | RZFeeser
   HTTP GET requests  */
   
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

func main() {

    // create a GET request to the page
    resp, err := http.Get("https://alta3.com")

    // check for errors
    if err != nil {
        log.Fatal(err)
    }

    // the client must close the response body when finished
    defer resp.Body.Close()

    // read the content of the body with "ReadAll()"
    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        log.Fatal(err)
    }

    // print received data to the console
    fmt.Println(string(body))
}

